package autocode

import (
	"encoding/json"
	"os"
	"strings"

	"golang.org/x/mod/modfile"
)

type Object struct {
	Server      string   `json:"Server"`      // 服务名
	Module      string   `json:"module"`      // 包
	Table       string   `json:"table"`       // 表
	Keyword     string   `json:"keyword"`     // 简写
	Comment     string   `json:"comment"`     // 备注
	Description string   `json:"description"` // 模块描述
	Type        string   `json:"type"`        // 类型，list/tree
	Fields      []*Field `json:"fields"`      // 所有字段
	Methods     []string `json:"methods"`     // 所有方法
}

type Field struct {
	Keyword   string         `json:"keyword"`   // 字段key
	Title     string         `json:"title"`     // 字段标题
	Type      string         `json:"type"`      // 字段类型
	Default   string         `json:"default"`   // 默认值
	Required  bool           `json:"required"`  // 是否必填
	Rules     map[string]any `json:"rules"`     // 校验规则
	Operation FieldOperation `json:"operation"` // 操作
	QueryType string         `json:"queryType"` // 查询方式
	Relation  *FieldRelation `json:"relation"`  // 关联模型
}

type FieldOperation struct {
	Create bool `json:"create"` // 插入
	Update bool `json:"update"` // 更新
	List   bool `json:"list"`   // 列表
	Get    bool `json:"get"`    // 查询
}

type FieldRelation struct {
	Type   string
	Object *Object
}

func GenByJson(conf []byte) (map[string]string, error) {
	var object = Object{}
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
		mdName := md + toUpperCamelCase(o.Keyword)
		curMp[mdName] = true
	}

	mp := make(map[string]bool)
	methods := []string{"Get", "List", "Create", "Import", "Export", "Update", "Delete", "BatchDelete"}
	for _, md := range methods {
		mdName := md + toUpperCamelCase(o.Keyword)
		mp[mdName] = curMp[mdName]
	}
	return mp
}

func (o *Object) HasMethod(mp map[string]bool, method string) bool {
	is, exist := mp[method]
	if !exist {
		return true
	}
	return is
}
