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
	repoModelPath = base.KratosxCliMod() + "/internal/autocode/template/repo/model.tpl"
	repoDataPath  = base.KratosxCliMod() + "/internal/autocode/template/repo/data.tpl"
)

type repo struct {
	mapping map[string]Mapping
}

type repoModel struct {
	Package   string
	Imports   []string
	ModelSort []string
	ModelMap  map[string]string
}

type repoData struct {
	Package  string
	Imports  []string
	DataSort []string
	DataMap  map[string]string
}

func GenRepo(object *Object) (map[string]string, error) {
	r := &repo{mapping: TypesMapping()}
	reply := map[string]string{}

	modelCode, err := r.renderModel(object)
	if err != nil {
		return nil, err
	}
	reply[r.modelPath(object)] = modelCode

	dataCode, err := r.renderData(object)
	if err != nil {
		return nil, err
	}
	reply[r.dataPath(object)] = dataCode

	return reply, nil
}

func (r *repo) dir() string {
	return strings.ToLower(fmt.Sprintf("internal/data"))
}

func (r *repo) modelPath(object *Object) string {
	return r.dir() + fmt.Sprintf("/model/%s.go", toLowerCase(object.Keyword))
}

func (r *repo) dataPath(object *Object) string {
	return r.dir() + fmt.Sprintf("/%s.go", toLowerCase(object.Module))
}

// isOptional 是否是可选项，这里暂志考虑基础数据类型
func (r *repo) isOptional(field *Field) bool {
	tp := r.mapping[field.Type].Struct
	if !field.Required || tp == _bool {
		return true
	}
	return false
}

func (r *repo) genModelTplVariable(object *Object) map[string]any {
	var (
		fields    []string
		relations []string
	)

	removeFields := []string{"Id", "CreatedAt", "UpdatedAt", "DeletedAt"}
	leave := 0
	for _, field := range object.Fields {
		stp := r.mapping[field.Type].Struct
		upKey := toUpperCamelCase(field.Keyword)
		opt := ""
		if r.isOptional(field) {
			opt = "*"
		}

		if inList(removeFields, upKey) {
			leave++
		} else {
			fields = append(fields, fmt.Sprintf("\t%s %s%s `json:\"%s\" gorm:\"%s\"`",
				upKey,
				opt,
				stp,
				toLowerCamelCase(field.Keyword),
				fmt.Sprintf("column:%s", toSnake(field.Keyword)),
			))
		}

		if field.Relation != nil {
			obj := field.Relation.Object
			if field.Relation.Type == _relationHasOne {
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

	switch leave {
	case 1, 2:
		fields = append(fields, "types.CreateModel")
	case 3:
		fields = append(fields, "types.BaseModel")
	case 4:
		fields = append(fields, "types.DeleteModel")
	}

	return map[string]any{
		"Fields": strings.Join(fields, "\n"),
		"Module": toLowerCase(object.Module),
		"Object": toUpperCamelCase(object.Keyword),
	}
}

func (r *repo) genModel(object *Object) (*repoModel, error) {
	oldModel := &repoModel{ModelMap: make(map[string]string)}
	byteData, err := os.ReadFile(r.modelPath(object))
	if err == nil {
		if res, err := r.scanModel(string(byteData)); err == nil {
			oldModel = res
		}
	}

	tp, err := os.ReadFile(repoModelPath)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("go").Parse(strings.TrimSpace(string(tp)))
	if err != nil {
		return nil, err
	}
	if err := tmpl.Execute(buf, r.genModelTplVariable(object)); err != nil {
		return nil, err
	}
	newModel, err := r.scanModel(buf.String())
	if err != nil {
		return nil, err
	}
	oldModel.Package = newModel.Package
	oldModel.Imports = append(oldModel.Imports, newModel.Imports...)
	oldModel.ModelSort = append(oldModel.ModelSort, newModel.ModelSort...)
	for key, val := range newModel.ModelMap {
		oldModel.ModelMap[key] = val
	}

	oldModel.Imports = uniqueStrings(oldModel.Imports)
	oldModel.ModelSort = uniqueStrings(oldModel.ModelSort)
	return oldModel, nil
}

func (r *repo) scanModel(src string) (*repoModel, error) {
	nodeToString := func(fset *token.FileSet, node ast.Node) (string, error) {
		var buf bytes.Buffer
		if err := format.Node(&buf, fset, node); err != nil {
			return "", err
		}
		return buf.String(), nil
	}

	doc := &repoModel{ModelMap: make(map[string]string)}

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
			doc.ModelSort = append(doc.ModelSort, funcName)
			doc.ModelMap[funcName] = body

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
					doc.ModelSort = append(doc.ModelSort, typeName)
					doc.ModelMap[typeName] = body
				}
			}
		}
	}

	return doc, nil
}

func (r *repo) renderModel(object *Object) (string, error) {
	tps, err := r.genModel(object)
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

	// Write each type definition in the specified order
	for _, typeName := range tps.ModelSort {
		typeDef, ok := tps.ModelMap[typeName]
		if !ok {
			continue // Skip if there is no definition for the type name
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

func (r *repo) getDataPreload(rela *FieldRelation, prefix string) []string {
	var preload []string
	var innerRelas []*FieldRelation
	for _, field := range rela.Object.Fields {
		relaKey := toUpperCamelCase(rela.Object.Keyword)
		if rela.Type == _relationHasMany {
			relaKey = pluralize(relaKey)
		}
		if field.Operation.Get {
			preload = append(preload, fmt.Sprintf(`Preload("%s")`, prefix+relaKey))
			if field.Relation != nil {
				innerRelas = append(innerRelas, field.Relation)
			}
		}
	}

	for _, item := range innerRelas {
		relaKey := toUpperCamelCase(rela.Object.Keyword)
		if rela.Type == _relationHasMany {
			relaKey = pluralize(relaKey)
		}
		prefix = prefix + relaKey + "."
		preload = append(preload, r.getDataPreload(item, prefix)...)
	}
	return uniqueStrings(preload)
}

func (r *repo) listDataPreload(rela *FieldRelation, prefix string) []string {
	var preload []string
	var innerRelas []*FieldRelation
	for _, field := range rela.Object.Fields {
		relaKey := toUpperCamelCase(rela.Object.Keyword)
		if rela.Type == _relationHasMany {
			relaKey = pluralize(relaKey)
		}
		if field.Operation.List {
			preload = append(preload, fmt.Sprintf(`Preload("%s")`, prefix+relaKey))
			if field.Relation != nil {
				innerRelas = append(innerRelas, field.Relation)
			}
		}
	}

	for _, item := range innerRelas {
		relaKey := toUpperCamelCase(rela.Object.Keyword)
		if rela.Type == _relationHasMany {
			relaKey = pluralize(relaKey)
		}
		prefix = prefix + relaKey + "."
		preload = append(preload, r.getDataPreload(item, prefix)...)
	}
	return uniqueStrings(preload)
}

func (r *repo) genDataTplVariable(object *Object) map[string]any {
	var (
		getFields       []string
		listFields      []string
		getTrashFields  []string
		listTrashFields []string
		getPreload      []string
		listPreload     []string
		queryCodes      []string
	)

	type byCodeType struct {
		Params string
		Where  string
		Method string
	}
	var (
		filedMap = object.FieldMap()
		byCodes  []byCodeType
	)
	for _, list := range object.Unique {
		var (
			keys   []string
			params []string
			where  []string
		)

		for _, item := range list {
			if toUpperCamelCase(item) == toUpperCamelCase("deleted_at") {
				continue
			}
			field, ok := filedMap[item]
			if !ok {
				continue
			}
			stp := r.mapping[field.Type].Struct
			lowKey := toLowerCamelCase(field.Keyword)
			params = append(params, fmt.Sprintf("%s %s", lowKey, stp))
			keys = append(keys, toUpperCamelCase(item))
			where = append(where, fmt.Sprintf(`Where("%s = ?",%s)`, lowKey, lowKey))
		}
		byCodes = append(byCodes, byCodeType{
			Params: strings.Join(params, ","),
			Method: strings.Join(keys, "And"),
			Where:  strings.Join(where, "."),
		})
	}

	for _, field := range object.Fields {
		upKey := toUpperCamelCase(field.Keyword)
		snakeKey := toSnake(field.Keyword)

		if field.Operation.Get {
			getTrashFields = append(getTrashFields, fmt.Sprintf(`"%s"`, snakeKey))
			if upKey != "DeletedAt" {
				getFields = append(getFields, fmt.Sprintf(`"%s"`, snakeKey))
			}
			if field.Relation != nil {
				getPreload = append(getPreload, r.getDataPreload(field.Relation, "")...)
			}
		}
		if (field.Operation.Get) && upKey != "DeletedAt" {
			listTrashFields = append(listTrashFields, fmt.Sprintf(`"%s"`, snakeKey))
			if upKey != "DeletedAt" {
				listFields = append(listFields, fmt.Sprintf(`"%s"`, snakeKey))
			}
			if field.Relation != nil {
				listPreload = append(listPreload, r.listDataPreload(field.Relation, "")...)
			}

			if field.QueryType != "" {
				switch strings.ToLower(field.QueryType) {
				case _in:
					tpl := `if req.%s != nil {
								db = db.Where("%s IN ?", *req.%s)
							}`
					code := fmt.Sprintf(tpl, pluralize(upKey), snakeKey, pluralize(upKey))
					queryCodes = append(queryCodes, code)
				case _notIn:
					tpl := `if req.%s != nil {
								db = db.Where("%s NOT IN ?", *req.%s)
							}`
					code := fmt.Sprintf(tpl, pluralize(upKey), snakeKey, pluralize(upKey))
					queryCodes = append(queryCodes, code)
				case _between:
					tpl := `if req.%s != nil {
								db = db.Where("%s BETWEEN ? AND ?", *req.%s[0], *req.%s[1])
							}`
					code := fmt.Sprintf(tpl, pluralize(upKey), snakeKey, pluralize(upKey))
					queryCodes = append(queryCodes, code)

				case _like:
					tpl := `if req.%s != nil {
								db = db.Where("%s LIKE ?", *req.%s+"%%")
							}`
					code := fmt.Sprintf(tpl, upKey, snakeKey, upKey)
					queryCodes = append(queryCodes, code)
				default:
					tpl := `if req.%s != nil {
								db = db.Where("%s %s ?", *req.%s)
							}`
					code := fmt.Sprintf(tpl, upKey, snakeKey, field.QueryType, upKey)
					queryCodes = append(queryCodes, code)
				}
			}
		}
	}

	return map[string]any{
		"ByCodes":         byCodes,
		"GetFields":       strings.Join(getFields, ", "),
		"ListFields":      strings.Join(listFields, ", "),
		"GetTrashFields":  strings.Join(getTrashFields, ", "),
		"ListTrashFields": strings.Join(getTrashFields, ", "),
		"Server":          object.Server,
		"Module":          toLowerCase(object.Module),
		"ModuleUpper":     toUpperCamelCase(object.Module),
		"ModuleLower":     toLowerCamelCase(object.Module),
		"Object":          toUpperCamelCase(object.Keyword),
		"HasGetPreload":   len(getPreload) != 0,
		"GetPreload":      strings.Join(getPreload, "."),
		"HasListPreload":  len(listPreload) != 0,
		"ListPreload":     strings.Join(listPreload, "."),
		"QueryCodes":      strings.Join(queryCodes, "\n"),
	}
}

func (r *repo) genData(object *Object) (*repoData, error) {
	oldData := &repoData{DataMap: make(map[string]string)}
	byteData, err := os.ReadFile(r.dataPath(object))
	if err == nil {
		if res, err := r.scanData(string(byteData)); err == nil {
			oldData = res
		}
	}

	tp, err := os.ReadFile(repoDataPath)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("go").Parse(strings.TrimSpace(string(tp)))
	if err != nil {
		return nil, err
	}
	if err := tmpl.Execute(buf, r.genDataTplVariable(object)); err != nil {
		return nil, err
	}
	newData, err := r.scanData(buf.String())
	if err != nil {
		return nil, err
	}
	oldData.Package = newData.Package
	oldData.Imports = append(oldData.Imports, newData.Imports...)
	oldData.DataSort = append(oldData.DataSort, newData.DataSort...)
	for key, val := range newData.DataMap {
		oldData.DataMap[key] = val
	}

	oldData.Imports = uniqueStrings(oldData.Imports)
	oldData.DataSort = uniqueStrings(oldData.DataSort)
	return oldData, nil
}

func (r *repo) scanData(src string) (*repoData, error) {
	nodeToString := func(fset *token.FileSet, node ast.Node) (string, error) {
		var buf bytes.Buffer
		if err := format.Node(&buf, fset, node); err != nil {
			return "", err
		}
		return buf.String(), nil
	}

	doc := &repoData{DataMap: make(map[string]string)}

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
			doc.DataSort = append(doc.DataSort, funcName)
			doc.DataMap[funcName] = body

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
					doc.DataSort = append(doc.DataSort, typeName)
					doc.DataMap[typeName] = body
				}
			}
		}
	}

	return doc, nil
}

func (r *repo) renderData(object *Object) (string, error) {
	tps, err := r.genData(object)
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
	for _, typeName := range tps.DataSort {
		typeDef, ok := tps.DataMap[typeName]
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
