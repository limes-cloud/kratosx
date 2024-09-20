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
	"text/template"

	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen/types"
)

type Repo struct {
	*gen.Builder
}

func NewRepoBuilder(b *gen.Builder) *Repo {
	return &Repo{Builder: b}
}

type byRepoCodeType struct {
	Params string
	Where  string
	Method string
}

type RepoCode struct {
	imports []string
	sort    []string
	bucket  map[string]string
}

func (b *Repo) handleRepoFuncType(name string, fn *ast.FuncType, fset *token.FileSet) (string, error) {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, fn); err != nil {
		return "", err
	}
	signature := strings.Trim(buf.String(), " ")
	signature = strings.ReplaceAll(signature, "func", name)
	return signature, nil
}

// ParseRepoByContent 解析repo文本，生成repo结构
func (p *Repo) ParseRepoByContent(content string) (*RepoCode, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", content, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	tps := &RepoCode{
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
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			interfaceType, ok := typeSpec.Type.(*ast.InterfaceType)
			if !ok {
				continue
			}
			for _, method := range interfaceType.Methods.List {
				for _, name := range method.Names {
					signature, err := p.handleRepoFuncType(name.Name, method.Type.(*ast.FuncType), fset)
					if err != nil {
						return nil, err
					}
					var comment strings.Builder
					if method.Doc != nil {
						for _, c := range method.Doc.List {
							comment.WriteString(c.Text)
							comment.WriteString("\n")
						}
					}
					fullMethod := comment.String() + signature
					tps.sort = append(tps.sort, name.Name)
					tps.bucket[name.Name] = fullMethod
				}
			}
		}
	}

	return tps, nil
}

// ScanRepo 扫描现在的已经存在的repo error定义
func (p *Repo) ScanRepo() (*RepoCode, error) {
	path := p.GoRepoPath()
	if !pkg.IsExistFile(path) {
		return &RepoCode{}, nil
	}
	code, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return p.ParseRepoByContent(string(code))
}

// makeGetByCodes 生成ByXxx的代码
func (p *Repo) makeGetByCodes(object *types.Table) []byRepoCodeType {
	var bucket = make(map[string]*types.Column)
	for _, column := range object.Columns {
		bucket[column.Name] = column
	}

	var byCodes []byRepoCodeType
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
			byCodes = append(byCodes, byRepoCodeType{
				Params: strings.Join(params, ","),
				Method: strings.Join(keys, "And"),
				Where:  strings.Join(where, "."),
			})
		}
	}
	return byCodes
}

// MakeRepo 根据模板生成repo
func (p *Repo) MakeRepo() (*RepoCode, error) {
	code, err := os.ReadFile(p.GoRepoTplPath())
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	tmpl, err := template.New("Repo").Parse(strings.TrimSpace(string(code)))
	if err != nil {
		return nil, err
	}
	renderData := map[string]any{
		"EnableBatchDelete": p.Table.EnableBatchDelete,
		"Module":            p.Module,
		"GetByCodes":        p.makeGetByCodes(p.Table),
		"Classify":          p.Table.Module,
		"ClassifyUpper":     pkg.ToUpperHump(p.Table.Module),
		"Object":            pkg.ToUpperHump(p.Table.Struct),
		"Title":             p.Table.Comment,
	}
	if err := tmpl.Execute(buf, renderData); err != nil {
		return nil, err
	}

	return p.ParseRepoByContent(buf.String())
}

func (p *Repo) RenderRepo(repo *RepoCode) string {
	var code = "package repository\n"
	if len(repo.imports) != 0 {
		code += fmt.Sprintf("import (\n%s\n)\n", strings.Join(repo.imports, "\n"))
	}

	var (
		lines []string
		trash = p.HasDeletedAt()
	)
	for _, item := range repo.sort {
		if !trash && strings.Contains(item, "Trash") {
			continue
		}
		lines = append(lines, repo.bucket[item])
	}

	code += fmt.Sprintf("type %sRepository interface{\n%s\n}", pkg.ToUpperHump(p.Table.Module), strings.Join(lines, "\n"))

	return p.FormatGoCode(code)
}

// GenRepo 生成error repo
func (p *Repo) GenRepo() (string, error) {
	// 扫描已经生成的repo
	scanRepo, err := p.ScanRepo()
	if err != nil {
		return "", fmt.Errorf("扫描repo代码失败，%s", err.Error())
	}

	// 生成新的repo
	makeRepo, err := p.MakeRepo()
	if err != nil {
		return "", fmt.Errorf("生成repo代码失败，%s", err.Error())
	}

	// 合并repo
	code := &RepoCode{bucket: make(map[string]string)}
	code.sort = append(makeRepo.sort, scanRepo.sort...)
	code.imports = append(makeRepo.imports, scanRepo.imports...)
	for key, val := range scanRepo.bucket {
		code.bucket[key] = val
	}
	for key, val := range makeRepo.bucket {
		code.bucket[key] = val
	}

	// 去重
	code.sort = pkg.UniqueArray(code.sort)
	code.imports = pkg.UniqueArray(code.imports)

	// 生成repo
	return p.RenderRepo(code), nil
}
