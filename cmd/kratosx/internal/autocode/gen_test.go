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
				Relations: []*FieldRelation{
					{
						Type:   _relationHasMany,
						Object: classify,
					},
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
	data := []byte(`{
  "table":"user",
  "keyword": "user",
  "module": "user",
  "comment": "用户信息",
  "description": "用户信息",
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
      "relations":[
        {
          "type": "hasMany",
          "rules": {"min_items": 1},
          "object": {
            "table":"user_job",
            "keyword": "user_job",
            "module": "job",
            "comment": "用户职位信息",
            "description": "用户职位信息",
            "fields": [
              {
                "keyword": "job_id",
                "title": "职位id",
                "type": "foreignKey",
                "required": true,
                "default": "",
                "operation": {
                  "create": true,
                  "update": false,
                  "list": false,
                  "get": false
                }
              }
            ]
          }
        },
        {
          "type": "hasMany",
          "rules": {"min_items": 1},
          "object": {
            "table":"user_role",
            "keyword": "user_role",
            "module": "role",
            "comment": "用户角色信息",
            "description": "用户角色信息",
            "fields": [
              {
                "keyword": "role_id",
                "title": "角色id",
                "type": "foreignKey",
                "required": true,
                "default": "",
                "operation": {
                  "create": true,
                  "update": false,
                  "list": false,
                  "get": false
                }
              }
            ]
          }
        },
        {
          "type": "hasMany",
          "object": {
            "table":"role",
            "keyword": "role",
            "module": "role",
            "comment": "角色信息",
            "description": "角色信息",
            "fields": [
              {
                "keyword": "id",
                "title": "主键",
                "type": "primaryKey",
                "required": true,
                "default": "",
                "operation": {
                  "create": false,
                  "update": false,
                  "list": false,
                  "get": true
                }
              },
              {
                "keyword": "name",
                "title": "名称",
                "type": "varchar64",
                "required": true,
                "default": "",
                "operation": {
                  "create": false,
                  "update": false,
                  "list": false,
                  "get": true
                }
              },
              {
                "keyword": "keyword",
                "title": "标识",
                "type": "char",
                "required": true,
                "default": "",
                "operation": {
                  "create": false,
                  "update": false,
                  "list": false,
                  "get": true
                }
              }
            ]
          }
        },
        {
          "type": "hasMany",
          "object": {
            "table":"job",
            "keyword": "job",
            "module": "job",
            "comment": "职位信息",
            "description": "职位信息",
            "fields": [
              {
                "keyword": "id",
                "title": "主键",
                "type": "primaryKey",
                "required": true,
                "default": "",
                "operation": {
                  "create": false,
                  "update": false,
                  "list": false,
                  "get": true
                }
              },
              {
                "keyword": "name",
                "title": "名称",
                "type": "varchar64",
                "required": true,
                "default": "",
                "operation": {
                  "create": false,
                  "update": false,
                  "list": false,
                  "get": true
                }
              },
              {
                "keyword": "keyword",
                "title": "标识",
                "type": "char",
                "required": true,
                "default": "",
                "operation": {
                  "create": false,
                  "update": false,
                  "list": false,
                  "get": true
                }
              }
            ]
          }
        }
      ]
    },
    {
      "keyword": "department_id",
      "title": "部门id",
      "type": "foreignKey",
      "required": true,
      "default": "",
      "operation": {
        "create": true,
        "update": true,
        "list": false,
        "get": true
      },
      "queryType": "=",
      "relations":[
        {
          "type": "hasOne",
          "object": {
            "table":"department",
            "keyword": "department",
            "module": "department",
            "comment": "部门信息",
            "description": "部门信息",
            "fields": [
              {
                "keyword": "id",
                "title": "主键",
                "type": "primaryKey",
                "required": true,
                "default": "",
                "operation": {
                  "create": false,
                  "update": false,
                  "list": false,
                  "get": true
                }
              },
              {
                "keyword": "name",
                "title": "名称",
                "type": "varchar64",
                "required": true,
                "default": "",
                "operation": {
                  "create": false,
                  "update": false,
                  "list": false,
                  "get": true
                }
              },
              {
                "keyword": "keyword",
                "title": "标识",
                "type": "char",
                "required": true,
                "default": "",
                "operation": {
                  "create": false,
                  "update": false,
                  "list": false,
                  "get": true
                }
              }
            ]
          }
        }
      ]
    },
    {
      "keyword": "role_id",
      "title": "角色id",
      "type": "foreignKey",
      "required": true,
      "default": "",
      "operation": {
        "create": false,
        "update": false,
        "list": false,
        "get": true
      },
      "queryType": "=",
      "relations": [
        {
          "type": "hasOne",
          "object": {
            "table":"role",
            "keyword": "role",
            "module": "role",
            "comment": "角色信息",
            "description": "角色信息",
            "fields": [
              {
                "keyword": "id",
                "title": "主键",
                "type": "primaryKey",
                "required": true,
                "default": "",
                "operation": {
                  "create": false,
                  "update": false,
                  "list": false,
                  "get": true
                }
              },
              {
                "keyword": "name",
                "title": "名称",
                "type": "varchar64",
                "required": true,
                "default": "",
                "operation": {
                  "create": false,
                  "update": false,
                  "list": false,
                  "get": false
                }
              },
              {
                "keyword": "keyword",
                "title": "标识",
                "type": "char",
                "required": true,
                "default": "",
                "operation": {
                  "create": false,
                  "update": false,
                  "list": false,
                  "get": false
                }
              }
            ]
          }
        }
      ]
    },
    {
      "keyword": "name",
      "title": "姓名",
      "type": "char",
      "required": true,
      "default": "",
      "operation": {
        "create": true,
        "update": true,
        "list": true,
        "get": true
      },
      "queryType": "like"
    },
    {
      "keyword": "nickname",
      "title": "昵称",
      "type": "varchar64",
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
      "keyword": "gender",
      "title": "性别",
      "type": "char",
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
      "keyword": "avatar",
      "title": "头像",
      "type": "varchar256",
      "required": false,
      "default": "",
      "operation": {
        "create": false,
        "update": true,
        "list": true,
        "get": true
      }
    },
    {
      "keyword": "phone",
      "title": "电话",
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
      "keyword": "email",
      "title": "邮箱",
      "type": "varchar64",
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
      "keyword": "password",
      "title": "密码",
      "type": "char",
      "required": true,
      "default": "",
      "operation": {
        "create": false,
        "update": false,
        "list": false,
        "get": false
      }
    },
    {
      "keyword": "status",
      "title": "状态",
      "type": "bool",
      "required": true,
      "default": "",
      "operation": {
        "create": false,
        "update": true,
        "list": true,
        "get": true
      }
    },
    {
      "keyword": "setting",
      "title": "用户设置",
      "type": "tinytext",
      "required": false,
      "default": "",
      "operation": {
        "create": false,
        "update": false,
        "list": false,
        "get": true
      }
    },
    {
      "keyword": "token",
      "title": "用户token",
      "type": "tinytext",
      "required": false,
      "default": "",
      "operation": {
        "create": false,
        "update": false,
        "list": false,
        "get": false
      }
    },
    {
      "keyword": "logged_at",
      "title": "登陆时间",
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
      },
      "queryType": "between"
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
    }
  ],
  "unique": [
    ["phone"],
    ["email"]
  ],
  "index": [
    ["logged_at"],
    ["created_at"],
    ["updated_at"]
  ],
  "methods": ["Get","List", "Create", "Update","UpdateStatus", "Delete"]
}`)
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
