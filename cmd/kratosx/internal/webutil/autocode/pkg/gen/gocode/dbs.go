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

type byDbsCodeType struct {
	Params string
	Where  string
	Method string
}

type Dbs struct {
	*gen.Builder
}

func NewDbsBuilder(b *gen.Builder) *Dbs {
	return &Dbs{Builder: b}
}

type DbsCode struct {
	pkg     string
	imports []string
	sort    []string
	bucket  map[string]string
}

func (p *Dbs) handleDbsFuncType(name string, fn *ast.FuncType, fset *token.FileSet) (string, error) {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, fn); err != nil {
		return "", err
	}
	signature := strings.Trim(buf.String(), " ")
	signature = strings.ReplaceAll(signature, "func", name)
	return signature, nil
}

// ParseDbsByContent 解析dbs文本，生成dbs结构
func (p *Dbs) ParseDbsByContent(content string) (*DbsCode, error) {
	nodeToString := func(fset *token.FileSet, node ast.Node) (string, error) {
		var buf bytes.Buffer
		if err := format.Node(&buf, fset, node); err != nil {
			return "", err
		}
		return buf.String(), nil
	}

	doc := &DbsCode{bucket: make(map[string]string)}

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

// ScanDbs 扫描现在的已经存在的dbs error定义
func (p *Dbs) ScanDbs() (*DbsCode, error) {
	path := p.GoDbsPath()
	if !pkg.IsExistFile(path) {
		return &DbsCode{}, nil
	}
	code, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return p.ParseDbsByContent(string(code))
}

// makeGetByCodes 生成ByXxx的代码
func (p *Dbs) makeGetByCodes(object *types.Table) []byDbsCodeType {
	var bucket = make(map[string]*types.Column)
	for _, column := range object.Columns {
		bucket[column.Name] = column
	}

	var byCodes []byDbsCodeType
	for _, index := range object.Indexes {
		if !index.Unique {
			continue
		}
		var (
			keys   []string
			params []string
			where  []string
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
		}
		if len(keys) != 0 {
			byCodes = append(byCodes, byDbsCodeType{
				Params: strings.Join(params, ","),
				Method: strings.Join(keys, "And"),
				Where:  strings.Join(where, "."),
			})
		}
	}
	return byCodes
}

// makeGetFields 生成查询字段
func (p *Dbs) makeGetFields(object *types.Table) []string {
	var (
		length int
		list   []string
	)

	for _, column := range object.Columns {
		if !column.Operation.Get {
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

// makeListFields 生成list字段
func (p *Dbs) makeListFields(object *types.Table) []string {
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

// makeGetTrashFields 生成查询字段
func (p *Dbs) makeGetTrashFields(object *types.Table) []string {
	list := p.makeGetFields(object)
	if len(list) == 1 && list[0] == "\"*\"" {
		return list
	}
	list = append(list, "deleted_at")
	return list
}

// makeListTrashFields 生成list字段
func (p *Dbs) makeListTrashFields(object *types.Table) []string {
	list := p.makeListFields(object)
	if len(list) == 1 && list[0] == "\"*\"" {
		return list
	}
	list = append(list, "deleted_at")
	return list
}

func (p *Dbs) getPreload(rela *types.Relation, tp string, prefix string) []string {
	var preload []string
	var innerRelas []*types.Relation
	for _, column := range rela.Table.Columns {
		if tp == "GET" && !column.Operation.Get {
			continue
		}
		if tp == "LIST" && !column.Operation.List {
			continue
		}
		relaKey := pkg.ToUpperHump(rela.Table.Name)
		if rela.Type == types.RelationTypeMany {
			relaKey = pkg.ToPluralize(relaKey)
		}

		preload = append(preload, fmt.Sprintf(`Preload("%s")`, prefix+relaKey))
		if column.Relations != nil {
			innerRelas = append(innerRelas, column.Relations...)
		}
	}

	for _, item := range innerRelas {
		relaKey := pkg.ToUpperHump(rela.Table.Name)
		if rela.Type == types.RelationTypeMany {
			relaKey = pkg.ToPluralize(relaKey)
		}

		prefix = prefix + relaKey + "."
		preload = append(preload, p.getPreload(item, tp, prefix)...)
	}
	return pkg.UniqueArray(preload)
}

// makePreload 生成预加载数据
func (p *Dbs) makePreload(object *types.Table, tp string) []string {
	var (
		list []string
	)

	for _, column := range object.Columns {
		if !column.Operation.Get {
			continue
		}
		for _, relation := range column.Relations {
			tmp := *relation
			list = append(list, p.getPreload(&tmp, tp, "")...)
		}
	}
	return list
}

func (p *Dbs) makeListQuery(object *types.Table) []string {
	var (
		list []string
	)

	for _, column := range object.Columns {
		if column.Query.Type == "" {
			continue
		}

		upperName := pkg.ToUpperHump(column.Name)
		snakeName := pkg.ToSnake(column.Name)
		pluralizeName := pkg.ToPluralize(upperName)
		switch strings.ToLower(column.Query.Type) {
		case "in":
			tpl := `if req.%s != nil {
								db = db.Where("%s IN ?", req.%s)
							}`
			code := fmt.Sprintf(tpl, pluralizeName, snakeName, pluralizeName)
			list = append(list, code)

		case "not in":
			tpl := `if req.%s != nil {
								db = db.Where("%s NOT IN ?", req.%s)
							}`
			code := fmt.Sprintf(tpl, pluralizeName, snakeName, pluralizeName)
			list = append(list, code)

		case "between":
			tpl := `if len(req.%s) == 2 {
								db = db.Where("%s BETWEEN ? AND ?", req.%s[0], req.%s[1])
							}`
			code := fmt.Sprintf(tpl, pluralizeName, snakeName, pluralizeName, pluralizeName)
			list = append(list, code)

		case "like":
			tpl := `if req.%s != nil {
								db = db.Where("%s LIKE ?", *req.%s+"%%")
							}`
			code := fmt.Sprintf(tpl, upperName, snakeName, upperName)
			list = append(list, code)

		default:
			tpl := `if req.%s != nil {
								db = db.Where("%s %s ?", *req.%s)
							}`
			code := fmt.Sprintf(tpl, upperName, snakeName, column.Query.Type, upperName)
			list = append(list, code)
		}
	}
	return list
}

// MakeDbs 根据模板生成dbs
func (p *Dbs) MakeDbs() (*DbsCode, error) {
	code, err := os.ReadFile(p.GoDbsTplPath())
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	tmpl, err := template.New("Dbs").Parse(strings.TrimSpace(string(code)))
	if err != nil {
		return nil, err
	}

	getPreloads := p.makePreload(p.Table, "GET")
	listPreloads := p.makePreload(p.Table, "LIST")

	renderData := map[string]any{
		"EnableBatchDelete": p.Table.EnableBatchDelete,
		"GetByCodes":        p.makeGetByCodes(p.Table),
		"GetFields":         strings.Join(p.makeGetFields(p.Table), ", "),
		"ListFields":        strings.Join(p.makeListFields(p.Table), ", "),
		"GetTrashFields":    strings.Join(p.makeGetTrashFields(p.Table), ", "),
		"ListTrashFields":   strings.Join(p.makeListTrashFields(p.Table), ", "),
		"HasGetPreload":     len(getPreloads) != 0,
		"GetPreload":        strings.Join(getPreloads, "."),
		"HasListPreload":    len(listPreloads) != 0,
		"ListPreload":       strings.Join(listPreloads, "."),
		"QueryCodes":        strings.Join(p.makeListQuery(p.Table), "\n"),
		"Module":            p.Module,
		"Object":            pkg.ToUpperHump(p.Table.Struct),
		"Title":             p.Table.Comment,
		"Classify":          pkg.ToUpperHump(p.Table.Module),
		"IsTree":            p.Table.Type == types.TableTypeTree,
	}

	if err := tmpl.Execute(buf, renderData); err != nil {
		return nil, err
	}

	codeStr := p.FormatGoCode(buf.String())

	return p.ParseDbsByContent(codeStr)
}

func (p *Dbs) RenderDbs(dbs *DbsCode) string {
	var code = "package dbs\n"
	if len(dbs.imports) != 0 {
		code += fmt.Sprintf("import (\n%s\n)\n", strings.Join(dbs.imports, "\n"))
	}

	var (
		lines []string
		trash = p.HasDeletedAt()
	)
	for _, item := range dbs.sort {
		if !trash && strings.Contains(item, "Trash") {
			continue
		}
		lines = append(lines, dbs.bucket[item])
	}

	code += strings.Join(lines, "\n")

	return p.FormatGoCode(code)
}

// GenDbs 生成error dbs
func (p *Dbs) GenDbs() (string, error) {
	// 扫描已经生成的dbs
	scanDbs, err := p.ScanDbs()
	if err != nil {
		return "", fmt.Errorf("扫描dbs代码失败，%s", err.Error())
	}

	// 生成新的dbs
	makeDbs, err := p.MakeDbs()
	if err != nil {
		return "", fmt.Errorf("生成dbs代码失败，%s", err.Error())
	}

	// 合并dbs
	code := &DbsCode{bucket: make(map[string]string)}
	code.sort = append(makeDbs.sort, scanDbs.sort...)
	code.imports = append(makeDbs.imports, scanDbs.imports...)
	for key, val := range scanDbs.bucket {
		code.bucket[key] = val
	}
	for key, val := range makeDbs.bucket {
		code.bucket[key] = val
	}

	// 去重
	code.sort = pkg.UniqueArray(code.sort)
	code.imports = pkg.UniqueArray(code.imports)

	// 生成dbs
	return p.RenderDbs(code), nil
}
