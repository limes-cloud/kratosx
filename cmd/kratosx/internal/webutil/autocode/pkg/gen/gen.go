package gen

import (
	"fmt"
	"os"
	"strings"

	"github.com/fsgo/go_fmt/gofmtapi"
	"golang.org/x/mod/modfile"
	"gorm.io/gorm"

	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen/types"
)

type Builder struct {
	Module   string
	NameRule string
	Server   string
	DB       *gorm.DB
	Table    *types.Table
	TplRoot  string
	SrvRoot  string
	WebRoot  string
}

func NewBuilder(db *gorm.DB, table *types.Table) *Builder {
	webRoot := "/Users/fangweiye/GolandProjects/go-platform/framework/kratos-layout"
	srvRoot := "/Users/fangweiye/GolandProjects/go-platform/framework/kratos-layout"
	tplRoot := "/Users/fangweiye/GolandProjects/go-platform/framework/kratosx/cmd/kratosx/internal/webutil/autocode/template"

	getMod := func(srvRoot string) string {
		modBytes, err := os.ReadFile(srvRoot + "/go.mod")
		if err == nil {
			return modfile.ModulePath(modBytes)
		}

		arr := strings.Split(srvRoot, "/")
		return arr[len(arr)-1]
	}

	getServer := func(srvRoot string) string {
		mod := getMod(srvRoot)
		arr := strings.Split(mod, "/")
		return arr[len(arr)-1]
	}

	return &Builder{
		NameRule: types.NameRuleHump,
		Module:   getMod(srvRoot),
		Server:   getServer(srvRoot),
		DB:       db,
		Table:    table,
		SrvRoot:  srvRoot,
		TplRoot:  tplRoot,
		WebRoot:  webRoot,
	}
}

// ProtoErrorPath 获取proto error代码地址
func (b *Builder) ProtoErrorPath() string {
	return b.SrvRoot + "/api/" + b.Server + "/errors/" + b.Server + "_error_reason.proto"
}

// ProtoErrorTplPath 获取proto error 模板地址
func (b *Builder) ProtoErrorTplPath() string {
	return b.TplRoot + "/proto/error.tpl"
}

// ProtoMessagePath 获取proto节后定义代码地址
func (b *Builder) ProtoMessagePath() string {
	filename := fmt.Sprintf("%s_%s.proto", b.Server, pkg.ToSnake(b.Table.Struct))
	return b.SrvRoot + "/api/" + b.Server + "/" + b.Table.Module + "/" + filename
}

// ProtoMessageTplPath 获取proto 结构定义模板地址
func (b *Builder) ProtoMessageTplPath() string {
	return b.TplRoot + "/proto/message.tpl"
}

// ProtoServicePath 获取proto service 代码地址
func (b *Builder) ProtoServicePath() string {
	filename := fmt.Sprintf("%s_%s_service.proto", b.Server, pkg.ToSnake(b.Table.Module))
	return b.SrvRoot + "/api/" + b.Server + "/" + b.Table.Module + "/" + filename
}

// ProtoServiceTplPath 获取proto service 模板地址
func (b *Builder) ProtoServiceTplPath() string {
	return b.TplRoot + "/proto/service.tpl"
}

// GoTypesPath 获取类型路径
func (b *Builder) GoTypesPath() string {
	return b.SrvRoot + "/internal/types/" + pkg.ToSnake(b.Table.Module) + ".go"
}

// GoRepoPath 获取repo路径
func (b *Builder) GoRepoPath() string {
	return b.SrvRoot + "/internal/domain/repository/" + pkg.ToSnake(b.Table.Module) + ".go"
}

// GoRepoTplPath 获取repo 模板路径
func (b *Builder) GoRepoTplPath() string {
	return b.TplRoot + "/go/repo.tpl"
}

// GoEntityPath 获取entity路径
func (b *Builder) GoEntityPath() string {
	return b.SrvRoot + "/internal/domain/entity/" + pkg.ToSnake(b.Table.Module) + ".go"
}

// GoEntityTplPath 获取entity模板路径
func (b *Builder) GoEntityTplPath() string {
	return b.TplRoot + "/go/entity.tpl"
}

// GoDbsPath 获取dbs路径
func (b *Builder) GoDbsPath() string {
	return b.SrvRoot + "/internal/infra/dbs/" + pkg.ToSnake(b.Table.Module) + ".go"
}

// GoDbsTplPath 获取dbs模板路径
func (b *Builder) GoDbsTplPath() string {
	return b.TplRoot + "/go/dbs.tpl"
}

// GoServicePath 获取service路径
func (b *Builder) GoServicePath() string {
	return b.SrvRoot + "/internal/domain/service/" + pkg.ToSnake(b.Table.Module) + ".go"
}

// GoServiceTplPath 获取service模板路径
func (b *Builder) GoServiceTplPath() string {
	return b.TplRoot + "/go/service.tpl"
}

// IsRelationTypeMany 是否存在多对多引用
func (b *Builder) IsRelationTypeMany(tp string) bool {
	return tp == types.RelationTypeMany
}

// GoAppPath 获取dbs路径
func (b *Builder) GoAppPath() string {
	return b.SrvRoot + "/internal/app/" + pkg.ToSnake(b.Table.Module) + ".go"
}

// GoAppTplPath 获取App模板路径
func (b *Builder) GoAppTplPath() string {
	return b.TplRoot + "/go/app.tpl"
}

// GoAppEntryPath 获取app路径
func (b *Builder) GoAppEntryPath() string {
	return b.SrvRoot + "/internal/app/entry.go"
}

// GoAppEntryTplPath 获取App模板路径
func (b *Builder) GoAppEntryTplPath() string {
	return b.TplRoot + "/go/entry.tpl"
}

// GoAppEntryTplPath 获取App模板路径
func (b *Builder) GoTsApiPath() string {
	return b.WebRoot + "/go/entry.tpl"
}

// HasDeletedAt 是否启用delete模板
func (b *Builder) HasDeletedAt() bool {
	for _, col := range b.Table.Columns {
		if col.IsDeletedAt() {
			return true
		}
	}
	return false
}

// FormatGoCode 格式化go代码
func (b *Builder) FormatGoCode(code string) string {
	format := gofmtapi.NewFormatter()
	options := gofmtapi.NewOptions()
	options.LocalModule = b.Module
	_, ac, _, err := format.Format("", []byte(code), options)
	if err != nil {
		return code
	}
	return string(ac)
}
