package gocode

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"
	"text/template"

	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen/types"
)

type bySrvCodeType struct {
	Fields string
	Params string
	Where  string
	Method string
}

type Service struct {
	*gen.Builder
}

func NewServiceBuilder(b *gen.Builder) *Service {
	return &Service{Builder: b}
}

type ServiceCode struct {
	pkg     string
	imports []string
	sort    []string
	bucket  map[string]string
}

func (p *Service) handleServiceFuncType(name string, fn *ast.FuncType, fset *token.FileSet) (string, error) {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, fn); err != nil {
		return "", err
	}
	signature := strings.Trim(buf.String(), " ")
	signature = strings.ReplaceAll(signature, "func", name)
	return signature, nil
}

// ParseServiceByContent 解析service文本，生成service结构
func (p *Service) ParseServiceByContent(content string) (*ServiceCode, error) {
	nodeToString := func(fset *token.FileSet, node ast.Node) (string, error) {
		var buf bytes.Buffer
		if err := format.Node(&buf, fset, node); err != nil {
			return "", err
		}
		return buf.String(), nil
	}

	doc := &ServiceCode{bucket: make(map[string]string)}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", content, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	doc.pkg = f.Name.Name
	for _, imp := range f.Imports {
		if imp.Name != nil {
			doc.imports = append(doc.imports, imp.Name.Name+" "+imp.Path.Value)
		} else {
			doc.imports = append(doc.imports, imp.Path.Value)
		}
	}

	for _, d := range f.Decls {
		switch decl := d.(type) {
		case *ast.FuncDecl:
			body, err := nodeToString(fset, decl)
			if err != nil {
				continue
			}
			funcName := decl.Name.Name
			doc.sort = append(doc.sort, funcName)
			doc.bucket[funcName] = body

		case *ast.GenDecl:
			if decl.Tok == token.TYPE { // Handle type declarations
				for _, spec := range decl.Specs {
					typeSpec, ok := spec.(*ast.TypeSpec)
					if !ok {
						continue
					}
					body, err := nodeToString(fset, decl)
					if err != nil {
						continue
					}

					typeName := typeSpec.Name.Name
					doc.sort = append(doc.sort, typeName)
					doc.bucket[typeName] = body
				}
			}
		}
	}

	return doc, nil
}

// ScanService 扫描现在的已经存在的service error定义
func (p *Service) ScanService() (*ServiceCode, error) {
	path := p.GoServicePath()
	if !pkg.IsExistFile(path) {
		return &ServiceCode{}, nil
	}
	code, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return p.ParseServiceByContent(string(code))
}

// makeGetByCodes 生成ByXxx的代码
func (p *Service) makeGetByCodes(object *types.Table) []bySrvCodeType {
	var bucket = make(map[string]*types.Column)
	for _, column := range object.Columns {
		bucket[column.Name] = column
	}

	var byCodes []bySrvCodeType
	for _, index := range object.Indexes {
		if !index.Unique {
			continue
		}
		var (
			keys   []string
			params []string
			where  []string
			fields []string
		)
		for _, name := range index.Names {
			if name == "deleted_at" {
				continue
			}

			column, ok := bucket[name]
			if !ok {
				continue
			}
			lowKey := pkg.ToLowerHump(column.Name)
			params = append(params, fmt.Sprintf("%s %s", lowKey, column.GoType()))
			keys = append(keys, pkg.ToUpperHump(column.Name))
			where = append(where, fmt.Sprintf(`Where("%s = ?",%s)`, lowKey, lowKey))
			fields = append(fields, lowKey)
		}
		if len(keys) != 0 {
			byCodes = append(byCodes, bySrvCodeType{
				Params: strings.Join(params, ","),
				Method: strings.Join(keys, "And"),
				Where:  strings.Join(where, "."),
				Fields: strings.Join(fields, ","),
			})
		}
	}
	return byCodes
}

// makeListFields 生成list字段
func (p *Service) makeListFields(object *types.Table) []string {
	var (
		length int
		list   []string
	)

	for _, column := range object.Columns {
		if !column.Operation.List {
			continue
		}
		if column.IsDeletedAt() {
			continue
		}
		length++
		list = append(list, fmt.Sprintf(`"%s"`, pkg.ToSnake(column.Name)))
	}
	if len(list) == length {
		list = []string{"\"*\""}
	}
	return list
}

// MakeService 根据模板生成service
func (p *Service) MakeService() (*ServiceCode, error) {
	code, err := os.ReadFile(p.GoServiceTplPath())
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	tmpl, err := template.New("Service").Parse(strings.TrimSpace(string(code)))
	if err != nil {
		return nil, err
	}

	renderData := map[string]any{
		"EnableBatchDelete": p.Table.EnableBatchDelete,
		"GetByCodes":        p.makeGetByCodes(p.Table),
		"Module":            p.Module,
		"Server":            p.Server,
		"Object":            pkg.ToUpperHump(p.Table.Struct),
		"Title":             p.Table.Comment,
		"Classify":          pkg.ToUpperHump(p.Table.Module),
		"IsTree":            p.Table.Type == types.TableTypeTree,
	}

	if err := tmpl.Execute(buf, renderData); err != nil {
		return nil, err
	}

	codeStr := p.FormatGoCode(buf.String())

	return p.ParseServiceByContent(codeStr)
}

func (p *Service) RenderService(service *ServiceCode) string {
	var code = "package service\n"
	if len(service.imports) != 0 {
		code += fmt.Sprintf("import (\n%s\n)\n", strings.Join(service.imports, "\n"))
	}

	var (
		lines []string
		trash = p.HasDeletedAt()
	)
	for _, item := range service.sort {
		if !trash && strings.Contains(item, "Trash") {
			continue
		}
		lines = append(lines, service.bucket[item])
	}

	code += strings.Join(lines, "\n")

	return p.FormatGoCode(code)
}

// GenService 生成error service
func (p *Service) GenService() (string, error) {
	// 扫描已经生成的service
	scanService, err := p.ScanService()
	if err != nil {
		return "", fmt.Errorf("扫描service代码失败，%s", err.Error())
	}

	// 生成新的service
	makeService, err := p.MakeService()
	if err != nil {
		return "", fmt.Errorf("生成service代码失败，%s", err.Error())
	}

	// 合并service
	code := &ServiceCode{bucket: make(map[string]string)}
	code.sort = append(makeService.sort, scanService.sort...)
	code.imports = append(makeService.imports, scanService.imports...)
	for key, val := range scanService.bucket {
		code.bucket[key] = val
	}
	for key, val := range makeService.bucket {
		code.bucket[key] = val
	}

	// 去重
	code.sort = pkg.UniqueArray(code.sort)
	code.imports = pkg.UniqueArray(code.imports)

	// 生成service
	return p.RenderService(code), nil
}
