package gen

import (
	"fmt"
	"strings"

	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen/types"
)

type sql struct {
	*gen.Builder
}

type SQLBuilder interface {
	GenTableSQL() string
}

func NewSQLBuilder(builder *gen.Builder) SQLBuilder {
	return &sql{
		Builder: builder,
	}
}

func (sql *sql) GenTableSQL() string {
	var columns []string
	for _, col := range sql.Table.Columns {
		columns = append(columns, sql.GenColumnsSQL(col))
	}

	for _, col := range sql.Table.Indexes {
		columns = append(columns, sql.GenIndexSQL(col))
	}

	return fmt.Sprintf(
		"CREATE TABLE `%s` (\n%s\n)ENGINE=InnoDB  CHARSET=utf8mb4 COMMENT='%s';",
		sql.Table.Name,
		strings.Join(columns, ",\n"),
		sql.Table.Comment,
	)
}

func (sql *sql) GenColumnsSQL(column *types.Column) string {
	if column.Name == "id" {
		return "`id` bigint unsigned NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键ID'"
	}
	if column.Name == "created_at" {
		return "`created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间'"
	}
	if column.Name == "updated_at" {
		return "`updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间'"
	}
	if column.Name == "deleted_at" {
		return "`deleted_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间'"
	}

	var args []string
	args = append(args, fmt.Sprintf("`%s`", column.Name))

	if column.ColumnType == "" {
		switch column.Type {
		case "numeric", "char", "varchar", "varbinary", "tinyint":
			column.ColumnType = fmt.Sprintf("%s(%s)", column.Type, column.Size)
		default:
			column.ColumnType = column.Type
		}
	}

	// 类型
	args = append(args, column.ColumnType)

	// 是否为无符号
	if column.Unsigned {
		args = append(args, "unsigned")
	}

	// 字符比较集
	if column.Collation != "" {
		args = append(args, "COLLATE "+column.Collation)
	}

	// 是否为空
	if !column.IsNull {
		args = append(args, "NOT NULL")
	}

	// 默认值
	if column.Default != "" {
		args = append(args, fmt.Sprintf("DEFAULT '%s'", column.Default))
	}

	// 备注
	args = append(args, fmt.Sprintf("COMMENT '%s'", column.Comment))

	return strings.Join(args, " ")
}

func (sql *sql) GenIndexSQL(index *types.Index) string {
	var args []string
	if index.Unique {
		args = append(args, "UNIQUE")
	}
	args = append(args, "INDEX")

	args = append(args, fmt.Sprintf("(%s)", strings.Join(index.Names, ",")))
	return strings.Join(args, " ")
}
