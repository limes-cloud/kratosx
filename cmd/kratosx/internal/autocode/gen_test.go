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
