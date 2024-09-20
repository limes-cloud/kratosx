package gocode

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"

	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen/types"
)

type Types struct {
	*gen.Builder
}

func NewTypesBuilder(b *gen.Builder) *Types {
	return &Types{Builder: b}
}

type TypesCode struct {
	imports []string
	sort    []string
	bucket  map[string]string
}

// ParseTypesByContent 解析types文本，生成types结构
func (p *Types) ParseTypesByContent(content string) (*TypesCode, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", content, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	tps := &TypesCode{
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

// ScanTypes 扫描现在的已经存在的types error定义
func (p *Types) ScanTypes() (*TypesCode, error) {
	path := p.GoTypesPath()
	if !pkg.IsExistFile(path) {
		return &TypesCode{}, nil
	}
	code, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return p.ParseTypesByContent(string(code))
}

func (p *Types) makeListTypes() (*TypesCode, error) {
	var rows []string

	// 列表页添加分页
	if p.Table.Type == types.TableTypeList {
		rows = append(rows, fmt.Sprintf("\tPage uint32 `json:\"page\"`"))
		rows = append(rows, fmt.Sprintf("\tPageSize uint32 `json:\"%s\"`", pkg.VariableName("PageSize", p.NameRule)))
	}

	// 添加排序
	rows = append(rows, "\tOrder *string `json:\"order\"`")
	rows = append(rows, fmt.Sprintf("\tOrderBy *string `json:\"%s\"`", pkg.VariableName("OrderBy", p.NameRule)))

	// 添加查询字段
	for _, column := range p.Table.Columns {
		if column.Query.Type == "" {
			continue
		}
		if column.Query.IsPluralize() {
			rows = append(rows, fmt.Sprintf("\t%s []%s `json:\"%s\"`",
				pkg.ToPluralize(pkg.ToUpperHump(column.Name)),
				column.GoType(),
				pkg.ToPluralize(pkg.VariableName(column.Name, p.NameRule)),
			))
		} else {
			rows = append(rows, fmt.Sprintf("\t%s *%s `json:\"%s\"`",
				pkg.ToUpperHump(column.Name),
				column.GoType(),
				pkg.VariableName(column.Name, p.NameRule),
			))
		}
	}
	key := fmt.Sprintf("List%sRequest", pkg.ToUpperHump(p.Table.Struct))
	code := fmt.Sprintf("type %s struct{\n%s\n}", key, strings.Join(rows, "\n"))
	return &TypesCode{
		sort: []string{key},
		bucket: map[string]string{
			key: code,
		},
	}, nil
}

func (p *Types) makeTrashListTypes() (*TypesCode, error) {
	code, err := p.makeListTypes()
	if err != nil {
		return nil, err
	}
	name := code.sort[0]
	code.sort[0] = fmt.Sprintf("ListTrash%sRequest", pkg.ToUpperHump(p.Table.Struct))
	code.bucket[name] = strings.ReplaceAll(code.bucket[name], name, code.sort[0])
	code.bucket[code.sort[0]] = code.bucket[name]
	delete(code.bucket, name)
	return code, nil
}

// MakeTypes 根据模板生成types
func (p *Types) MakeTypes() (*TypesCode, error) {
	listCode, err := p.makeListTypes()
	if err != nil {
		return nil, err
	}

	listTrashCode, err := p.makeTrashListTypes()
	if err != nil {
		return nil, err
	}

	listCode.sort = append(listCode.sort, listTrashCode.sort...)
	for ind, item := range listTrashCode.bucket {
		listCode.bucket[ind] = item
	}

	return listCode, nil
}

func (p *Types) RenderTypes(types *TypesCode) string {
	var code = "package types\n"
	if len(types.imports) != 0 {
		code += fmt.Sprintf("import (\n%s\n)\n", strings.Join(types.imports, "\n"))
	}

	trash := p.HasDeletedAt()
	for _, item := range types.sort {
		if !trash && strings.Contains(item, "Trash") {
			continue
		}
		code += types.bucket[item] + "\n"
	}

	return p.FormatGoCode(code)
}

// GenTypes 生成error types
func (p *Types) GenTypes() (string, error) {
	// 扫描已经生成的types
	scanTypes, err := p.ScanTypes()
	if err != nil {
		return "", fmt.Errorf("扫描types代码失败，%s", err.Error())
	}

	// 生成新的types
	makeTypes, err := p.MakeTypes()
	if err != nil {
		return "", fmt.Errorf("生成types代码失败，%s", err.Error())
	}

	// 合并types
	code := &TypesCode{bucket: make(map[string]string)}
	code.sort = append(makeTypes.sort, scanTypes.sort...)
	for key, val := range scanTypes.bucket {
		code.bucket[key] = val
	}
	for key, val := range makeTypes.bucket {
		code.bucket[key] = val
	}

	// 去重
	code.sort = pkg.UniqueArray(code.sort)

	// 生成types
	return p.RenderTypes(code), nil
}
