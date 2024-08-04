package proto

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"

	"github.com/emicklei/proto"

	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen/types"
)

type Service struct {
	*gen.Builder
}

type ServiceCode struct {
	pkg     string
	options []string
	imports []string
	sort    []string
	bucket  map[string]string
}

func NewServiceBuilder(builder *gen.Builder) *Service {
	return &Service{Builder: builder}
}

func (p *Service) dir(object *types.Table) string {
	return strings.ToLower(fmt.Sprintf("api/%s/%s", p.Server, object.Module))
}

func (p *Service) version() string {
	return "v1"
}

func (p *Service) goPackage(object *types.Table) string {
	return "./v1;v1"
}

func (p *Service) packageName(object *types.Table) string {
	return strings.ReplaceAll(p.Server+"/"+p.dir(object)+"/"+p.version(), "/", ".")
}

func (p *Service) javaPackage(object *types.Table) string {
	return strings.ReplaceAll(p.Server+"/"+p.dir(object)+"/"+p.version(), "/", ".")
}

func (p *Service) javaClass(object *types.Table) string {
	s := strings.Split(p.dir(object)+"/"+p.version(), "/")
	return pkg.ToUpperHump(object.Module) + pkg.ToUpperHump(s[len(s)-1])
}

func (p *Service) objectName(object *types.Table) string {
	return pkg.ToUpperHump(object.Struct)
}

func (p *Service) objectComment(object *types.Table) string {
	return pkg.ToUpperHump(object.Comment)
}

// ParseServiceByContent 解析proto文本，生成proto结构
func (p *Service) ParseServiceByContent(content string) (*ServiceCode, error) {
	reply := &ServiceCode{bucket: make(map[string]string)}

	parser := proto.NewParser(strings.NewReader(content))
	definition, err := parser.Parse()
	if err != nil {
		return nil, err
	}

	// 正则扫描定义
	// re := regexp.MustCompile(`// (.+)\nexport function (\w+)([\s]*?)\(([\s\S]*?)\n\}`)
	re := regexp.MustCompile(`service\s+(\w+)\s*\{([\s\S]*)\}`)
	matches := re.FindAllStringSubmatch(content, -1)
	if len(matches) != 1 && len(matches[0]) != 3 {
		return nil, errors.New("匹配数据错误")
	}

	rpcCode := strings.TrimSpace(matches[0][2])
	codeArr := strings.Split(rpcCode, "\n\n")

	keyRe := regexp.MustCompile(`rpc\s+(\w+)`)

	for _, code := range codeArr {
		code = strings.TrimSpace(code)
		matches := keyRe.FindAllStringSubmatch(code, -1)
		if len(matches) != 0 && len(matches[0]) != 2 {
			continue
		}
		reply.sort = append(reply.sort, matches[0][1])
		reply.bucket[matches[0][1]] = "  " + code + "\n"
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
			if option.Constant.Source == "" {
				return
			}
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

// ScanService 扫描现在的已经存在的proto error定义
func (p *Service) ScanService() (*ServiceCode, error) {
	path := p.ProtoServicePath()
	if !pkg.IsExistFile(path) {
		return &ServiceCode{}, nil
	}
	code, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return p.ParseServiceByContent(string(code))
}

// MakeService 根据模板生成proto
func (p *Service) MakeService() (*ServiceCode, error) {
	tp, err := os.ReadFile(p.ProtoServiceTplPath())
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("Message").Parse(strings.TrimSpace(string(tp)))
	if err != nil {
		return nil, err
	}

	renderData := map[string]any{
		"EnableBatchDelete":        p.Table.EnableBatchDelete,
		"Package":                  p.packageName(p.Table),
		"GoPackage":                p.goPackage(p.Table),
		"JavaPackage":              p.javaPackage(p.Table),
		"JavaClass":                p.javaClass(p.Table),
		"Classify":                 pkg.ToSnake(p.Table.Module),
		"Server":                   pkg.ToSnake(p.Server),
		"Object":                   p.objectName(p.Table),
		"ObjectPluralizeLowerCase": pkg.ToPluralize(pkg.ToLowerHump(p.objectName(p.Table))),
		"ObjectLowerCase":          pkg.ToSnake(p.objectName(p.Table)),
		"ServerLowerCase":          pkg.ToLowerHump(p.Server),
		"ModuleLowerCase":          pkg.ToLowerHump(p.Table.Module),
		"Title":                    p.Table.Comment,
	}
	if err := tmpl.Execute(buf, renderData); err != nil {
		return nil, err
	}

	return p.ParseServiceByContent(buf.String())
}

func (p *Service) RenderService(data *ServiceCode) string {
	content := "syntax = \"proto3\";\n\n"
	content += "package " + data.pkg + ";\n\n"
	for _, val := range data.options {
		content += val + ";\n"
	}

	content += "\n"
	for _, val := range data.imports {
		content += val + ";\n"
	}

	trash := p.HasDeletedAt()

	content += "\n"
	content += fmt.Sprintf("service %s{\n\n", pkg.ToUpperHump(p.Table.Module))
	for _, val := range data.sort {
		if !trash && strings.Contains(val, "Trash") {
			continue
		}
		content += data.bucket[val] + "\n"
	}
	content += "}"
	return content
}

// GenService 生成error proto
func (p *Service) GenService() (string, error) {
	// 扫描已经生成的proto
	scanService, err := p.ScanService()
	if err != nil {
		return "", fmt.Errorf("扫描proto代码失败，%s", err.Error())
	}

	// 生成新的proto
	makeService, err := p.MakeService()
	if err != nil {
		return "", fmt.Errorf("生成proto代码失败，%s", err.Error())
	}

	// 合并proto
	code := &ServiceCode{bucket: make(map[string]string), pkg: makeService.pkg}
	code.sort = append(makeService.sort, scanService.sort...)
	code.imports = append(makeService.imports, scanService.imports...)
	code.options = append(makeService.options, scanService.options...)

	for key, val := range scanService.bucket {
		code.bucket[key] = val
	}
	for key, val := range makeService.bucket {
		code.bucket[key] = val
	}

	// 去重
	code.sort = pkg.UniqueArray(code.sort)
	code.imports = pkg.UniqueArray(code.imports)
	code.options = pkg.UniqueArray(code.options)

	// 生成proto
	return p.RenderService(code), nil
}
