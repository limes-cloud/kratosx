package entity

import (
    {{if .IsTree}}
    "github.com/limes-cloud/kratosx/pkg/tree"
    {{end}}
    "github.com/limes-cloud/kratosx/types"
)


{{range $key,$val := .Entities}}
type {{$key}} struct {
{{$val}}
}
{{end}}

{{if .IsTree}}
// ID 获取树ID
func (m *{{.Object}}) ID() uint32 {
	return m.Id
}

// Parent 获取父ID
func (m *{{.Object}}) Parent() uint32 {
	return m.ParentId
}

// AppendChildren 添加子节点
func (m *{{.Object}}) AppendChildren(child any) {
	menu := child.(*{{.Object}})
	m.Children = append(m.Children, menu)
}

// ChildrenNode 获取子节点
func (m *{{.Object}}) ChildrenNode() []tree.Tree {
	var list []tree.Tree
	for _, item := range m.Children {
		list = append(list, item)
	}
	return list
}
{{end}}