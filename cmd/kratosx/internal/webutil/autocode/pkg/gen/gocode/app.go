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

type byAppCodeType struct {
	Params    string
	Where     string
	Method    string
	Fields    []string
	Composite bool
}

type App struct {
	*gen.Builder
}

func NewAppBuilder(b *gen.Builder) *App {
	return &App{Builder: b}
}

type AppCode struct {
	pkg     string
	imports []string
	sort    []string
	bucket  map[string]string
}

func (p *App) handleAppFuncType(name string, fn *ast.FuncType, fset *token.FileSet) (string, error) {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, fn); err != nil {
		return "", err
	}
	signature := strings.Trim(buf.String(), " ")
	signature = strings.ReplaceAll(signature, "func", name)
	return signature, nil
}

// ParseAppByContent 解析dbs文本，生成dbs结构
func (p *App) ParseAppByContent(content string) (*AppCode, error) {
	nodeToString := func(fset *token.FileSet, node ast.Node) (string, error) {
		var buf bytes.Buffer
		if err := format.Node(&buf, fset, node); err != nil {
			return "", err
		}
		return buf.String(), nil
	}

	doc := &AppCode{bucket: make(map[string]string)}

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

// ScanApp 扫描现在的已经存在的app error定义
func (p *App) ScanApp() (*AppCode, error) {
	path := p.GoAppPath()
	if !pkg.IsExistFile(path) {
		return &AppCode{}, nil
	}
	code, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return p.ParseAppByContent(string(code))
}

// makeGetByCodes 生成ByXxx的代码
func (p *App) makeGetByCodes(object *types.Table) []byAppCodeType {
	var bucket = make(map[string]*types.Column)
	for _, column := range object.Columns {
		bucket[column.Name] = column
	}

	var byCodes []byAppCodeType
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
			fields = append(fields, pkg.ToUpperHump(column.Name))
		}
		if len(keys) != 0 {
			byCodes = append(byCodes, byAppCodeType{
				Params: strings.Join(params, ","),
				Method: strings.Join(keys, "And"),
				Where:  strings.Join(where, "."),
				Fields: fields,
			})
		}
	}
	return byCodes
}

func (p *App) makeListQuery(object *types.Table) []string {
	var (
		list []string
	)

	for _, column := range object.Columns {
		if column.Query.Type == "" {
			continue
		}

		key := pkg.ToUpperHump(column.Name)
		if column.Query.IsPluralize() {
			list = append(list, pkg.ToPluralize(key))
		} else {
			list = append(list, key)
		}
	}
	return list
}

// MakeApp 根据模板生成app
func (p *App) MakeApp() (*AppCode, error) {
	code, err := os.ReadFile(p.GoAppTplPath())
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	tmpl, err := template.New("App").Parse(strings.TrimSpace(string(code)))
	if err != nil {
		return nil, err
	}

	renderData := map[string]any{
		"EnableBatchDelete": p.Table.EnableBatchDelete,
		"GetByCodes":        p.makeGetByCodes(p.Table),
		"QueryFields":       p.makeListQuery(p.Table),
		"Module":            p.Module,
		"Server":            p.Server,
		"Object":            pkg.ToUpperHump(p.Table.Struct),
		"Title":             p.Table.Comment,
		"Classify":          pkg.ToUpperHump(p.Table.Module),
		"ClassifyLowerCase": pkg.ToSnake(p.Table.Module),
		"IsTree":            p.Table.Type == types.TableTypeTree,
	}

	if err := tmpl.Execute(buf, renderData); err != nil {
		return nil, err
	}

	codeStr := p.FormatGoCode(buf.String())

	return p.ParseAppByContent(codeStr)
}

func (p *App) RenderApp(app *AppCode) string {
	var code = "package app\n"
	if len(app.imports) != 0 {
		code += fmt.Sprintf("import (\n%s\n)\n", strings.Join(app.imports, "\n"))
	}

	var (
		lines []string
		trash = p.HasDeletedAt()
	)
	for _, item := range app.sort {
		if !trash && strings.Contains(item, "Trash") {
			continue
		}
		lines = append(lines, app.bucket[item])
	}

	code += strings.Join(lines, "\n")

	return p.FormatGoCode(code)
}

// GenApp 生成error app
func (p *App) GenApp() (string, error) {
	// 扫描已经生成的app
	scanApp, err := p.ScanApp()
	if err != nil {
		return "", fmt.Errorf("扫描app代码失败，%s", err.Error())
	}

	// 生成新的app
	makeApp, err := p.MakeApp()
	if err != nil {
		return "", fmt.Errorf("生成app代码失败，%s", err.Error())
	}

	// 合并app
	code := &AppCode{bucket: make(map[string]string)}
	code.sort = append(makeApp.sort, scanApp.sort...)
	code.imports = append(makeApp.imports, scanApp.imports...)
	for key, val := range scanApp.bucket {
		code.bucket[key] = val
	}
	for key, val := range makeApp.bucket {
		code.bucket[key] = val
	}

	// 去重
	code.sort = pkg.UniqueArray(code.sort)
	code.imports = pkg.UniqueArray(code.imports)

	// 生成app
	return p.RenderApp(code), nil
}

func (p *App) GenAppEntry() (string, error) {
	code, err := os.ReadFile(p.GoAppEntryTplPath())
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	tmpl, err := template.New("AppEntry").Parse(strings.TrimSpace(string(code)))
	if err != nil {
		return "", err
	}

	renderData := map[string]any{
		"Module":            p.Module,
		"Server":            p.Server,
		"Object":            pkg.ToUpperHump(p.Table.Struct),
		"Title":             p.Table.Comment,
		"Classify":          pkg.ToUpperHump(p.Table.Module),
		"ClassifyLowerCase": pkg.ToSnake(p.Table.Module),
		"IsTree":            p.Table.Type == types.TableTypeTree,
	}

	if err := tmpl.Execute(buf, renderData); err != nil {
		return "", err
	}

	return p.FormatGoCode(buf.String()), nil
}
