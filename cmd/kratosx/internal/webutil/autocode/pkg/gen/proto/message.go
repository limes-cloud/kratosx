package proto

import (
	"bytes"
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

type Message struct {
	*gen.Builder
}

type MessageCode struct {
	pkg     string
	options []string
	imports []string
	sort    []string
	bucket  map[string]string
}

type MessageStruct struct {
	Alias        string
	RelationType string
	Rules        map[string]any
	Keyword      string
	Fields       []*MessageStructField
	IsOneOf      bool
	Relations    []*MessageStruct
}

type MessageStructField struct {
	Decorate string
	Keyword  string
	Type     string
	Validate string
}

func NewMessageBuilder(builder *gen.Builder) *Message {
	return &Message{Builder: builder}
}

func (p *Message) dir(object *types.Table) string {
	return strings.ToLower(fmt.Sprintf("api/%s/%s", p.Server, object.Module))
}

func (p *Message) version() string {
	return "v1"
}

func (p *Message) goPackage(object *types.Table) string {
	return "./v1;v1"
}

func (p *Message) packageName(object *types.Table) string {
	return strings.ReplaceAll(p.Server+"/"+p.dir(object)+"/"+p.version(), "/", ".")
}

func (p *Message) javaPackage(object *types.Table) string {
	return strings.ReplaceAll(p.Server+"/"+p.dir(object)+"/"+p.version(), "/", ".")
}

func (p *Message) javaClass(object *types.Table) string {
	s := strings.Split(p.dir(object)+"/"+p.version(), "/")
	return pkg.ToUpperHump(object.Module) + pkg.ToUpperHump(s[len(s)-1])
}

func (p *Message) objectName(object *types.Table) string {
	return pkg.ToUpperHump(object.Struct)
}

func (p *Message) objectComment(object *types.Table) string {
	return pkg.ToUpperHump(object.Comment)
}

// ParseMessageByContent 解析Message文本，生成Message结构
func (p *Message) ParseMessageByContent(content string) (*MessageCode, error) {
	reply := &MessageCode{bucket: make(map[string]string)}

	parser := proto.NewParser(strings.NewReader(content))
	definition, err := parser.Parse()
	if err != nil {
		return nil, err
	}

	// 正则扫描定义
	re := regexp.MustCompile(`message (\w+)([\s]*?)\{([\s\S]*?)\n\}`)
	matches := re.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		if len(match) == 4 { // 0 是整个匹配项，1 是消息名称，2 是消息体
			body := match[0]
			name := match[1]
			reply.sort = append(reply.sort, name)
			reply.bucket[name] = body
		}
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
		proto.WithMessage(func(message *proto.Message) {
			if message.Comment == nil || message.Comment.Message() == "" {
				return
			}
			body, ok := reply.bucket[message.Name]
			if !ok {
				return
			}
			reply.bucket[message.Name] = "// " + message.Comment.Message() + "\n" + body
		}),
	)

	return reply, nil
}

// ScanMessage 扫描已有Message
func (p *Message) ScanMessage() (*MessageCode, error) {
	path := p.ProtoMessagePath()
	if !pkg.IsExistFile(path) {
		return &MessageCode{}, nil
	}
	code, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return p.ParseMessageByContent(string(code))
}

// ruleString 规则转string
func (p *Message) ruleString(tp string, rules map[string]any) string {
	if len(rules) != 0 {
		var (
			tpl = "[(validate.rules).%s = {%s}]"
			arr []string
		)
		for key, val := range rules {
			switch val.(type) {
			case []any:
				var valArr []string
				list, _ := val.([]any)
				for _, v := range list {
					if _, ok := v.(string); ok {
						valArr = append(valArr, fmt.Sprintf(`"%v"`, v))
					} else {
						valArr = append(valArr, fmt.Sprintf(`%v`, v))
					}
				}
				rv := fmt.Sprintf("[%s]", strings.Join(valArr, ","))
				arr = append(arr, key+": "+rv)
			case string:
				arr = append(arr, fmt.Sprintf(`%s: "%s"`, key, val))
			default:
				arr = append(arr, key+": "+fmt.Sprint(val))
			}
		}
		return fmt.Sprintf(tpl, tp, strings.Join(arr, ","))
	}
	return ""
}

// renderStruct 渲染message struct
func (p *Message) renderStruct(msg *MessageStruct, spec string) string {
	set := map[string]bool{}
	text := fmt.Sprintf(spec+"message %s {\n", msg.Keyword)
	oldText := text
	for _, item := range msg.Relations {
		relation := *item

		// 判断引用类型
		pf := &MessageStructField{
			Keyword: pkg.VariableName(relation.Keyword, p.NameRule),
			Type:    pkg.ToUpperHump(relation.Keyword),
		}
		if !msg.IsOneOf {
			pf.Decorate = "optional "
		}
		pf.Validate = p.ruleString("repeated", relation.Rules)
		if p.IsRelationTypeMany(relation.RelationType) {
			pf.Keyword = pkg.ToPluralize(pf.Keyword)
			pf.Decorate = "repeated "
		}
		if relation.Alias != "" {
			pf.Keyword = relation.Alias
		}
		if !set[relation.Keyword] {
			relationText := p.renderStruct(&relation, spec+"  ")
			if relationText != "" {
				msg.Fields = append(msg.Fields, pf)
				text += relationText + "\n"
				set[relation.Keyword] = true
			}
		} else {
			msg.Fields = append(msg.Fields, pf)
		}
	}

	var rows []string
	for index, field := range msg.Fields {
		// uint32 field = number[(validate.rules).uint32 = {gt: 0}];
		row := fmt.Sprintf(spec+"  %s%s %s = %d%s;", field.Decorate, field.Type, field.Keyword, index+1, field.Validate)
		rows = append(rows, row)
	}

	if msg.IsOneOf && len(rows) > 1 {
		for ind, row := range rows {
			rows[ind] = "  " + spec + row
		}
		oneof := "  " + spec + "oneof params{\n"
		oneof += strings.Join(rows, "\n") + "\n"
		oneof += "  " + spec + "}\n"
		text += oneof
	} else {
		text += strings.Join(rows, "\n") + "\n"
	}

	if text == oldText && spec != "" {
		return ""
	}

	text += spec + "}"

	return text
}

// makeUpdateRequest 生成 create request
func (p *Message) makeCreateRequest(table *types.Table, deep ...bool) *MessageStruct {
	keyword := pkg.ToUpperHump(table.Struct)
	if len(deep) == 0 || !deep[0] {
		keyword = fmt.Sprintf("Create%sRequest", pkg.ToUpperHump(table.Struct))
	}

	msg := &MessageStruct{Keyword: keyword}

	for _, column := range table.Columns {
		if column.Operation.Create {
			tp := column.ProtoType()
			pf := &MessageStructField{
				Keyword:  pkg.VariableName(column.Name, p.NameRule),
				Type:     tp,
				Validate: p.ruleString(tp, column.Rules),
			}
			if column.IsProtoOption() {
				pf.Decorate = "optional "
			}
			msg.Fields = append(msg.Fields, pf)
		}

		for _, relation := range column.Relations {
			temp := *relation
			pm := p.makeCreateRequest(temp.Table, true)
			pm.RelationType = relation.Type
			pm.Rules = relation.Rules
			msg.Relations = append(msg.Relations, pm)
		}
	}
	return msg
}

// makeUpdateRequest 生成 update request
func (p *Message) makeUpdateRequest(table *types.Table, deep ...bool) *MessageStruct {
	keyword := pkg.ToUpperHump(table.Struct)
	if len(deep) == 0 || !deep[0] {
		keyword = fmt.Sprintf("Update%sRequest", pkg.ToUpperHump(table.Struct))
	}
	msg := &MessageStruct{Keyword: keyword}

	for _, column := range table.Columns {
		if column.Operation.Update {
			tp := column.ProtoType()
			pf := &MessageStructField{
				Keyword:  pkg.VariableName(column.Name, p.NameRule),
				Type:     tp,
				Validate: p.ruleString(tp, column.Rules),
			}
			if column.Name != "id" {
				pf.Decorate = "optional "
			}
			msg.Fields = append(msg.Fields, pf)
		}

		for _, relation := range column.Relations {
			temp := *relation
			pm := p.makeUpdateRequest(temp.Table)
			pm.RelationType = relation.Type
			pm.Rules = relation.Rules
			msg.Relations = append(msg.Relations, pm)
		}
	}
	return msg
}

// makeGetRequest 生成 get request
func (p *Message) makeGetRequest(table *types.Table, trash bool, deep ...bool) *MessageStruct {
	keyword := pkg.ToUpperHump(table.Struct)
	if len(deep) == 0 || !deep[0] {
		if trash {
			keyword = fmt.Sprintf("GetTrash%sRequest", pkg.ToUpperHump(table.Struct))
		} else {
			keyword = fmt.Sprintf("Get%sRequest", pkg.ToUpperHump(table.Struct))
		}
	}

	msg := &MessageStruct{
		IsOneOf: true,
		Keyword: keyword,
		Fields: []*MessageStructField{
			{
				Keyword:  "id",
				Type:     "uint32",
				Validate: "[(validate.rules).uint32 = {gte: 1}]",
			},
		},
	}

	columns := map[string]*types.Column{}

	for _, column := range table.Columns {
		if column.IsDeletedAt() {
			continue
		}
		columns[column.Name] = column
	}

	for _, index := range table.Indexes {
		if !index.Unique {
			continue
		}

		var (
			unique []string
			fields []*MessageStructField
		)

		for _, name := range index.Names {
			column, ok := columns[name]
			if !ok {
				continue
			}

			unique = append(unique, pkg.ToUpperHump(name))
			tp := column.ProtoType()
			fields = append(fields, &MessageStructField{
				Keyword:  pkg.VariableName(column.Name, p.NameRule),
				Type:     tp,
				Validate: p.ruleString(tp, column.Rules),
			})
		}
		if len(unique) == 0 {
			continue
		}

		if len(unique) > 1 {
			keyword := strings.Join(unique, "And")
			relation := &MessageStruct{
				Keyword:      keyword,
				RelationType: types.RelationTypeOne,
				Fields:       fields,
			}
			msg.Relations = append(msg.Relations, relation)
		} else {
			msg.Fields = append(msg.Fields, fields...)
		}
	}

	return msg
}

// makeGetReply 生成 get reply
func (p *Message) makeGetReply(table *types.Table, trash bool, deep ...bool) *MessageStruct {
	keyword := pkg.ToUpperHump(table.Struct)
	if len(deep) == 0 || !deep[0] {
		if trash {
			keyword = fmt.Sprintf("GetTrash%sReply", pkg.ToUpperHump(table.Struct))
		} else {
			keyword = fmt.Sprintf("Get%sReply", pkg.ToUpperHump(table.Struct))
		}
	}
	msg := &MessageStruct{Keyword: keyword}
	for _, column := range table.Columns {
		if !column.Operation.Get {
			continue
		}

		if column.IsDeletedAt() && !trash {
			continue
		}

		tp := column.ProtoType()
		field := &MessageStructField{
			Keyword: pkg.VariableName(column.Name, p.NameRule),
			Type:    tp,
			// Validate: p.ruleString(tp, column.Rules),
		}
		if column.IsProtoOption() {
			field.Decorate = "optional "
		}
		msg.Fields = append(msg.Fields, field)

		for _, item := range column.Relations {
			relation := *item
			pm := p.makeGetReply(relation.Table, trash, true)
			pm.RelationType = relation.Type
			msg.Relations = append(msg.Relations, pm)
		}
	}
	return msg
}

func (p *Message) makeListRequest(table *types.Table, trash bool, deep ...bool) *MessageStruct {
	keyword := pkg.ToUpperHump(table.Struct)
	if len(deep) == 0 || !deep[0] {
		if trash {
			keyword = fmt.Sprintf("ListTrash%sRequest", pkg.ToUpperHump(table.Struct))
		} else {
			keyword = fmt.Sprintf("List%sRequest", pkg.ToUpperHump(table.Struct))
		}
	}
	msg := &MessageStruct{Keyword: keyword}

	if table.Type == types.TableTypeList {
		msg.Fields = append(msg.Fields, []*MessageStructField{
			{
				Keyword:  "page",
				Type:     "uint32",
				Validate: "[(validate.rules).uint32 = {gte: 1}]",
			},
			{
				Keyword:  "pageSize",
				Type:     "uint32",
				Validate: "[(validate.rules).uint32 = {gte: 1,lte:50}]",
			},
		}...)
	}

	// 获取索引，用于排序
	var indexes = []string{`"id"`}
	for _, index := range table.Indexes {
		if len(index.Names) == 0 {
			continue
		}
		indexName := index.Names[0]
		if indexName == "deleted_at" {
			continue
		}
		indexes = append(indexes, fmt.Sprintf(`"%s"`, pkg.ToSnake(indexName)))
	}
	indexes = pkg.UniqueArray(indexes)

	// 添加排序字段
	msg.Fields = append(msg.Fields, []*MessageStructField{
		{
			Decorate: "optional ",
			Keyword:  "order",
			Type:     "string",
			Validate: "[(validate.rules).string = {in: [\"asc\",\"desc\"]}]",
		},
		{
			Decorate: "optional ",
			Keyword:  "orderBy",
			Type:     "string",
			Validate: fmt.Sprintf("[(validate.rules).string = {in: [%s]}]", strings.Join(indexes, ",")),
		},
	}...)

	for _, column := range table.Columns {
		if column.Query.Type == "" {
			continue
		}

		tp := column.ProtoType()
		field := &MessageStructField{
			Decorate: "optional ",
			Keyword:  pkg.VariableName(column.Name, p.NameRule),
			Type:     tp,
			Validate: p.ruleString(tp, column.Rules),
		}

		if column.Query.IsPluralize() {
			field.Decorate = "repeated "
			field.Keyword = pkg.ToPluralize(field.Keyword)
		}

		msg.Fields = append(msg.Fields, field)
	}
	return msg
}

func (p *Message) makeListReply(table *types.Table, trash bool, deep ...bool) *MessageStruct {
	keyword := pkg.ToUpperHump(table.Struct)
	if len(deep) == 0 || !deep[0] {
		if trash {
			keyword = fmt.Sprintf("ListTrash%sReply", pkg.ToUpperHump(table.Struct))
		} else {
			keyword = fmt.Sprintf("List%sReply", pkg.ToUpperHump(table.Struct))
		}
	}

	msg := &MessageStruct{
		Keyword: pkg.ToUpperHump(table.Struct),
	}

	for _, column := range table.Columns {
		if !column.Operation.List {
			continue
		}
		if column.IsDeletedAt() && !trash {
			continue
		}

		tp := column.ProtoType()
		field := &MessageStructField{
			Keyword: pkg.VariableName(column.Name, p.NameRule),
			Type:    tp,
		}
		if column.IsProtoOption() {
			field.Decorate = "optional "
		}

		msg.Fields = append(msg.Fields, field)
		for _, item := range column.Relations {
			relation := *item
			pm := p.makeListReply(relation.Table, trash, true)
			pm.RelationType = relation.Type
			msg.Relations = append(msg.Relations, pm)
		}
	}

	if table.Type == types.TableTypeTree {
		msg.Fields = append(msg.Fields, &MessageStructField{
			Decorate: "repeated ",
			Keyword:  "children",
			Type:     pkg.ToUpperHump(table.Struct),
		})
	}

	if len(deep) == 0 || !deep[0] {
		msg.Alias = "list"
		msg.RelationType = types.RelationTypeMany
		parent := &MessageStruct{
			Keyword:   keyword,
			Relations: []*MessageStruct{msg},
		}
		if table.Type != types.TableTypeTree {
			parent.Fields = append(parent.Fields, &MessageStructField{
				Keyword: "total",
				Type:    "uint32",
			})
		}
		return parent
	}

	return msg
}

func (p *Message) MakeMessage() (*MessageCode, error) {
	tp, err := os.ReadFile(p.ProtoMessageTplPath())
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("Message").Parse(strings.TrimSpace(string(tp)))
	if err != nil {
		return nil, err
	}

	renderData := map[string]any{
		"EnableBatchDelete": p.Table.EnableBatchDelete,
		"Package":           p.packageName(p.Table),
		"GoPackage":         p.goPackage(p.Table),
		"JavaPackage":       p.javaPackage(p.Table),
		"JavaClass":         p.javaClass(p.Table),
		"Object":            p.objectName(p.Table),
		"CreateRequest":     p.renderStruct(p.makeCreateRequest(p.Table), ""),
		"UpdateRequest":     p.renderStruct(p.makeUpdateRequest(p.Table), ""),
		"GetRequest":        p.renderStruct(p.makeGetRequest(p.Table, false), ""),
		"GetReply":          p.renderStruct(p.makeGetReply(p.Table, false), ""),
		"GetTrashRequest":   p.renderStruct(p.makeGetRequest(p.Table, true), ""),
		"GetTrashReply":     p.renderStruct(p.makeGetReply(p.Table, true), ""),
		"ListRequest":       p.renderStruct(p.makeListRequest(p.Table, false), ""),
		"ListReply":         p.renderStruct(p.makeListReply(p.Table, false), ""),
		"ListTrashRequest":  p.renderStruct(p.makeListRequest(p.Table, true), ""),
		"ListTrashReply":    p.renderStruct(p.makeListReply(p.Table, true), ""),
	}
	if err := tmpl.Execute(buf, renderData); err != nil {
		return nil, err
	}

	return p.ParseMessageByContent(buf.String())
}

func (p *Message) RenderMessage(msg *MessageCode) string {
	content := "syntax = \"proto3\";\n\n"
	content += "package " + msg.pkg + ";\n\n"
	for _, val := range msg.options {
		content += val + ";\n"
	}

	content += "\n"
	for _, val := range msg.imports {
		content += val + ";\n"
	}

	trash := p.HasDeletedAt()
	content += "\n"
	for _, val := range msg.sort {
		if !trash && strings.Contains(val, "Trash") {
			continue
		}
		content += msg.bucket[val] + "\n\n"

		// mth := strings.TrimSuffix(val, "Request")
		// mth = strings.TrimSuffix(mth, "Reply")
		// if p.HasMethod(mth) {
		//
		// }
	}
	return content
}

func (p *Message) GenMessage() (string, error) {
	// 扫描已经生成的proto
	scanMessage, err := p.ScanMessage()
	if err != nil {
		return "", fmt.Errorf("扫描proto message代码失败，%s", err.Error())
	}

	// 生成新的proto
	makeMessage, err := p.MakeMessage()
	if err != nil {
		return "", fmt.Errorf("生成proto message代码失败，%s", err.Error())
	}

	// 合并proto
	code := &MessageCode{bucket: make(map[string]string), pkg: makeMessage.pkg}
	code.sort = append(makeMessage.sort, scanMessage.sort...)
	code.imports = append(makeMessage.imports, scanMessage.imports...)
	code.options = append(makeMessage.options, scanMessage.options...)

	for key, val := range scanMessage.bucket {
		code.bucket[key] = val
	}
	for key, val := range makeMessage.bucket {
		code.bucket[key] = val
	}

	// 去重
	code.sort = pkg.UniqueArray(code.sort)
	code.imports = pkg.UniqueArray(code.imports)
	code.options = pkg.UniqueArray(code.options)

	// 生成proto
	return p.RenderMessage(code), nil
}
