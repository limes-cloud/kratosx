package gocode

import (
	"bytes"
	"fmt"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen/types"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"
	"text/template"
)

type Entity struct {
	*gen.Builder
}

func NewEntityBuilder(b *gen.Builder) *Entity {
	return &Entity{Builder: b}
}

type EntityCode struct {
	imports []string
	sort    []string
	bucket  map[string]string
}

func (p *Entity) handleEntityFuncType(name string, fn *ast.FuncType, fset *token.FileSet) (string, error) {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, fn); err != nil {
		return "", err
	}
	signature := strings.Trim(buf.String(), " ")
	signature = strings.ReplaceAll(signature, "func", name)
	return signature, nil
}

// ParseEntityByContent 解析entity文本，生成entity结构
func (p *Entity) ParseEntityByContent(content string) (*EntityCode, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", content, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	tps := &EntityCode{
		bucket: make(map[string]string),
	}

	// 解析imports
	for _, imp := range f.Imports {
		if imp.Name != nil {
			tps.imports = append(tps.imports, imp.Name.Name+" "+imp.Path.Value)
		} else {
			tps.imports = append(tps.imports, imp.Path.Value)
		}
	}

	// 解析定义
	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		if genDecl.Tok == token.TYPE { // We are only interested in type declarations
			for _, spec := range genDecl.Specs {
				typeSpec := spec.(*ast.TypeSpec)
				var buf bytes.Buffer
				if err := printer.Fprint(&buf, fset, typeSpec); err != nil {
					return nil, err
				}
				typeDef := buf.String()

				// Check for comments associated with the type
				var comments strings.Builder
				if genDecl.Doc != nil {
					for _, comment := range genDecl.Doc.List {
						comments.WriteString(comment.Text)
						comments.WriteString("\n")
					}
				}

				fullTypeDef := comments.String() + "type " + typeDef
				tps.sort = append(tps.sort, typeSpec.Name.Name)
				tps.bucket[typeSpec.Name.Name] = fullTypeDef
			}
		}
	}

	return tps, nil
}

// ScanEntity 扫描现在的已经存在的entity error定义
func (p *Entity) ScanEntity() (*EntityCode, error) {
	path := p.GoEntityPath()
	if !pkg.IsExistFile(path) {
		return &EntityCode{}, nil
	}
	code, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return p.ParseEntityByContent(string(code))
}

func (p *Entity) genEntityStruct(object *types.Table) map[string]string {
	var (
		fields          []string
		relations       []string
		relationObjects []*types.Table
		entities        = make(map[string]string)
	)
	array := pkg.New([]string{"id", "created_at", "updated_at", "deleted_at"})
	count := 0
	for _, column := range object.Columns {
		opt := ""
		if column.IsProtoOption() {
			opt = "*"
		}
		if array.Has(column.Name) {
			count++
			continue
		}
		fields = append(fields, fmt.Sprintf("\t%s %s%s `gorm:\"column:%s\" json:\"%s\"` //%s",
			pkg.ToUpperHump(column.Name),
			opt,
			column.GoType(),
			pkg.ToSnake(column.Name),
			pkg.VariableName(column.Name, p.NameRule),
			column.Comment,
		))

		for _, item := range column.Relations {
			obj := item.Table
			relationObjects = append(relationObjects, obj)
			if item.Type == types.RelationTypeOne {
				relations = append(relations, fmt.Sprintf("\t%s *%s `gorm:\"foreignKey:%s;references:%s\" json:\"%s\"` // %s",
					pkg.ToUpperHump(obj.Name),
					pkg.ToUpperHump(obj.Name),
					pkg.ToSnake(item.Column),
					pkg.ToSnake(column.Name),
					pkg.VariableName(obj.Name, p.NameRule),
					obj.Comment,
				))
			} else {
				relations = append(relations, fmt.Sprintf("\t%s []*%s `gorm:\"foreignKey:%s;references:%s\" json:\"%s\"` // %s",
					pkg.ToPluralize(pkg.ToUpperHump(obj.Name)),
					pkg.ToUpperHump(obj.Name),
					pkg.ToSnake(item.Column),
					pkg.ToSnake(column.Name),
					pkg.ToPluralize(pkg.VariableName(obj.Name, p.NameRule)),
					obj.Comment,
				))
			}
		}
	}
	fields = append(fields, relations...)
	if object.Type == types.TableTypeTree {
		fields = append(fields, fmt.Sprintf("\tChildren []*%s `gorm:\"-\" json:\"children\"` // 子节点",
			pkg.ToUpperHump(object.Struct),
		))
	}

	// 变换模型
	switch count {
	case 1:
		fields = append([]string{"Id  uint32 `json:\"id\" gorm:\"primaryKey;autoIncrement\"` // 主键ID"}, fields...)
	case 2:
		fields = append([]string{"types.CreateModel"}, fields...)
	case 3:
		fields = append([]string{"types.BaseModel"}, fields...)
	case 4:
		fields = append([]string{"types.DeleteModel"}, fields...)
	}

	entities[pkg.ToUpperHump(object.Struct)] = strings.Join(fields, "\n")

	//for _, rela := range relationObjects {
	//	cents := p.genEntityStruct(rela)
	//	for key, val := range cents {
	//		entities[key] = val
	//	}
	//}
	return entities
}

// MakeEntity 根据模板生成entity
func (p *Entity) MakeEntity() (*EntityCode, error) {
	code, err := os.ReadFile(p.GoEntityTplPath())
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	tmpl, err := template.New("Entity").Parse(strings.TrimSpace(string(code)))
	if err != nil {
		return nil, err
	}
	renderData := map[string]any{
		"Entities": p.genEntityStruct(p.Table),
		"Module":   p.Module,
		"Object":   pkg.ToUpperHump(p.Table.Struct),
		"Title":    p.Table.Comment,
	}
	if err := tmpl.Execute(buf, renderData); err != nil {
		return nil, err
	}

	return p.ParseEntityByContent(buf.String())
}

func (p *Entity) RenderEntity(entity *EntityCode) string {
	var code = "package entity\n"
	if len(entity.imports) != 0 {
		code += fmt.Sprintf("import (\n%s\n)\n", strings.Join(entity.imports, "\n"))
	}

	var lines []string
	for _, item := range entity.sort {
		lines = append(lines, entity.bucket[item])
	}

	code += strings.Join(lines, "\n")

	return p.FormatGoCode(code)
}

// GenEntity 生成error entity
func (p *Entity) GenEntity() (string, error) {
	// 扫描已经生成的entity
	scanEntity, err := p.ScanEntity()
	if err != nil {
		return "", fmt.Errorf("扫描entity代码失败，%s", err.Error())
	}

	// 生成新的entity
	makeEntity, err := p.MakeEntity()
	if err != nil {
		return "", fmt.Errorf("生成entity代码失败，%s", err.Error())
	}

	// 合并entity
	code := &EntityCode{bucket: make(map[string]string)}
	code.sort = append(makeEntity.sort, scanEntity.sort...)
	code.imports = append(makeEntity.imports, scanEntity.imports...)
	for key, val := range scanEntity.bucket {
		code.bucket[key] = val
	}
	for key, val := range makeEntity.bucket {
		code.bucket[key] = val
	}

	// 去重
	code.sort = pkg.UniqueArray(code.sort)
	code.imports = pkg.UniqueArray(code.imports)

	// 生成entity
	return p.RenderEntity(code), nil
}
