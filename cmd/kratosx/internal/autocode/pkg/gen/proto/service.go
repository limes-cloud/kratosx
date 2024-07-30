package proto

import (
	"fmt"
	"github.com/emicklei/proto"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/autocode/pkg"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/autocode/pkg/gen"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/autocode/pkg/gen/types"
	"os"
	"regexp"
	"strings"
)

type Server struct {
	*gen.Builder
}

type ServiceCode struct {
	pkg     string
	options []string
	imports []string
	sort    []string
	bucket  map[string]string
}

func NewServerBuilder(builder *gen.Builder) *Server {
	return &Server{Builder: builder}
}

func (p *Server) dir(object *types.Table) string {
	return strings.ToLower(fmt.Sprintf("api/%s/%s", p.Server, object.Module))
}

func (p *Server) version() string {
	return "v1"
}

func (p *Server) goPackage(object *types.Table) string {
	return "./v1;v1"
}

func (p *Server) packageName(object *types.Table) string {
	return strings.ReplaceAll(p.Server+"/"+p.dir(object)+"/"+p.version(), "/", ".")
}

func (p *Server) javaPackage(object *types.Table) string {
	return strings.ReplaceAll(p.Server+"/"+p.dir(object)+"/"+p.version(), "/", ".")
}

func (p *Server) javaClass(object *types.Table) string {
	s := strings.Split(p.dir(object)+"/"+p.version(), "/")
	return pkg.ToUpperHump(object.Module) + pkg.ToUpperHump(s[len(s)-1])
}

func (p *Server) objectName(object *types.Table) string {
	return pkg.ToUpperHump(object.Struct)
}

func (p *Server) objectComment(object *types.Table) string {
	return pkg.ToUpperHump(object.Comment)
}

// ParseServerByContent 解析proto文本，生成proto结构
func (p *Server) ParseServerByContent(content string) (*ServiceCode, error) {
	reply := &ServiceCode{bucket: make(map[string]string)}

	parser := proto.NewParser(strings.NewReader(content))
	definition, err := parser.Parse()
	if err != nil {
		return nil, err
	}

	// 正则扫描定义
	//re := regexp.MustCompile(`// (.+)\nexport function (\w+)([\s]*?)\(([\s\S]*?)\n\}`)
	re := regexp.MustCompile(`(\/\/[^\n]*)?\n\s*rpc\s+(\w+)\s*\((\w+)\)\s*returns\s*\((\w+)\)\s*\{(.*?)\n\s*\}`)

	matches := re.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		if len(match) == 4 { // 0 是整个匹配项，1 是消息名称，2 是消息体
			body := match[0]
			name := match[1]
			reply.sort = append(reply.sort, name)
			reply.bucket[name] = body
		}
	}

	// 解析 import\option\package\message
	proto.Walk(definition,
		proto.WithPackage(func(p *proto.Package) {
			reply.pkg = p.Name
		}),
		proto.WithImport(func(m *proto.Import) {
			reply.imports = append(reply.imports, fmt.Sprintf(`import "%s"`, m.Filename))
		}),
		proto.WithOption(func(option *proto.Option) {
			text := ""
			if option.Constant.IsString {
				text = fmt.Sprintf(`option %s = "%s"`, option.Name, option.Constant.Source)
			} else {
				text = fmt.Sprintf(`option %s = %s`, option.Name, option.Constant.Source)
			}
			reply.options = append(reply.options, text)
		}),
	)

	return reply, nil
}

// ScanServer 扫描现在的已经存在的proto error定义
func (p *Server) ScanServer() (*ServiceCode, error) {
	path := p.ProtoServerPath()
	if !pkg.IsExistFile(path) {
		return &ServiceCode{}, nil
	}
	code, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return p.ParseServerByContent(string(code))
}

// MakeServer 根据模板生成proto
func (p *Server) MakeServer() (*ServiceCode, error) {
	code, err := os.ReadFile(p.ProtoServerTplPath())
	if err != nil {
		return nil, err
	}
	return p.ParseServerByContent(string(code))
}

func (p *Server) RenderServer(data *ServiceCode) string {
	content := "syntax = \"proto3\";\n\n"
	content += data.pkg + ";\n\n"
	for _, val := range data.options {
		content += val + "\n"
	}

	content += "\n"
	for _, val := range data.imports {
		content += val + "\n"
	}

	trash := p.HasDeletedAt()

	content += "\n"
	content += fmt.Sprintf("service %s{\n\n", p.Server)
	for _, val := range data.sort {
		if trash && !strings.Contains(val, "Trash") {
			continue
		}
		content += data.bucket[val] + "\n"
	}
	content += "}"
	return content
}

// GenServer 生成error proto
func (p *Server) GenServer() (string, error) {
	// 扫描已经生成的proto
	scanServer, err := p.ScanServer()
	if err != nil {
		return "", fmt.Errorf("扫描proto代码失败，%s", err.Error())
	}

	// 生成新的proto
	makeServer, err := p.MakeServer()
	if err != nil {
		return "", fmt.Errorf("生成proto代码失败，%s", err.Error())
	}

	// 合并proto
	makeServer.sort = append(makeServer.sort, scanServer.sort...)
	for key, val := range scanServer.bucket {
		makeServer.bucket[key] = val
	}

	// 去重
	makeServer.sort = pkg.UniqueArray(makeServer.sort)

	// 生成proto
	return p.RenderServer(makeServer), nil
}
