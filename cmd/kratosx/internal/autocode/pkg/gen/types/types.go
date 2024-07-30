package gen

import (
	"gorm.io/gen/field"
)

const (
	TableTypeList = "list"
	TableTypeTree = "tree"
)

type Table struct {
	Type    string    `json:"type"`    // 类型
	Module  string    `json:"module"`  // 包名
	Name    string    `json:"name"`    // 表名
	Struct  string    `json:"struct"`  // 结构名
	Comment string    `json:"comment"` // 表备注
	Columns []*Column `json:"columns"` // 表字段列表
	Indexes []*Index  `json:"indexes"` // 表索引列表
}

type Column struct {
	Name       string            `json:"name"`        // 字段名
	Type       string            `json:"type"`        // 字段类型
	ColumnType string            `json:"column_type"` // 字段全称类型
	Size       string            `json:"size"`        // 字段大小
	Comment    string            `json:"comment"`     // 字段备注
	Collation  string            `json:"collation"`   // 字符编码
	IsNull     bool              `json:"is_null"`     // 是否为空
	Default    string            `json:"default"`     // 默认值
	Unsigned   bool              `json:"unsigned"`    // 是否无符号
	Relations  []*field.Relation `json:"relations"`   // 引用信息
	Operation  FieldOperation    `json:"operation"`   // 操作方式
	Rules      []string          `json:"rules"`       // 校验规则
}

type FieldOperation struct {
	Create bool `json:"create"` // 插入
	Update bool `json:"update"` // 更新
	List   bool `json:"list"`   // 列表
	Get    bool `json:"get"`    // 查询
}

type Relation struct {
	Table  Table  `json:"table"`  // 引用表
	Column string `json:"column"` // 引用字段
}

type Index struct {
	Names  []string `json:"name"`   // 索引名
	Unique bool     `json:"unique"` // 是否为唯一索引
}

// IsRelation 是否存在引用
func (c *Column) IsRelation() bool {
	return len(c.Relations) == 0
}

// GoType go语言类型
//func (c *Column) GoType() string {
//	switch c.Type {
//	case "numeric":
//		return "int32"
//	case "integer":
//		return "int32"
//	case "int":
//		return "int32"
//	case "smallint":
//		return "int32"
//	case "mediumint":
//		return "int32"
//	case "bigint":
//		return "int64"
//	case "float":
//		return "float32"
//	case "real":
//		return "float64"
//	case "double":
//		return "float64"
//	case "decimal":
//		return "float64"
//	case "char":
//		return "string"
//	case "varchar":
//		return "string"
//	case "tinytext":
//		return "string"
//	case "mediumtext":
//		return "string"
//	case "longtext":
//		return "string"
//	case "binary":
//		return "[]byte"
//	case "varbinary":
//		return "[]byte"
//	case "tinyblob":
//		return "[]byte"
//	case "blob":
//		return "[]byte"
//	case "mediumblob":
//		return "[]byte"
//	case "longblob":
//		return "[]byte"
//	case "text":
//		return "string"
//	case "json":
//		return "string"
//	case "enum":
//		return "string"
//	case "time":
//		return "time.Time"
//	case "date":
//		return "time.Time"
//	case "datetime":
//		return "time.Time"
//	case "timestamp":
//		return "time.Time"
//	case "year":
//		return "int32"
//	case "bit":
//		return "[]uint8"
//	case "boolean":
//		return "bool"
//	case "tinyint":
//		if strings.HasPrefix(strings.TrimSpace(c.ColumnType), "tinyint(1)") || c.Size == "1" {
//			return "bool"
//		}
//		return "int32"
//	default:
//		return "string"
//	}
//
//}
