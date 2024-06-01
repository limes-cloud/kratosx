package autocode

import (
	"encoding/json"
	"os"
	"strings"

	"golang.org/x/mod/modfile"
)

type Object struct {
	Server      string     `json:"Server"`      // 服务名
	Module      string     `json:"module"`      // 包
	Table       string     `json:"table"`       // 表
	Keyword     string     `json:"keyword"`     // 简写
	Comment     string     `json:"comment"`     // 备注
	Description string     `json:"description"` // 模块描述
	Type        string     `json:"type"`        // 类型，list/tree
	Fields      []*Field   `json:"fields"`      // 所有字段
	Methods     []string   `json:"methods"`     // 所有方法
	Unique      [][]string `json:"unique"`
	Index       [][]string `json:"index"`
}

type Field struct {
	Keyword   string           `json:"keyword"`   // 字段key
	Title     string           `json:"title"`     // 字段标题
	IsOrder   bool             `json:"order"`     // 是否排序
	Type      string           `json:"type"`      // 字段类型
	Default   string           `json:"default"`   // 默认值
	Required  bool             `json:"required"`  // 是否必填
	Rules     map[string]any   `json:"rules"`     // 校验规则
	Operation FieldOperation   `json:"operation"` // 操作
	QueryType string           `json:"queryType"` // 查询方式
	Relations []*FieldRelation `json:"relations"` // 关联模型
}

type FieldOperation struct {
	Create bool `json:"create"` // 插入
	Update bool `json:"update"` // 更新
	List   bool `json:"list"`   // 列表
	Get    bool `json:"get"`    // 查询
}

type FieldRelation struct {
	Type   string
	Rules  map[string]any
	Object *Object
}

func GenByJson(conf []byte) (map[string]string, error) {
	var object = Object{Server: serverName()}
	if err := json.Unmarshal(conf, &object); err != nil {
		return nil, err
	}
	return Gen(&object)
}

func Gen(object *Object) (map[string]string, error) {
	reply := make(map[string]string)

	protoReply, err := GenProto(object)
	if err != nil {
		return nil, err
	}
	for key, val := range protoReply {
		reply[key] = val
	}

	bizReply, err := GenBiz(object)
	if err != nil {
		return nil, err
	}
	for key, val := range bizReply {
		reply[key] = val
	}

	srvReply, err := GenService(object)
	if err != nil {
		return nil, err
	}
	for key, val := range srvReply {
		reply[key] = val
	}

	repoReply, err := GenRepo(object)
	if err != nil {
		return nil, err
	}
	for key, val := range repoReply {
		reply[key] = val
	}

	return reply, nil
}

func serverName() string {
	modBytes, err := os.ReadFile("go.mod")
	if err != nil {
		if modBytes, err = os.ReadFile("../go.mod"); err != nil {
			return ""
		}
	}
	return modfile.ModulePath(modBytes)
}

func (o *Object) StructName() string {
	return toUpperCamelCase(o.Keyword)
}

func (o *Object) ServerName() string {
	arr := strings.Split(o.Server, "/")
	return arr[len(arr)-1]
}

func (o *Object) MethodStatus() map[string]bool {
	curMp := make(map[string]bool)
	for _, md := range o.Methods {
		if md == "UpdateStatus" {
			mdName := "Update" + toUpperCamelCase(o.Keyword) + "Status"
			curMp[mdName] = true
		} else {
			mdName := md + toUpperCamelCase(o.Keyword)
			curMp[mdName] = true
		}
	}

	mp := make(map[string]bool)
	methods := []string{"Get", "List", "Create", "Import", "Export", "Update", "Delete",
		"UpdateStatus", "ListTrash", "GetTrash", "ListTrash", "DeleteTrash", "RevertTrash"}

	treeMethods := []string{
		"Get" + toUpperCamelCase(o.Keyword) + "ChildrenIds",
		"Get" + toUpperCamelCase(o.Keyword) + "ParentIds",
		"append" + toUpperCamelCase(o.Keyword) + "Children",
		"remove" + toUpperCamelCase(o.Keyword) + "Parent",
	}
	for _, tm := range treeMethods {
		mp[tm] = o.Type == _objectTypeTree
	}

	for _, md := range methods {
		if md == "UpdateStatus" {
			mdName := "Update" + toUpperCamelCase(o.Keyword) + "Status"
			mp[mdName] = curMp[mdName]
		} else {
			mdName := md + toUpperCamelCase(o.Keyword)
			mp[mdName] = curMp[mdName]
		}
	}
	return mp
}

func (o *Object) HasMethod(mp map[string]bool, method string) bool {
	if strings.Contains(method, "By") && strings.HasPrefix(method, "Get") {
		method = method[:strings.Index(method, "By")]
	}
	is, exist := mp[method]
	if !exist {
		return true
	}
	return is
}

func (o *Object) FieldMap() map[string]*Field {
	var m = map[string]*Field{}
	for _, field := range o.Fields {
		m[field.Keyword] = field
	}
	return m
}
