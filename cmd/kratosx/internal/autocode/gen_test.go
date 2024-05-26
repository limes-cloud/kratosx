package autocode

import (
	"os"
	"testing"
)

func TestGen(t *testing.T) {
	classify := &Object{
		Server:      serverName(),
		Table:       "notice_classify",
		Keyword:     "classify",
		Module:      "notice",
		Comment:     "通知分类",
		Description: "主要负责通知的分类",
		Type:        _objectTypeList,
		Fields: []*Field{
			{
				Keyword:  "id",
				Title:    "主键",
				Type:     _primaryKey,
				Default:  "1",
				Required: true,
				Operation: FieldOperation{
					Create: true,
					Update: true,
					List:   true,
					Get:    true,
				},
				QueryType: "=",
			},
			{
				Keyword:  "title",
				Title:    "标题",
				Type:     _varchar64,
				Default:  "",
				Required: true,
				Operation: FieldOperation{
					Create: true,
					Update: true,
					List:   true,
					Get:    true,
				},
				QueryType: "like",
			},
		},
	}

	notice := Object{
		Server:      serverName(),
		Table:       "table_notice",
		Keyword:     "notice",
		Module:      "notice",
		Comment:     "通知",
		Description: "主要负责通知",
		Type:        _objectTypeList,
		Fields: []*Field{
			{
				Keyword:  "id",
				Title:    "主键",
				Type:     _primaryKey,
				Default:  "1",
				Required: true,
				Operation: FieldOperation{
					Create: true,
					Update: true,
					List:   true,
					Get:    true,
				},
				QueryType: "=",
			},
			{
				Keyword:  "classify_id",
				Title:    "分类id",
				Type:     _foreignKey,
				Default:  "1",
				Required: true,
				Operation: FieldOperation{
					Create: true,
					Update: true,
					List:   true,
					Get:    true,
				},
				QueryType: "=",
				Relation: &FieldRelation{
					Type:   _relationHasMany,
					Object: classify,
				},
			},
			{
				Keyword:  "title",
				Title:    "标题",
				Type:     _varchar64,
				Default:  "",
				Required: true,
				Operation: FieldOperation{
					Create: true,
					Update: true,
					List:   true,
					Get:    true,
				},
				QueryType: "like",
			},
			{
				Keyword:  "created_at",
				Title:    "创建时间",
				Type:     _time,
				Default:  "",
				Required: true,
				Operation: FieldOperation{
					Create: true,
					Update: true,
					List:   true,
					Get:    true,
				},
				QueryType: _between,
			},
		},
		Methods: []string{"Get", "List", "Create", "Import", "Export", "Update", "Delete", "BatchDelete"},
	}

	reply, err := Gen(&notice)
	for path, text := range reply {
		autoMkDir(path)
		if err := os.WriteFile(path, []byte(text), os.ModePerm); err != nil {
			t.Error(err)
			return
		}
	}
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(reply)
}

func TestGenJson(t *testing.T) {
	data := []byte(`
{
  "table":"dictionary",
  "keyword": "dictionary",
  "module": "dictionary",
  "comment": "字典目录",
  "description": "字典目录",
  "type": "list",
  "fields": [
    {
      "keyword": "id",
      "title": "主键",
      "type": "primaryKey",
      "required": true,
      "default": "",
      "operation": {
        "create": false,
        "update": true,
        "list": true,
        "get": true
      },
      "relation":{
        "type": "hasMany",
        "object": {
          "table": "dictionary_value",
          "keyword": "dictionary_value",
          "module": "dictionary",
          "comment": "字典值目录",
          "description": "字典值目录",
          "type": "list",
          "fields": [
            {
              "keyword": "label",
              "title": "标签",
              "type": "varchar128",
              "required": true,
              "default": "",
              "operation": {
                "create": true,
                "update": true,
                "list": false,
                "get": true
              }
            },
            {
              "keyword": "value",
              "title": "标识",
              "type": "varchar64",
              "required": true,
              "default": "",
              "operation": {
                "create": true,
                "update": true,
                "list": false,
                "get": true
              }
            },
            {
              "keyword": "type",
              "title": "类型",
              "type": "char",
              "required": false,
              "default": "",
              "operation": {
                "create": true,
                "update": true,
                "list": false,
                "get": true
              }
            },
            {
              "keyword": "extra",
              "title": "扩展数据",
              "type": "tinytext",
              "required": false,
              "default": "",
              "operation": {
                "create": true,
                "update": true,
                "list": false,
                "get": true
              }
            }
          ]
        }
      }
    },
    {
      "keyword": "keyword",
      "title": "目录标识",
      "type": "char",
      "required": true,
      "default": "",
      "operation": {
        "create": true,
        "update": true,
        "list": true,
        "get": true
      },
      "queryType": "="
    },
    {
      "keyword": "name",
      "title": "目录名称",
      "type": "varchar64",
      "required": true,
      "default": "",
      "operation": {
        "create": true,
        "update": true,
        "list": true,
        "get": true
      }
    },
    {
      "keyword": "description",
      "title": "目录描述",
      "type": "varchar256",
      "required": false,
      "default": "",
      "operation": {
        "create": true,
        "update": true,
        "list": true,
        "get": true
      }
    },
    {
      "keyword": "created_at",
      "title": "创建时间",
      "type": "time",
      "required": true,
      "default": "",
      "operation": {
        "create": false,
        "update": false,
        "list": true,
        "get": true
      }
    },
    {
      "keyword": "updated_at",
      "title": "更新时间",
      "type": "time",
      "required": true,
      "default": "",
      "operation": {
        "create": false,
        "update": false,
        "list": true,
        "get": true
      }
    },
    {
      "keyword": "deleted_at",
      "title": "删除时间",
      "type": "time",
      "required": true,
      "default": "",
      "operation": {
        "create": false,
        "update": false,
        "list": true,
        "get": true
      }
    }
  ],
"unique": [["keyword","deleted_at"]],
  "methods": ["Get", "List", "Create", "Import", "Export", "Update", "Delete", "GetTrash", "ListTrash", "DeleteTrash", "RevertTrash"]
}
`)
	reply, err := GenByJson(data)
	for path, text := range reply {
		autoMkDir(path)
		if err := os.WriteFile(path, []byte(text), os.ModePerm); err != nil {
			t.Error(err)
			return
		}
	}
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(reply)
}
