package autocode

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"

	proto3 "github.com/emicklei/proto"
)

type ts struct {
	store map[string]*msg
}

type msg struct {
	isAllOption bool
}

var (
	t    *ts
	once sync.Once
)

type tsInterface struct {
	sort []string
	m    map[string]string
}

func GenTypeScript(path string, content string) (string, string, error) {
	once.Do(func() {
		t = &ts{store: make(map[string]*msg)}
	})

	if strings.HasSuffix(path, "error_reason.proto") {
		return "", "", nil
	}

	if strings.HasSuffix(path, "service.proto") {
		code, err := t.renderProtoApi(content)
		if err != nil {
			return "", "", err
		}
		return t.apiPath(path), code, nil
	}

	if strings.HasSuffix(path, ".proto") {
		code, err := t.renderProtoMessage(path, content)
		if err != nil {
			return "", "", err
		}
		return t.typePath(path), code, nil
	}
	return "", "", nil
}

func (t *ts) apiPath(path string) string {
	path = "web_ts/" + path
	path = path[:strings.LastIndex(path, "/")]
	path = path + "/api.ts"
	return path
}

func (t *ts) typePath(path string) string {
	path = "web_ts/" + path
	path = path[:strings.LastIndex(path, "/")]
	path = path + "/type.ts"
	return path
}

// parseProto 解析 .proto 文件并返回语法树
func (t *ts) renderProtoMessage(path, content string) (string, error) {
	parser := proto3.NewParser(strings.NewReader(content))
	definition, err := parser.Parse()
	if err != nil {
		return "", err
	}
	return t.generateTypeScript(path, definition), nil
}

func (t *ts) scanMessage(content string) tsInterface {
	ti := tsInterface{m: make(map[string]string)}
	re := regexp.MustCompile(`export interface (\w+)([\s]*?)\{([\s\S]*?)\n\}`)
	matches := re.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		if len(match) == 4 { // 0 是整个匹配项，1 是消息名称，2 是消息体
			messageBody := match[0]
			messageName := match[1]
			ti.sort = append(ti.sort, messageName)
			ti.m[messageName] = messageBody
		}
	}
	return ti
}

// generateTypeScript 根据解析出的 .proto 语法树生成 TypeScript 代码
func (t *ts) generateTypeScript(path string, definition *proto3.Proto) string {
	var (
		tsCode strings.Builder
		ti     = tsInterface{m: make(map[string]string)}
	)
	oriByte, err := os.ReadFile(t.typePath(path))
	if err == nil {
		ti = t.scanMessage(string(oriByte))
	}

	//tsCode.WriteString("/* eslint-disable @typescript-eslint/no-empty-interface */\n")

	proto3.Walk(definition,
		proto3.WithMessage(func(m *proto3.Message) {
			itr := t.genInterface(m)
			ti.sort = append(ti.sort, itr.sort...)
			for k, v := range itr.m {
				ti.m[k] = v
			}
		}),
	)
	ti.sort = uniqueStrings(ti.sort)
	for _, k := range ti.sort {
		tsCode.WriteString(ti.m[k] + "\n\n")
	}
	//t.store = append(t.store, ti.sort...)
	return tsCode.String()
}

// ProtoToTSType 将给定的proto类型字符串转换为对应的TypeScript类型字符串
func (t *ts) ProtoToTSType(protoType string) string {
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
		// 处理复杂类型如message、repeated和map，以及Well-Known Types
		// 这里需要根据具体的proto文件中的定义来返回正确的TypeScript类型
		// 对于message类型，假设已有一个同名的TypeScript接口或类
		// 对于repeated和map类型，则需要额外的信息来确定元素类型
		// 对于Well-Known Types，使用对应的TypeScript类型或者any类型
		if t.isWellKnownType(protoType) {
			return t.wellKnownTypeToTS(protoType)
		} else if t.isRepeatedType(protoType) {
			elementType := t.getElementType(protoType)
			return "Array<" + t.ProtoToTSType(elementType) + ">"
		} else if t.isMapType(protoType) {
			keyType, valueType := t.getMapTypes(protoType)
			return fmt.Sprintf("{ [key: %s]: %s }", t.ProtoToTSType(keyType), t.ProtoToTSType(valueType))
		}
		// 默认情况下，假设它是一个message类型
		return protoType // 使用相同的类型名称
	}
}

func (t *ts) genInterface(m *proto3.Message) tsInterface {
	var (
		code = "export "
		ti   = tsInterface{m: make(map[string]string)}
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
			for k, v := range res.m {
				ti.m[k] = v
			}
		}
	}
	t.store[m.Name] = &msg{isAllOption: optionLen == len(m.Elements)}
	code += "}"
	ti.sort = append(ti.sort, m.Name)
	ti.m[m.Name] = code
	return ti
}

// 以下函数是为了完整性而添加的，但没有具体实现。
// 实际中，这些函数的实现取决于proto文件的具体内容。
func (t *ts) isWellKnownType(protoType string) bool {
	// 判断是否是Well-Known Type
	return false
}

func (t *ts) wellKnownTypeToTS(protoType string) string {
	// 将Well-Known Type映射到TypeScript类型
	return "any"
}

func (t *ts) isRepeatedType(protoType string) bool {
	// 判断是否是repeated类型
	return false
}

func (t *ts) getElementType(protoType string) string {
	// 获取repeated类型的元素类型
	return ""
}

func (t *ts) isMapType(protoType string) bool {
	// 判断是否是map类型
	return false
}

func (t *ts) getMapTypes(protoType string) (string, string) {
	// 获取map类型的键和值类型
	return "", ""
}

// generateAxiosTypescript 生成Axios请求的TypeScript代码
func (t *ts) renderProtoApi(content string) (string, error) {
	parser := proto3.NewParser(strings.NewReader(content))
	definition, err := parser.Parse()
	if err != nil {
		return "", err
	}

	var (
		tsCode  strings.Builder
		imports []string
	)

	proto3.Walk(definition,
		proto3.WithService(func(s *proto3.Service) {
			for _, e := range s.Elements {
				if rpc, ok := e.(*proto3.RPC); ok {
					methodName := rpc.Name // Capitalize the first letter
					httpOption := t.getHttpOption(rpc)
					comment := ""
					if httpOption != nil {
						requestType := rpc.RequestType
						responseType := rpc.ReturnsType
						for _, cmt := range rpc.Comment.Lines {
							comment = comment + "//" + cmt + "\n"
						}
						respMsg := t.store[responseType]
						if responseType == "google.protobuf.Empty" || respMsg == nil {
							responseType = ""
						} else {
							imports = append(imports, responseType)
							responseType = fmt.Sprintf("<%s>", responseType)
						}

						reqMsg := t.store[requestType]
						if requestType == "google.protobuf.Empty" || reqMsg == nil {
							requestType = ""
						} else {
							imports = append(imports, requestType)
						}
						tsCode.WriteString(comment)
						// Generate TypeScript function
						switch httpOption.Verb {
						case "get", "delete":
							if requestType != "" {
								if reqMsg != nil && reqMsg.isAllOption {
									requestType = "params?: " + requestType
								} else {
									requestType = "params: " + requestType
								}
							}
							tsCode.WriteString(fmt.Sprintf("export function %s(%s) {\n", methodName, requestType))
							if requestType == "" {
								tsCode.WriteString(fmt.Sprintf("    return axios.%s%s('%s');\n", httpOption.Verb, responseType, httpOption.Path))
							} else {
								tsCode.WriteString(fmt.Sprintf("    return axios.%s%s('%s', { params });\n", httpOption.Verb, responseType, httpOption.Path))
							}
						case "post", "put":
							if requestType != "" {
								if reqMsg != nil && reqMsg.isAllOption {
									requestType = "data?: " + requestType
								} else {
									requestType = "data: " + requestType
								}
							}
							tsCode.WriteString(fmt.Sprintf("export function %s(%s) {\n", methodName, requestType))
							if requestType == "" {
								tsCode.WriteString(fmt.Sprintf("    return axios.%s%s('%s');\n", httpOption.Verb, responseType, httpOption.Path))
							} else {
								tsCode.WriteString(fmt.Sprintf("    return axios.%s%s('%s', data);\n", httpOption.Verb, responseType, httpOption.Path))
							}
						}
						tsCode.WriteString("}\n\n")
					}
				}
			}
		}),
	)
	importTypes := "import axios from 'axios';\n"
	importTypes = importTypes + fmt.Sprintf("import {\n\t%s\n} from './type';\n\n", strings.Join(uniqueStrings(imports), ",\n\t"))
	return importTypes + tsCode.String(), nil
}

type httpOption struct {
	Verb string
	Path string
}

// getHttpOption 获取HTTP选项
func (t *ts) getHttpOption(rpc *proto3.RPC) *httpOption {
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
