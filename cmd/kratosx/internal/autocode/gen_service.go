package autocode

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"strings"
	"text/template"

	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/base"
)

var (
	srvServicePath  = base.KratosxCliMod() + "/internal/autocode/template/service/service.tpl"
	srvEntrancePath = base.KratosxCliMod() + "/internal/autocode/template/service/entrance.tpl"
)

type service struct {
}

type serverSrv struct {
	Package string
	Imports []string
	Sort    []string
	Map     map[string]string
}

func GenService(object *Object) (map[string]string, error) {
	b := &service{}
	reply := map[string]string{}

	srvCode, err := b.renderSrv(object)
	if err != nil {
		return nil, err
	}
	reply[b.srvPath(object)] = srvCode

	entCode, err := b.renderEntrance(object)
	if err != nil {
		return nil, err
	}
	reply[b.entPath()] = entCode
	return reply, nil
}

func (b *service) dir() string {
	return strings.ToLower(fmt.Sprintf("internal/service"))
}

func (b *service) srvPath(object *Object) string {
	return b.dir() + "/" + toLowerCamelCase(object.Module) + ".go"
}

func (b *service) entPath() string {
	return b.dir() + "/entrance.go"
}

func (b *service) genSrvTplVariable(object *Object) map[string]any {
	return map[string]any{
		"Server":      object.Server,
		"ServerName":  object.ServerName(),
		"Module":      toLowerCase(object.Module),
		"ModuleUpper": toUpperCamelCase(object.Module),
		"Object":      toUpperCamelCase(object.Keyword),
		"Title":       object.Comment,
	}
}

func (b *service) genSrv(object *Object) (*serverSrv, error) {
	oldSrv := &serverSrv{Map: make(map[string]string)}
	byteData, err := os.ReadFile(b.srvPath(object))
	if err == nil {
		res, err := b.scanSrv(string(byteData))
		if err != nil {
			return nil, err
		}
		oldSrv = res

	}

	tp, err := os.ReadFile(srvServicePath)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("go").Parse(strings.TrimSpace(string(tp)))
	if err != nil {
		return nil, err
	}

	if err := tmpl.Execute(buf, b.genSrvTplVariable(object)); err != nil {
		return nil, err
	}

	newSrv, err := b.scanSrv(buf.String())
	if err != nil {
		return nil, err
	}
	oldSrv.Package = newSrv.Package
	oldSrv.Imports = append(oldSrv.Imports, newSrv.Imports...)
	oldSrv.Sort = append(oldSrv.Sort, newSrv.Sort...)
	for key, val := range newSrv.Map {
		if oldVal := oldSrv.Map[key]; strings.Contains(oldVal, _fixedCode) {
			continue
		}
		oldSrv.Map[key] = val
	}

	oldSrv.Imports = uniqueStrings(oldSrv.Imports)
	oldSrv.Sort = uniqueStrings(oldSrv.Sort)
	return oldSrv, nil
}

func (b *service) scanSrv(src string) (*serverSrv, error) {
	nodeToString := func(fset *token.FileSet, node ast.Node) (string, error) {
		var buf bytes.Buffer
		if err := format.Node(&buf, fset, node); err != nil {
			return "", err
		}
		return buf.String(), nil
	}

	doc := &serverSrv{Map: make(map[string]string)}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	doc.Package = f.Name.Name
	for _, imp := range f.Imports {
		if imp.Name != nil {
			doc.Imports = append(doc.Imports, imp.Name.Name+" "+imp.Path.Value)
		} else {
			doc.Imports = append(doc.Imports, imp.Path.Value)
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
			doc.Sort = append(doc.Sort, funcName)
			doc.Map[funcName] = body

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
					doc.Sort = append(doc.Sort, typeName)
					doc.Map[typeName] = body
				}
			}
		}
	}

	return doc, nil
}

func (b *service) renderSrv(object *Object) (string, error) {
	tps, err := b.genSrv(object)
	if err != nil {
		return "", err
	}

	var (
		sb strings.Builder
		md = object.MethodStatus()
	)
	// Write the package statement
	fmt.Fprintf(&sb, "package %s\n\n", tps.Package)

	// Write the imports if any
	if len(tps.Imports) > 0 {
		sb.WriteString("import (\n")
		for _, imp := range tps.Imports {
			fmt.Fprintf(&sb, "\t%s\n", imp)
		}
		sb.WriteString(")\n\n")
	}

	// Write each type definition in the specified order
	for _, typeName := range tps.Sort {
		typeDef, ok := tps.Map[typeName]
		if !ok {
			continue // Skip if there is no definition for the type name
		}
		if !object.HasMethod(md, typeName) {
			continue // Skip if there is no definition for the function name
		}

		// Write the type definition including comments
		sb.WriteString(typeDef)
		sb.WriteString("\n\n") // Add an extra line after each type for readability
	}

	formattedCode, err := format.Source([]byte(sb.String()))
	if err != nil {
		return sb.String(), nil // Return the error if the code could not be formatted
	}

	return string(formattedCode), nil
}

func (b *service) renderEntrance(object *Object) (string, error) {
	tp, err := os.ReadFile(srvEntrancePath)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	tmpl, err := template.New("go").Parse(strings.TrimSpace(string(tp)))
	if err != nil {
		return "", err
	}
	if err := tmpl.Execute(buf, b.genSrvTplVariable(object)); err != nil {
		return "", err
	}

	formattedCode, err := format.Source(buf.Bytes())
	if err != nil {
		return buf.String(), nil // Return the error if the code could not be formatted
	}

	return string(formattedCode), nil
}
