package model

import (
	"github.com/limes-cloud/kratosx/types"
)

type {{.Object}} struct {
{{.Fields}}
}

{{if .IsTree}}
type {{.Object}}Closure struct {
	ID       uint32 `json:"id" gorm:"column:id"`
	Parent   uint32 `json:"parent" gorm:"column:parent"`
	Children uint32 `json:"children" gorm:"column:children"`
}
{{end}}