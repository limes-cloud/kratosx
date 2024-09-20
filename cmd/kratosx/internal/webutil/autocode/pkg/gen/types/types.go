package types

import (
	"strings"

	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg"
)

const (
	TableTypeList = "list"
	TableTypeTree = "tree"
)

const (
	NameRuleHump  = "hump"
	NameRuleSnake = "snake"
)

const (
	RelationTypeOne  = "one"
	RelationTypeMany = "many"
)

const (
	ColumnQueryTypeIn      = "in"
	ColumnQueryTypeNotIn   = "not in"
	ColumnQueryTypeBetween = "between"
)

type Table struct {
	Type              string    `json:"type"`              // 类型
	Module            string    `json:"module"`            // 包名
	Name              string    `json:"name"`              // 表名
	Struct            string    `json:"struct"`            // 结构名
	Comment           string    `json:"comment"`           // 表备注
	Columns           []*Column `json:"columns"`           // 表字段列表
	Indexes           []*Index  `json:"indexes"`           // 表索引列表
	EnableBatchDelete bool      `json:"enableBatchDelete"` // 是否开启批量删除
}

type Column struct {
	Name       string          `json:"name"`       // 字段名
	Type       string          `json:"type"`       // 字段类型
	ColumnType string          `json:"columnType"` // 字段全称类型
	Size       string          `json:"size"`       // 字段大小
	Comment    string          `json:"comment"`    // 字段备注
	Collation  string          `json:"collation"`  // 字符编码
	IsNull     bool            `json:"is_null"`    // 是否为空
	Default    string          `json:"default"`    // 默认值
	Unsigned   bool            `json:"unsigned"`   // 是否无符号
	Relations  []*Relation     `json:"relations"`  // 引用信息
	Operation  ColumnOperation `json:"operation"`  // 操作方式
	Rules      map[string]any  `json:"rules"`      // 校验规则
	Query      ColumnQuery     `json:"query"`      // 查询方式
}

type ColumnQuery struct {
	Type string `json:"type"`
}

type ColumnOperation struct {
	Create bool `json:"create"` // 插入
	Update bool `json:"update"` // 更新
	List   bool `json:"list"`   // 列表
	Get    bool `json:"get"`    // 查询
}

type Relation struct {
	Rules  map[string]any `json:"rules"`  // 校验规则
	Type   string         `json:"type"`   // 引用类型
	Column string         `json:"column"` // 引用字段
	Table  *Table         `json:"table"`  // 引用表
}

type Index struct {
	Names  []string `json:"names"`  // 索引名
	Unique bool     `json:"unique"` // 是否为唯一索引
}

// IsRelation 是否存在引用
func (c *Column) IsRelation() bool {
	return len(c.Relations) == 0
}

func (c *Column) IsProtoOption() bool {
	if c.IsNull || c.ProtoType() == "bool" {
		return true
	}
	return false
}

func (c *Column) IsDeletedAt() bool {
	return c.Name == "deleted_at"
}

func (c *Column) ProtoType() string {
	switch c.GoType() {
	case "int32":
		return "int32"
	case "[]uint8", "[]byte":
		return "bytes"
	case "float32", "float64":
		return "float"
	case "time.Time":
		return "google.protobuf.Timestamp"
	case "int64", "uint64", "uint32":
		return "uint32"
	case "bool":
		return "bool"
	default:
		return "string"
	}
}

// GoType go语言类型
func (c *Column) GoType() string {
	switch c.Type {
	case "numeric":
		return "int32"
	case "integer":
		return "int32"
	case "int":
		return "int32"
	case "smallint":
		return "int32"
	case "mediumint":
		return "int32"
	case "bigint":
		return "uint32"
	case "float":
		return "float32"
	case "real":
		return "float64"
	case "double":
		return "float64"
	case "decimal":
		return "float64"
	case "char":
		return "string"
	case "varchar":
		return "string"
	case "tinytext":
		return "string"
	case "mediumtext":
		return "string"
	case "longtext":
		return "string"
	case "binary":
		return "[]byte"
	case "varbinary":
		return "[]byte"
	case "tinyblob":
		return "[]byte"
	case "blob":
		return "[]byte"
	case "mediumblob":
		return "[]byte"
	case "longblob":
		return "[]byte"
	case "text":
		return "string"
	case "json":
		return "string"
	case "enum":
		return "string"
	case "time":
		return "time.Time"
	case "date":
		return "time.Time"
	case "datetime":
		return "time.Time"
	case "timestamp":
		return "time.Time"
	case "year":
		return "int32"
	case "bit":
		return "[]uint8"
	case "boolean", "bool":
		return "bool"
	case "tinyint":
		if strings.HasPrefix(strings.TrimSpace(c.ColumnType), "tinyint(1)") || c.Size == "1" {
			return "bool"
		}
		return "int32"
	default:
		return "string"
	}
}

func (c *ColumnQuery) IsPluralize() bool {
	var ps = []string{
		ColumnQueryTypeIn,
		ColumnQueryTypeNotIn,
		ColumnQueryTypeBetween,
	}

	return pkg.InList(ps, c.Type)
}
