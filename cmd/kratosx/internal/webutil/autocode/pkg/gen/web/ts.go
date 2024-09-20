package web

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	proto3 "github.com/emicklei/proto"

	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen"
)

type Ts struct {
	*gen.Builder
	options map[string]bool
}

type Result struct {
	Interface string
	Api       string
}

type TypeScriptCode struct {
	sort   []string
	bucket map[string]string
}

type httpOption struct {
	Verb string
	Path string
}

func NewTsBuilder(builder *gen.Builder) *Ts {
	return &Ts{
		Builder: builder,
	}
}

func (t *Ts) protoParser(path string) (*proto3.Proto, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	parser := proto3.NewParser(bytes.NewReader(content))
	definition, err := parser.Parse()
	if err != nil {
		return nil, err
	}
	return definition, nil
}

// ProtoToTSType 将给定的proto类型字符串转换为对应的TypeScript类型字符串
func (t *Ts) ProtoToTSType(protoType string) string {
	switch protoType {
	case "double", "float", "int32", "uint32", "sint32", "fixed32", "sfixed32":
		return "number"
	case "int64", "uint64", "sint64", "fixed64", "sfixed64":
		// 考虑使用 `string` 或 `number` 取决于是否使用BigInt或者其他库
		// 这里我们选择 `string` 作为默认类型以避免精度问题
		return "string"
	case "bool":
		return "boolean"
	case "string":
		return "string"
	case "bytes":
		return "Uint8Array"
	case "enum":
		// TypeScript中的枚举可以使用字符串或数字类型
		// 这里我们选择 `number` 作为默认类型
		return "number"
	default:
		return protoType // 使用相同的类型名称
	}
}

func (t *Ts) genInterface(m *proto3.Message) TypeScriptCode {
	var (
		code = "export "
		ti   = TypeScriptCode{bucket: make(map[string]string)}
	)
	if len(m.Elements) == 0 {
		return ti
	}

	code = code + fmt.Sprintf("interface %s {\n", m.Name)
	optionLen := 0
	for _, e := range m.Elements {
		if field, ok := e.(*proto3.NormalField); ok {
			name := field.Name
			if field.Optional {
				name = name + "?"
				optionLen++
			}
			tp := t.ProtoToTSType(field.Type)
			if field.Repeated {
				tp = tp + "[]"
			}
			code = code + fmt.Sprintf("  %s: %s;\n", name, tp)
		}
		if cm, ok := e.(*proto3.Message); ok {
			res := t.genInterface(cm)
			ti.sort = append(ti.sort, res.sort...)
			for k, v := range res.bucket {
				ti.bucket[k] = v
			}
		}
	}

	code += "}"
	if len(m.Elements) != 0 {
		ti.sort = append(ti.sort, m.Name)
		ti.bucket[m.Name] = code
		t.options[m.Name] = len(m.Elements) == optionLen
	}

	return ti
}

func (t *Ts) GenTypeScript(msg string, api string) (*Result, error) {
	var result Result
	code, err := t.GenTypeByProtoPath(msg)
	if err != nil {
		return nil, err
	}
	result.Interface = code

	code, err = t.GenApiByProtoPath(msg)
	if err != nil {
		return nil, err
	}
	result.Api = code
	return &result, nil
}

// GenTypeByProtoPath 生成
func (t *Ts) GenTypeByProtoPath(path string) (string, error) {
	proto, err := t.protoParser(path)
	if err != nil {
		return "", err
	}

	var (
		tsCode = strings.Builder{}
		ti     = TypeScriptCode{bucket: make(map[string]string)}
	)

	proto3.Walk(proto,
		proto3.WithMessage(func(m *proto3.Message) {
			itr := t.genInterface(m)
			ti.sort = append(ti.sort, itr.sort...)
			for k, v := range itr.bucket {
				ti.bucket[k] = v
			}
		}),
	)

	for _, k := range ti.sort {
		tsCode.WriteString(ti.bucket[k] + "\n\n")
	}
	// 生成interface
	return tsCode.String(), nil
}

// GenApiByProtoPath 生成
func (t *Ts) GenApiByProtoPath(path string) (string, error) {
	proto, err := t.protoParser(path)
	if err != nil {
		return "", err
	}

	var (
		ts          = strings.Builder{}
		api         = TypeScriptCode{bucket: make(map[string]string)}
		importTypes []string
	)

	proto3.Walk(proto,
		proto3.WithService(func(s *proto3.Service) {
			for _, e := range s.Elements {
				if rpc, ok := e.(*proto3.RPC); ok {
					methodName := rpc.Name // Capitalize the first letter
					hopt := t.getHttpOption(rpc)
					tsCode := strings.Builder{}
					comment := ""

					if hopt != nil {
						requestType := rpc.RequestType
						responseType := rpc.ReturnsType

						isReqOpt, isReq := t.options[requestType]
						_, isRsp := t.options[responseType]

						for _, cmt := range rpc.Comment.Lines {
							comment = comment + "//" + cmt + "\n"
						}
						if responseType == "google.protobuf.Empty" || !isRsp {
							responseType = ""
						} else {
							importTypes = append(importTypes, responseType)
							responseType = fmt.Sprintf("<%s>", responseType)
						}

						if requestType == "google.protobuf.Empty" || !isReq {
							requestType = ""
						} else {
							importTypes = append(importTypes, requestType)
						}
						tsCode.WriteString(comment)
						// Generate TypeScript function
						switch hopt.Verb {
						case "get", "delete":
							if requestType != "" {
								if isReq && isReqOpt {
									requestType = "params?: " + requestType
								} else {
									requestType = "params: " + requestType
								}
							}
							tsCode.WriteString(fmt.Sprintf("export function %s(%s) {\n", methodName, requestType))
							if requestType == "" {
								tsCode.WriteString(fmt.Sprintf("    return axios.%s%s('%s');\n", hopt.Verb, responseType, hopt.Path))
							} else {
								tsCode.WriteString(fmt.Sprintf("    return axios.%s%s('%s', { params });\n", hopt.Verb, responseType, hopt.Path))
							}
						case "post", "put":
							if requestType != "" {
								if isReq && isReqOpt {
									requestType = "data?: " + requestType
								} else {
									requestType = "data: " + requestType
								}
							}
							tsCode.WriteString(fmt.Sprintf("export function %s(%s) {\n", methodName, requestType))
							if requestType == "" {
								tsCode.WriteString(fmt.Sprintf("    return axios.%s%s('%s');\n", hopt.Verb, responseType, hopt.Path))
							} else {
								tsCode.WriteString(fmt.Sprintf("    return axios.%s%s('%s', data);\n", hopt.Verb, responseType, hopt.Path))
							}
						}
						tsCode.WriteString("}")
						api.sort = append(api.sort, methodName)
						api.bucket[methodName] = tsCode.String()
					}
				}
			}
		}),
	)

	ts.WriteString("import axios from 'axios';")
	ts.WriteString(fmt.Sprintf("import {\n\t%s\n} from './type';", strings.Join(importTypes, ",\n\t")))

	for _, k := range api.sort {
		ts.WriteString(api.bucket[k] + "\n\n")
	}
	// 生成interface
	return ts.String(), nil
}

// getHttpOption 获取HTTP选项
func (t *Ts) getHttpOption(rpc *proto3.RPC) *httpOption {
	for _, e := range rpc.Elements {
		if option, ok := e.(*proto3.Option); ok {
			if option.Name == "(google.api.http)" {
				for _, p := range option.AggregatedConstants {
					return &httpOption{
						Verb: strings.Trim(p.Name, "{}"),
						Path: p.Literal.Source,
					}
				}
			}
		}
	}
	return nil
}
