package autocode

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

	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/base"
)

var (
	bizRepoPath   = base.KratosxCliMod() + "/internal/autocode/template/biz/repo.tpl"
	bizTypesPath  = base.KratosxCliMod() + "/internal/autocode/template/biz/types.tpl"
	bizEntityPath = base.KratosxCliMod() + "/internal/autocode/template/biz/entity.tpl"
	bizBizPath    = base.KratosxCliMod() + "/internal/autocode/template/biz/biz.tpl"
)

type biz struct {
	mapping map[string]Mapping
}

type bizRepo struct {
	Package      string
	Imports      []string
	FunctionSort []string
	FunctionMap  map[string]string
}

type bizTypes struct {
	Package   string
	Imports   []string
	TypesSort []string
	TypesMap  map[string]string
}

type bizEntity struct {
	Package    string
	Imports    []string
	EntitySort []string
	EntityMap  map[string]string
}

type bizBiz struct {
	Package string
	Imports []string
	BizSort []string
	BizMap  map[string]string
}

func GenBiz(object *Object) (map[string]string, error) {
	b := &biz{mapping: TypesMapping()}
	reply := map[string]string{}

	repoCode, err := b.renderRepo(object)
	if err != nil {
		return nil, err
	}
	reply[b.repoPath(object)] = repoCode

	typesCode, err := b.renderTypes(object)
	if err != nil {
		return nil, err
	}
	reply[b.typesPath(object)] = typesCode

	entityCode, err := b.renderEntity(object)
	if err != nil {
		return nil, err
	}
	reply[b.entityPath(object)] = entityCode

	bizCode, err := b.renderBiz(object)
	if err != nil {
		return nil, err
	}
	reply[b.bizPath(object)] = bizCode

	return reply, nil
}

func (b *biz) dir(object *Object) string {
	return strings.ToLower(fmt.Sprintf("internal/biz/%s", object.Module))
}

func (b *biz) repoPath(object *Object) string {
	return b.dir(object) + "/repo.go"
}

func (b *biz) entityPath(object *Object) string {
	return b.dir(object) + "/entity.go"
}

func (b *biz) bizPath(object *Object) string {
	return b.dir(object) + "/biz.go"
}

func (b *biz) typesPath(object *Object) string {
	return b.dir(object) + "/types.go"
}

func (b *biz) handleRepoFuncType(name string, fn *ast.FuncType, fset *token.FileSet) (string, error) {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, fn); err != nil {
		return "", err
	}
	signature := strings.Trim(buf.String(), " ")
	signature = strings.ReplaceAll(signature, "func", name)
	return signature, nil
}

func (b *biz) genRepo(object *Object) (*bizRepo, error) {
	srv := &bizRepo{FunctionMap: make(map[string]string)}
	byteData, err := os.ReadFile(b.repoPath(object))
	if err == nil {
		res, err := b.scanRepo(string(byteData))
		if err != nil {
			return nil, err
		}
		srv = res
	}

	tp, err := os.ReadFile(bizRepoPath)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("go").Parse(strings.TrimSpace(string(tp)))
	if err != nil {
		return nil, err
	}
	renderData := map[string]any{
		"Module": toLowerCase(object.Module),
		"Object": toUpperCamelCase(object.Keyword),
		"Title":  object.Comment,
	}
	if err := tmpl.Execute(buf, renderData); err != nil {
		return nil, err
	}

	newPrv, err := b.scanRepo(buf.String())
	if err != nil {
		return nil, err
	}

	filedMap := object.FieldMap()
	for _, list := range object.Unique {
		var (
			keys    []string
			params  []string
			codeTpl = `// %s 获取指定的%s
							%s(ctx kratosx.Context, %s) (*%s, error)`
		)
		for _, item := range list {
			if toUpperCamelCase(item) == toUpperCamelCase("deleted_at") {
				continue
			}
			field, ok := filedMap[item]
			if !ok {
				continue
			}
			stp := b.mapping[field.Type].Struct
			params = append(params, fmt.Sprintf("%s %s", toLowerCamelCase(field.Keyword), stp))
			keys = append(keys, toUpperCamelCase(item))
		}
		funcName := fmt.Sprintf("Get%sBy%s", toUpperCamelCase(object.Keyword), strings.Join(keys, "And"))
		code := fmt.Sprintf(codeTpl, funcName, object.Comment, funcName, strings.Join(params, ","), toUpperCamelCase(object.Keyword))
		newPrv.FunctionSort = append(newPrv.FunctionSort, funcName)
		newPrv.FunctionMap[funcName] = code
	}

	srv.Package = newPrv.Package
	srv.Imports = append(srv.Imports, newPrv.Imports...)
	srv.FunctionSort = append(srv.FunctionSort, newPrv.FunctionSort...)
	for key, val := range newPrv.FunctionMap {
		if oldVal := srv.FunctionMap[key]; strings.Contains(oldVal, _fixedCode) {
			continue
		}
		srv.FunctionMap[key] = val
	}

	srv.Imports = uniqueStrings(srv.Imports)
	srv.FunctionSort = uniqueStrings(srv.FunctionSort)
	return srv, nil
}

func (b *biz) scanRepo(src string) (*bizRepo, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	repo := &bizRepo{
		Package:     f.Name.Name,
		Imports:     []string{},
		FunctionMap: make(map[string]string),
	}

	for _, imp := range f.Imports {
		if imp.Name != nil {
			repo.Imports = append(repo.Imports, imp.Name.Name+" "+imp.Path.Value)
		} else {
			repo.Imports = append(repo.Imports, imp.Path.Value)
		}
	}

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
					signature, err := b.handleRepoFuncType(name.Name, method.Type.(*ast.FuncType), fset)
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
					repo.FunctionSort = append(repo.FunctionSort, name.Name)
					repo.FunctionMap[name.Name] = fullMethod
				}
			}
		}
	}

	return repo, nil
}

func (b *biz) renderRepo(object *Object) (string, error) {
	repo, err := b.genRepo(object)
	if err != nil {
		return "", err
	}

	var (
		sb strings.Builder
		md = object.MethodStatus()
	)
	// Write the package statement
	fmt.Fprintf(&sb, "package %s\n\n", repo.Package)

	// Write the imports if any
	if len(repo.Imports) > 0 {
		sb.WriteString("import (\n")
		for _, imp := range repo.Imports {
			fmt.Fprintf(&sb, "\t%s\n", imp)
		}
		sb.WriteString(")\n\n")
	}

	// Begin the interface definition
	sb.WriteString("type Repo interface {\n")

	// Write each function signature in the specified order
	for _, funcName := range repo.FunctionSort {
		funcDef, ok := repo.FunctionMap[funcName]
		if !ok || !object.HasMethod(md, funcName) {
			continue // Skip if there is no definition for the function name
		}

		// Split the function definition into lines for formatting
		lines := strings.Split(funcDef, "\n")
		for i, line := range lines {
			// Write comments with proper indentation
			if strings.HasPrefix(line, "//") {
				if i > 0 {
					sb.WriteString("\n") // Add a blank line before comments for readability
				}
				fmt.Fprintf(&sb, "\t%s\n", line)
			} else {
				// Write the function signature
				fmt.Fprintf(&sb, "\t%s\n", line)
			}
		}
		sb.WriteString("\n") // Add an extra line after each function for readability
	}

	// Close the interface definition
	sb.WriteString("}\n")

	formattedCode, err := format.Source([]byte(sb.String()))
	if err != nil {
		return sb.String(), nil // Return the error if the code could not be formatted
	}

	return string(formattedCode), nil
}

func (b *biz) genTypesTplVariable(object *Object) map[string]any {
	var (
		getFields  []string
		listFields []string
		queryConds []string
	)
	if object.Type == _objectTypeList {
		listFields = append(listFields, "\tPage uint32 `json:\"page\"`")
		listFields = append(listFields, "\tPageSize uint32 `json:\"pageSize\"`")
	}
	listFields = append(listFields, "\tOrder *string `json:\"order\"`")
	listFields = append(listFields, "\tOrderBy *string `json:\"orderBy\"`")

	var unique = map[string]bool{"Id": true}
	for _, list := range object.Unique {
		for _, item := range list {
			unique[toUpperCamelCase(item)] = true
		}
	}
	delete(unique, toUpperCamelCase("deleted_at"))

	for _, field := range object.Fields {
		tp := b.mapping[field.Type].Struct
		if unique[toUpperCamelCase(field.Keyword)] {
			getFields = append(getFields, fmt.Sprintf("\t%s *%s `json:\"%s\"`",
				toUpperCamelCase(field.Keyword),
				tp,
				toLowerCamelCase(field.Keyword),
			))
		}

		if field.QueryType == "" {
			continue
		}
		switch strings.ToLower(field.QueryType) {
		case _in, _notIn, _between:
			queryConds = append(queryConds, fmt.Sprintf("\t%s []%s `json:\"%s\"`",
				pluralize(toUpperCamelCase(field.Keyword)),
				tp,
				pluralize(toLowerCamelCase(field.Keyword)),
			))
		default:
			queryConds = append(queryConds, fmt.Sprintf("\t%s *%s `json:\"%s\"`",
				toUpperCamelCase(field.Keyword),
				tp,
				toLowerCamelCase(field.Keyword),
			))
		}
	}
	listFields = append(listFields, queryConds...)
	return map[string]any{
		"GetFields":    strings.Join(getFields, "\n"),
		"ListFields":   strings.Join(listFields, "\n"),
		"ExportFields": strings.Join(queryConds, "\n"),
		"Module":       toLowerCase(object.Module),
		"Object":       toUpperCamelCase(object.Keyword),
	}
}

func (b *biz) genTypes(object *Object) (*bizTypes, error) {
	oldTypes := &bizTypes{TypesMap: make(map[string]string)}
	byteData, err := os.ReadFile(b.typesPath(object))
	if err == nil {
		res, err := b.scanTypes(string(byteData))
		if err != nil {
			return nil, err
		}
		oldTypes = res
	}

	tp, err := os.ReadFile(bizTypesPath)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("go").Parse(strings.TrimSpace(string(tp)))
	if err != nil {
		return nil, err
	}

	if err := tmpl.Execute(buf, b.genTypesTplVariable(object)); err != nil {
		return nil, err
	}

	newTypes, err := b.scanTypes(buf.String())
	if err != nil {
		return nil, err
	}
	oldTypes.Package = newTypes.Package
	oldTypes.Imports = append(oldTypes.Imports, newTypes.Imports...)
	oldTypes.TypesSort = append(oldTypes.TypesSort, newTypes.TypesSort...)
	for key, val := range newTypes.TypesMap {
		if oldVal := oldTypes.TypesMap[key]; strings.Contains(oldVal, _fixedCode) {
			continue
		}
		oldTypes.TypesMap[key] = val
	}

	oldTypes.Imports = uniqueStrings(oldTypes.Imports)
	oldTypes.TypesSort = uniqueStrings(oldTypes.TypesSort)
	return oldTypes, nil
}

func (b *biz) scanTypes(src string) (*bizTypes, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	tps := &bizTypes{
		Package:  f.Name.Name,
		Imports:  []string{},
		TypesMap: make(map[string]string),
	}

	// Collect imports
	for _, imp := range f.Imports {
		if imp.Name != nil {
			tps.Imports = append(tps.Imports, imp.Name.Name+" "+imp.Path.Value)
		} else {
			tps.Imports = append(tps.Imports, imp.Path.Value)
		}
	}

	// Collect type definitions
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
				tps.TypesSort = append(tps.TypesSort, typeSpec.Name.Name)
				tps.TypesMap[typeSpec.Name.Name] = fullTypeDef
			}
		}
	}

	return tps, nil
}

func (b *biz) renderTypes(object *Object) (string, error) {
	tps, err := b.genTypes(object)
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
	for _, typeName := range tps.TypesSort {
		typeDef, ok := tps.TypesMap[typeName]
		if !ok {
			continue // Skip if there is no definition for the type name
		}
		typeName = strings.TrimSuffix(typeName, "Request")
		typeName = strings.TrimSuffix(typeName, "Reply")
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

func (b *biz) genEntityStruct(object *Object) map[string]string {
	var (
		fields          []string
		relations       []string
		relationObjects []*Object
		ents            = make(map[string]string)
	)
	for _, field := range object.Fields {
		tp := b.mapping[field.Type].Struct
		opt := ""
		if b.isOptional(field) {
			opt = "*"
		}
		fields = append(fields, fmt.Sprintf("\t%s %s%s `json:\"%s\"`",
			toUpperCamelCase(field.Keyword),
			opt,
			tp,
			toLowerCamelCase(field.Keyword),
		))

		for _, item := range field.Relations {
			obj := item.Object
			relationObjects = append(relationObjects, obj)
			if item.Type == _relationHasOne {
				relations = append(relations, fmt.Sprintf("\t%s *%s `json:\"%s\"`",
					toUpperCamelCase(obj.Keyword),
					toUpperCamelCase(obj.Keyword),
					toLowerCamelCase(obj.Keyword),
				))
			} else {
				relations = append(relations, fmt.Sprintf("\t%s []*%s `json:\"%s\"`",
					pluralize(toUpperCamelCase(obj.Keyword)),
					toUpperCamelCase(obj.Keyword),
					pluralize(toLowerCamelCase(obj.Keyword)),
				))
			}
		}
	}
	fields = append(fields, relations...)
	if object.Type == _objectTypeTree {
		fields = append(fields, fmt.Sprintf("\tChildren []*%s `json:\"Children\"`",
			toUpperCamelCase(object.Keyword),
		))
	}
	ents[toUpperCamelCase(object.Keyword)] = strings.Join(fields, "\n")

	for _, rela := range relationObjects {
		cents := b.genEntityStruct(rela)
		for key, val := range cents {
			ents[key] = val
		}
	}
	return ents
}
func (b *biz) genEntityTplVariable(object *Object) map[string]any {
	sc := b.genEntityStruct(object)
	return map[string]any{
		"Ents":   sc,
		"Module": toLowerCase(object.Module),
		"Object": toUpperCamelCase(object.Keyword),
		"IsTree": object.Type == _objectTypeTree,
	}
}

func (b *biz) genEntity(object *Object) (*bizEntity, error) {
	oldEntity := &bizEntity{EntityMap: make(map[string]string)}
	byteData, err := os.ReadFile(b.entityPath(object))
	if err == nil {
		res, err := b.scanEntity(string(byteData))
		if err != nil {
			return nil, err
		}
		oldEntity = res

	}

	tp, err := os.ReadFile(bizEntityPath)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("go").Parse(strings.TrimSpace(string(tp)))
	if err != nil {
		return nil, err
	}

	if err := tmpl.Execute(buf, b.genEntityTplVariable(object)); err != nil {
		return nil, err
	}

	newEntity, err := b.scanEntity(buf.String())
	if err != nil {
		return nil, err
	}
	oldEntity.Package = newEntity.Package
	oldEntity.Imports = append(oldEntity.Imports, newEntity.Imports...)
	oldEntity.EntitySort = append(oldEntity.EntitySort, newEntity.EntitySort...)
	for key, val := range newEntity.EntityMap {
		if oldVal := oldEntity.EntityMap[key]; strings.Contains(oldVal, _fixedCode) {
			continue
		}
		oldEntity.EntityMap[key] = val
	}

	oldEntity.Imports = uniqueStrings(oldEntity.Imports)
	oldEntity.EntitySort = uniqueStrings(oldEntity.EntitySort)
	return oldEntity, nil
}

func (b *biz) scanEntity(src string) (*bizEntity, error) {
	nodeToString := func(fset *token.FileSet, node ast.Node) (string, error) {
		var buf bytes.Buffer
		if err := format.Node(&buf, fset, node); err != nil {
			return "", err
		}
		return buf.String(), nil
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	tps := &bizEntity{
		Package:   f.Name.Name,
		Imports:   []string{},
		EntityMap: make(map[string]string),
	}

	// Collect imports
	for _, imp := range f.Imports {
		if imp.Name != nil {
			tps.Imports = append(tps.Imports, imp.Name.Name+" "+imp.Path.Value)
		} else {
			tps.Imports = append(tps.Imports, imp.Path.Value)
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
			tps.EntitySort = append(tps.EntitySort, funcName)
			tps.EntityMap[funcName] = body

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
					tps.EntitySort = append(tps.EntitySort, typeName)
					tps.EntityMap[typeName] = body
				}
			}
		}
	}

	// Collect type definitions
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
				tps.EntitySort = append(tps.EntitySort, typeSpec.Name.Name)
				tps.EntityMap[typeSpec.Name.Name] = fullTypeDef
			}
		}
	}

	return tps, nil
}

func (b *biz) renderEntity(object *Object) (string, error) {
	tps, err := b.genEntity(object)
	if err != nil {
		return "", err
	}

	var (
		sb strings.Builder
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

	treeMd := []string{"ID", "Parent", "AppendChildren", "ChildrenNode"}
	// Write each type definition in the specified order
	for _, typeName := range tps.EntitySort {
		typeDef, ok := tps.EntityMap[typeName]
		if !ok {
			continue // Skip if there is no definition for the type name
		}

		if object.Type != _objectTypeTree && inList(treeMd, typeName) {
			continue
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

func (b *biz) genBizTplVariable(object *Object) map[string]any {
	var (
		filedMap = object.FieldMap()
		getCodes []string
	)
	for _, list := range object.Unique {
		var (
			keys    []string
			params  []string
			conds   []string
			codeTpl = "if %s {\n\tres, err = u.repo.%s(ctx, %s)\n\t}"
		)

		for _, item := range list {
			if toUpperCamelCase(item) == toUpperCamelCase("deleted_at") {
				continue
			}
			field, ok := filedMap[item]
			if !ok {
				continue
			}
			conds = append(conds, fmt.Sprintf("req.%s != nil", toUpperCamelCase(field.Keyword)))
			params = append(params, fmt.Sprintf("*req.%s", toUpperCamelCase(field.Keyword)))
			keys = append(keys, toUpperCamelCase(item))
		}
		if len(keys) != 0 {
			funcName := fmt.Sprintf("Get%sBy%s", toUpperCamelCase(object.Keyword), strings.Join(keys, "And"))
			code := fmt.Sprintf(codeTpl, strings.Join(conds, " && "), funcName, strings.Join(params, ","))
			getCodes = append(getCodes, code)
		}
	}

	getCodesStr := strings.Join(getCodes, "else ")
	if getCodesStr != "" {
		getCodesStr = getCodesStr + " else "
	}
	return map[string]any{
		"GetCodes":   getCodesStr,
		"Server":     object.Server,
		"ServerName": object.ServerName(),
		"Module":     toLowerCase(object.Module),
		"Object":     toUpperCamelCase(object.Keyword),
		"IsTree":     object.Type == _objectTypeTree,
		"Title":      object.Comment,
	}
}

func (b *biz) genBiz(object *Object) (*bizBiz, error) {
	oldBiz := &bizBiz{BizMap: make(map[string]string)}
	byteData, err := os.ReadFile(b.bizPath(object))
	if err == nil {
		res, err := b.scanBiz(string(byteData))
		if err != nil {
			return nil, err
		}
		oldBiz = res
	}

	tp, err := os.ReadFile(bizBizPath)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("go").Parse(strings.TrimSpace(string(tp)))
	if err != nil {
		return nil, err
	}

	if err := tmpl.Execute(buf, b.genBizTplVariable(object)); err != nil {
		return nil, err
	}

	newBiz, err := b.scanBiz(buf.String())
	if err != nil {
		return nil, err
	}

	oldBiz.Package = newBiz.Package
	oldBiz.Imports = append(oldBiz.Imports, newBiz.Imports...)
	oldBiz.BizSort = append(oldBiz.BizSort, newBiz.BizSort...)
	for key, val := range newBiz.BizMap {
		if oldVal := oldBiz.BizMap[key]; strings.Contains(oldVal, _fixedCode) {
			continue
		}
		oldBiz.BizMap[key] = val
	}

	oldBiz.Imports = uniqueStrings(oldBiz.Imports)
	oldBiz.BizSort = uniqueStrings(oldBiz.BizSort)
	return oldBiz, nil
}

func (b *biz) scanBiz(src string) (*bizBiz, error) {
	nodeToString := func(fset *token.FileSet, node ast.Node) (string, error) {
		var buf bytes.Buffer
		if err := format.Node(&buf, fset, node); err != nil {
			return "", err
		}
		return buf.String(), nil
	}

	biz := &bizBiz{BizMap: make(map[string]string)}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	biz.Package = f.Name.Name
	for _, imp := range f.Imports {
		if imp.Name != nil {
			biz.Imports = append(biz.Imports, imp.Name.Name+" "+imp.Path.Value)
		} else {
			biz.Imports = append(biz.Imports, imp.Path.Value)
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
			biz.BizSort = append(biz.BizSort, funcName)
			biz.BizMap[funcName] = body

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
					biz.BizSort = append(biz.BizSort, typeName)
					biz.BizMap[typeName] = body
				}
			}
		}
	}

	return biz, nil
}

func (b *biz) renderBiz(object *Object) (string, error) {
	tps, err := b.genBiz(object)
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
	for _, typeName := range tps.BizSort {
		typeDef, ok := tps.BizMap[typeName]
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

// isOptional 是否是可选项，这里暂志考虑基础数据类型
func (b *biz) isOptional(field *Field) bool {
	tp := b.mapping[field.Type].Struct
	if !field.Required || tp == _bool {
		return true
	}
	return false
}
