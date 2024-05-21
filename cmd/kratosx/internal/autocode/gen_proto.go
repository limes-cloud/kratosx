package autocode

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"
)

const (
	protoMsgPath = "internal/autocode/template/proto/message.tpl"
	protoSrvPath = "internal/autocode/template/proto/service.tpl"
	protoErrPath = "internal/autocode/template/proto/error.tpl"
)

type proto struct {
	mapping map[string]Mapping
}

type protoService struct {
	Package      string
	Server       string
	Options      []string
	Imports      []string
	FunctionSort []string
	FunctionMap  map[string]string
}

type protoMessage struct {
	Package     string
	Options     []string
	Imports     []string
	MessageSort []string
	MessageMap  map[string]string
}

type protoMessageBody struct {
	RelationType string
	Keyword      string
	Fields       []*protoMessageField
	Relations    []*protoMessageBody
}

type protoMessageField struct {
	Decorate string
	Keyword  string
	Type     string
	Validate string
}

type protoError struct {
	ErrorSort []string
	ErrorMap  map[string]string
}

func GenProto(object *Object) (map[string]string, error) {
	p := &proto{mapping: TypesMapping()}

	reply := map[string]string{}
	tempProtoMsg, err := p.renderTemplateProtoMsg(object)
	if err != nil {
		return nil, err
	}
	reply[p.msgPath(object)] = tempProtoMsg

	tempProtoSrv, err := p.renderTemplateProtoSrv(object)
	if err != nil {
		return nil, err
	}
	reply[p.srvPath(object)] = tempProtoSrv

	tempProtoErr, err := p.renderError(object)
	if err != nil {
		return nil, err
	}
	reply[p.errorPath(object)] = tempProtoErr

	return reply, nil
}

func (p *proto) renderMsg(msg *protoMessageBody, prefixSpace string) string {
	// message Create{{.Object}}Request {
	//	{{.CreateRequest}}
	// }
	text := fmt.Sprintf(prefixSpace+"message %s {\n", msg.Keyword)
	for _, relation := range msg.Relations {
		// 判断引用类型
		pf := &protoMessageField{
			Decorate: "optional ",
			Keyword:  toLowerCamelCase(relation.Keyword),
			Type:     toUpperCamelCase(relation.Keyword),
		}
		if msg.RelationType == _relationHasMany {
			pf.Keyword = pluralize(pf.Keyword)
			pf.Decorate = "repeated "
		}
		msg.Fields = append(msg.Fields, pf)

		relationText := p.renderMsg(relation, prefixSpace+"  ")
		text += relationText + "\n"
	}

	for index, field := range msg.Fields {
		// uint32 field = number[(validate.rules).uint32 = {gt: 0}];
		row := fmt.Sprintf(prefixSpace+"  %s%s %s = %d%s;\n", field.Decorate, field.Type, toLowerCamelCase(field.Keyword), index+1, field.Validate)
		text += row
	}

	text += prefixSpace + "}"

	return text
}

func (p *proto) genCreateRequestMsg(object *Object) *protoMessageBody {
	msg := &protoMessageBody{
		Keyword: object.StructName(),
	}
	for _, field := range object.Fields {
		if !field.Operation.Create {
			continue
		}

		tp := p.mapping[field.Type].Proto
		pf := &protoMessageField{
			Keyword:  field.Keyword,
			Type:     tp,
			Validate: p.ruleToString(tp, field.Rules),
		}
		if p.isOptional(field) {
			pf.Decorate = "optional "
		}
		msg.Fields = append(msg.Fields, pf)
		if field.Relation != nil {
			pm := p.genCreateRequestMsg(field.Relation.Object)
			msg.Relations = append(msg.Relations, pm)
			msg.RelationType = field.Relation.Type
		}
	}
	return msg
}

func (p *proto) genUpdateRequestMsg(object *Object) *protoMessageBody {
	msg := &protoMessageBody{
		Keyword: object.StructName(),
	}
	for _, field := range object.Fields {
		if !field.Operation.Update {
			continue
		}

		tp := p.mapping[field.Type].Proto
		pf := &protoMessageField{
			Keyword:  field.Keyword,
			Type:     tp,
			Validate: p.ruleToString(tp, field.Rules),
		}
		if p.isOptional(field) {
			pf.Decorate = "optional "
		}
		msg.Fields = append(msg.Fields, pf)
		if field.Relation != nil {
			pm := p.genUpdateRequestMsg(field.Relation.Object)
			msg.Relations = append(msg.Relations, pm)
			msg.RelationType = field.Relation.Type
		}
	}
	return msg
}

func (p *proto) genGetReplyMsg(object *Object) *protoMessageBody {
	msg := &protoMessageBody{
		Keyword: object.StructName(),
	}
	for _, field := range object.Fields {
		if !field.Operation.Get {
			continue
		}

		tp := p.mapping[field.Type].Proto
		pf := &protoMessageField{
			Keyword: field.Keyword,
			Type:    tp,
		}
		if p.isOptional(field) {
			pf.Decorate = "optional "
		}
		msg.Fields = append(msg.Fields, pf)
		if field.Relation != nil {
			pm := p.genGetReplyMsg(field.Relation.Object)
			msg.Relations = append(msg.Relations, pm)
			msg.RelationType = field.Relation.Type
		}
	}
	return msg
}

func (p *proto) genExportRequestMsg(object *Object) *protoMessageBody {
	msg := &protoMessageBody{
		Keyword: fmt.Sprintf("Export%sRequest", object.StructName()),
	}

	for _, field := range object.Fields {
		if field.QueryType == "" {
			continue
		}
		tp := p.mapping[field.Type].Proto
		pf := &protoMessageField{
			Keyword: field.Keyword,
			Type:    tp,
		}
		switch strings.ToLower(field.QueryType) {
		case _in, _notIn, _between:
			pf.Decorate = "repeated "
		default:
			pf.Decorate = "optional "
		}
		msg.Fields = append(msg.Fields, pf)
	}
	return msg
}

func (p *proto) genListRequestMsg(object *Object) *protoMessageBody {
	msg := &protoMessageBody{
		Keyword: fmt.Sprintf("List%sRequest", object.StructName()),
	}
	if object.Type == _objectTypeList {
		msg.Fields = append(msg.Fields, []*protoMessageField{
			{
				Keyword:  "page",
				Type:     "uint32",
				Validate: "[(validate.rules).uint32 = {gt: 0}]",
			},
			{
				Keyword:  "pageSize",
				Type:     "uint32",
				Validate: "[(validate.rules).uint32 = {gt: 0,lte:50}]",
			},
		}...)
	}

	msg.Fields = append(msg.Fields, []*protoMessageField{
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
		},
	}...)

	for _, field := range object.Fields {
		if field.QueryType == "" {
			continue
		}
		tp := p.mapping[field.Type].Proto
		pf := &protoMessageField{
			Keyword: field.Keyword,
			Type:    tp,
		}
		switch strings.ToLower(field.QueryType) {
		case _in, _notIn, _between:
			pf.Decorate = "repeated "
		default:
			pf.Decorate = "optional "
		}
		msg.Fields = append(msg.Fields, pf)
	}
	return msg
}

func (p *proto) genListReplyMsg(object *Object) *protoMessageBody {
	msg := &protoMessageBody{
		Keyword: object.StructName(),
	}

	for _, field := range object.Fields {
		if !field.Operation.List {
			continue
		}

		tp := p.mapping[field.Type].Proto
		pf := &protoMessageField{
			Keyword: field.Keyword,
			Type:    tp,
		}
		if p.isOptional(field) {
			pf.Decorate = "optional "
		}
		msg.Fields = append(msg.Fields, pf)
		if field.Relation != nil {
			pm := p.genListReplyMsg(field.Relation.Object)
			msg.Relations = append(msg.Relations, pm)
			msg.RelationType = field.Relation.Type
		}
	}

	if object.Type == _objectTypeTree {
		msg.Fields = append(msg.Fields, &protoMessageField{
			Decorate: "repeated ",
			Keyword:  "children",
			Type:     object.StructName(),
		})
	}
	return msg
}

// isOptional 是否是可选项，这里暂志考虑基础数据类型
func (p *proto) isOptional(field *Field) bool {
	tp := p.mapping[field.Type].Proto
	if !field.Required || tp == _bool {
		return true
	}
	return false
}

func (p *proto) initObjectRules(object *Object) {
	for _, field := range object.Fields {
		if field.Rules == nil {
			field.Rules = make(map[string]any)
		}
		tp := p.mapping[field.Type].Proto
		if field.Required {
			switch tp {
			case _string:
				if field.Rules["min_len"] == nil {
					field.Rules["min_len"] = 1
				}
			}
		}
		switch field.Type {
		case _primaryKey, _foreignKey:
			field.Rules["gt"] = 0
		}
	}
}

func (p *proto) ruleToString(tp string, rules map[string]any) string {
	// [(validate.rules).uint32 = {gt: 0,lte:50}]
	if len(rules) != 0 {
		var (
			tpl = "[(validate.rules).%s = {%s}]"
			arr []string
		)
		for key, val := range rules {
			arr = append(arr, key+": "+fmt.Sprint(val))
		}
		return fmt.Sprintf(tpl, tp, strings.Join(arr, ","))
	}
	return ""
}

func (p *proto) dir(object *Object) string {
	return strings.ToLower(fmt.Sprintf("api/%s/%s", object.ServerName(), object.Module))
}

func (p *proto) msgPath(object *Object) string {
	name := strings.ToLower(fmt.Sprintf("%s_%s.proto", object.ServerName(), object.Keyword))
	return p.dir(object) + "/" + name
}

func (p *proto) errorPath(object *Object) string {
	path := strings.ToLower(fmt.Sprintf("api/%s/errors", object.ServerName()))
	return strings.ToLower(fmt.Sprintf("%s/%s_error_reason.proto", path, object.ServerName()))
}

func (p *proto) srvPath(object *Object) string {
	name := strings.ToLower(fmt.Sprintf("%s_%s_service.proto", object.ServerName(), object.Module))
	return p.dir(object) + "/" + name
}

func (p *proto) version() string {
	return "v1"
}

func (p *proto) goPackage(object *Object) string {
	return "./v1;v1"
	// return object.Server + "/" + p.dir(object) + ";" + s[len(s)-1]
}

func (p *proto) packageName(object *Object) string {
	return strings.ReplaceAll(object.ServerName()+"/"+p.dir(object)+"/"+p.version(), "/", ".")
}

func (p *proto) javaPackage(object *Object) string {
	return strings.ReplaceAll(object.ServerName()+"/"+p.dir(object)+"/"+p.version(), "/", ".")
}

func (p *proto) javaClass(object *Object) string {
	s := strings.Split(p.dir(object)+"/"+p.version(), "/")
	return toUpperCamelCase(object.Keyword) + toUpperCamelCase(s[len(s)-1])
}

func (p *proto) objectName(object *Object) string {
	return toUpperCamelCase(object.Keyword)
}

func (p *proto) objectComment(object *Object) string {
	return toUpperCamelCase(object.Comment)
}

func (p *proto) scanMessage(content string) *protoMessage {
	reply := &protoMessage{MessageMap: make(map[string]string)}
	lines := strings.Split(content, "\n")
	for i := 0; i < len(lines); i++ {
		text := strings.TrimSpace(lines[i])
		if strings.HasPrefix(text, "import") {
			reply.Imports = append(reply.Imports, text)
		}
		if strings.HasPrefix(text, "option") {
			reply.Options = append(reply.Options, text)
		}
		if strings.HasPrefix(text, "message ") {
			break
		}
	}

	re := regexp.MustCompile(`message (\w+) \{([\s\S]*?)\n\}`)
	matches := re.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		if len(match) == 3 { // 0 是整个匹配项，1 是消息名称，2 是消息体
			messageBody := match[0]
			messageName := match[1]
			reply.MessageSort = append(reply.MessageSort, messageName)
			reply.MessageMap[messageName] = messageBody
		}
	}
	return reply
}

func (p *proto) genMessage(object *Object) (*protoMessage, error) {
	msg := &protoMessage{MessageMap: make(map[string]string)}
	path := p.msgPath(object)
	byteData, err := os.ReadFile(path)
	if err == nil {
		msg = p.scanMessage(string(byteData))
	}

	p.initObjectRules(object)
	createRequestMessage := p.genCreateRequestMsg(object)
	createRequestMessage.Keyword = fmt.Sprintf("Create%sRequest", createRequestMessage.Keyword)

	updateRequestMessage := p.genUpdateRequestMsg(object)
	updateRequestMessage.Keyword = fmt.Sprintf("Update%sRequest", updateRequestMessage.Keyword)

	getReplyMessage := p.genGetReplyMsg(object)
	getReplyMessage.Keyword = fmt.Sprintf("Get%sReply", getReplyMessage.Keyword)

	listRequestMessage := p.genListRequestMsg(object)
	listReplyMessage := p.genListReplyMsg(object)
	exportRequestMessage := p.genExportRequestMsg(object)

	// 渲染模板
	tp, err := os.ReadFile(protoMsgPath)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("proto").Parse(strings.TrimSpace(string(tp)))
	if err != nil {
		return nil, err
	}

	renderData := map[string]any{
		"Package":       p.packageName(object),
		"GoPackage":     p.goPackage(object),
		"JavaPackage":   p.javaPackage(object),
		"JavaClass":     p.javaClass(object),
		"Object":        p.objectName(object),
		"CreateRequest": p.renderMsg(createRequestMessage, ""),
		"ExportRequest": p.renderMsg(exportRequestMessage, ""),
		"UpdateRequest": p.renderMsg(updateRequestMessage, ""),
		"GetReply":      p.renderMsg(getReplyMessage, ""),
		"ListReply":     p.renderMsg(listReplyMessage, "  "),
		"ListRequest":   p.renderMsg(listRequestMessage, ""),
	}
	if err := tmpl.Execute(buf, renderData); err != nil {
		return nil, err
	}

	newMsg := p.scanMessage(buf.String())

	msg.Package = fmt.Sprintf("package %s", p.packageName(object))
	msg.Imports = append(msg.Imports, newMsg.Imports...)
	msg.Options = append(msg.Options, newMsg.Options...)
	msg.MessageSort = append(msg.MessageSort, newMsg.MessageSort...)
	for key, val := range newMsg.MessageMap {
		msg.MessageMap[key] = val
	}

	msg.Imports = uniqueStrings(msg.Imports)
	msg.Options = uniqueStrings(msg.Options)
	msg.MessageSort = uniqueStrings(msg.MessageSort)
	return msg, nil
}

func (p *proto) renderTemplateProtoMsg(object *Object) (string, error) {
	msg, err := p.genMessage(object)
	if err != nil {
		return "", err
	}
	md := object.MethodStatus()

	content := "syntax = \"proto3\";\n\n"
	content += msg.Package + ";\n\n"
	for _, val := range msg.Options {
		content += val + "\n"
	}

	content += "\n"
	for _, val := range msg.Imports {
		content += val + "\n"
	}

	content += "\n"
	for _, val := range msg.MessageSort {
		mth := strings.TrimSuffix(val, "Request")
		mth = strings.TrimSuffix(mth, "Reply")
		if object.HasMethod(md, mth) {
			content += msg.MessageMap[val] + "\n\n"
		}
	}
	return content, nil
}

func (p *proto) scanService(content string) *protoService {
	reply := &protoService{FunctionMap: make(map[string]string)}

	lines := strings.Split(content, "\n")

	var i = 0
	for ; i < len(lines); i++ {
		text := strings.TrimSpace(lines[i])
		if strings.HasPrefix(text, "import") {
			reply.Imports = append(reply.Imports, text)
		}
		if strings.HasPrefix(text, "option") {
			reply.Options = append(reply.Options, text)
		}
		if strings.HasPrefix(text, "service ") {
			i++
			break
		}
	}

	var (
		comment   = ""
		rpcBody   = ""
		preMethod = ""
		rpcMap    = map[string]string{}
		rpcSort   []string
	)
	for ; i < len(lines); i++ {
		text := strings.TrimSpace(lines[i])
		if text == "" {
			continue
		}

		if strings.HasPrefix(text, "//") {
			comment += lines[i] + "\n"
		}

		if strings.HasPrefix(text, "rpc") {
			if preMethod != "" {
				rpcMap[preMethod] = rpcBody
				rpcBody = ""
				rpcSort = append(rpcSort, preMethod)
				preMethod = strings.Fields(text)[1]
			}
			rpcBody = comment + rpcBody
			preMethod = strings.Fields(text)[1]
			comment = ""
		}
		if comment == "" {
			rpcBody += lines[i] + "\n"
		}
	}
	if rpcBody != "" {
		rpcBody = comment + rpcBody
		rpcBody = rpcBody[:strings.LastIndex(rpcBody, "\n")-1]
		rpcMap[preMethod] = rpcBody
		rpcSort = append(rpcSort, preMethod)
	}

	reply.FunctionMap = rpcMap
	reply.FunctionSort = rpcSort
	return reply
}

func (p *proto) genService(object *Object) (*protoService, error) {
	srv := &protoService{FunctionMap: make(map[string]string)}
	path := p.srvPath(object)
	byteData, err := os.ReadFile(path)
	if err == nil {
		srv = p.scanService(string(byteData))
	}

	srv.Imports = append(srv.Imports,
		fmt.Sprintf("import \"%s\";", strings.ToLower(fmt.Sprintf("%s_%s.proto", object.ServerName(), object.Keyword))),
	)

	tp, err := os.ReadFile(protoSrvPath)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	tmpl, err := template.New("proto").Parse(strings.TrimSpace(string(tp)))
	if err != nil {
		return nil, err
	}

	renderData := map[string]any{
		"Package":         p.packageName(object),
		"GoPackage":       p.goPackage(object),
		"JavaPackage":     p.javaPackage(object),
		"JavaClass":       p.javaClass(object),
		"Object":          p.objectName(object),
		"ObjectLowerCase": toLowerCamelCase(p.objectName(object)),
		"ServerLowerCase": toLowerCamelCase(object.ServerName()),
		"ModuleLowerCase": toLowerCamelCase(object.Module),
		"Title":           p.objectComment(object),
	}

	if err := tmpl.Execute(buf, renderData); err != nil {
		return nil, err
	}

	newPrv := p.scanService(buf.String())
	srv.Package = fmt.Sprintf("package %s", p.packageName(object))
	srv.Server = toUpperCamelCase(object.Module)
	srv.Imports = append(srv.Imports, newPrv.Imports...)
	srv.Options = append(srv.Options, newPrv.Options...)
	srv.FunctionSort = append(srv.FunctionSort, newPrv.FunctionSort...)
	for key, val := range newPrv.FunctionMap {
		srv.FunctionMap[key] = val
	}

	srv.Imports = uniqueStrings(srv.Imports)
	srv.Options = uniqueStrings(srv.Options)
	srv.FunctionSort = uniqueStrings(srv.FunctionSort)
	return srv, nil
}

func (p *proto) renderTemplateProtoSrv(object *Object) (string, error) {
	prv, err := p.genService(object)
	if err != nil {
		return "", err
	}

	md := object.MethodStatus()

	content := "syntax = \"proto3\";\n\n"
	content += prv.Package + ";\n\n"
	for _, val := range prv.Options {
		content += val + "\n"
	}

	content += "\n"
	for _, val := range prv.Imports {
		content += val + "\n"
	}

	content += "\n"
	content += fmt.Sprintf("service %s{\n\n", prv.Server)
	for _, val := range prv.FunctionSort {
		if object.HasMethod(md, val) {
			content += prv.FunctionMap[val] + "\n"
		}
	}
	content += "}"
	return content, nil
}

func (p *proto) genError(object *Object) (*protoError, error) {
	oldError := &protoError{ErrorMap: make(map[string]string)}
	byteData, err := os.ReadFile(p.errorPath(object))
	if err == nil {
		if res, err := p.scanError(string(byteData)); err == nil {
			oldError = res
		}
	}

	tp, err := os.ReadFile(protoErrPath)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("go").Parse(strings.TrimSpace(string(tp)))
	if err != nil {
		return nil, err
	}

	if err := tmpl.Execute(buf, nil); err != nil {
		return nil, err
	}

	newError, err := p.scanError(buf.String())
	if err != nil {
		return nil, err
	}

	oldError.ErrorSort = append(oldError.ErrorSort, newError.ErrorSort...)
	for key, val := range newError.ErrorMap {
		oldError.ErrorMap[key] = val
	}
	oldError.ErrorSort = uniqueStrings(oldError.ErrorSort)

	return oldError, nil
}

func (p *proto) scanError(content string) (*protoError, error) {
	repo := &protoError{
		ErrorSort: []string{},
		ErrorMap:  make(map[string]string),
	}
	enumEntryRegex := regexp.MustCompile(`(\w+)\s*=\s*\d+\s*\[(.*?)\];`)
	matches := enumEntryRegex.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		if len(match) == 3 {
			repo.ErrorSort = append(repo.ErrorSort, match[1])
			repo.ErrorMap[match[1]] = fmt.Sprintf("[%s]", match[2])
		}
	}

	return repo, nil
}

func (p *proto) renderError(object *Object) (string, error) {
	tps, err := p.genError(object)
	if err != nil {
		return "", err
	}

	tpl := `syntax = "proto3";

package errors;

import "errors/errors.proto";
option go_package = "./;errors";

enum ErrorReason {
  // 设置缺省错误码
  option (errors.default_code) = 500;

%s
}`

	var lines []string
	for ind, name := range tps.ErrorSort {
		lines = append(lines, fmt.Sprintf("  %s = %d%s;", name, ind, tps.ErrorMap[name]))
	}

	return fmt.Sprintf(tpl, strings.Join(lines, "\n")), nil
}
