package code

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen/gocode"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen/proto"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen/types"
)

func TestCode(t *testing.T) {
	initTable := func() *types.Table {
		var table types.Table
		content, err := os.ReadFile("./comment.json")
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal(content, &table); err != nil {
			panic(err)
		}
		return &table
	}

	builder := gen.NewBuilder(nil, initTable())

	// proto 代码构造
	proto := proto.NewBuilder(builder)

	// 生成proto error
	protoErrorCode, err := proto.GenError()
	if err != nil {
		panic(err)
	}
	_ = pkg.WriteCode(builder.ProtoErrorPath(), protoErrorCode)

	// 生成proto message
	protoMsgCode, err := proto.GenMessage()
	if err != nil {
		fmt.Println("🚫 generate proto error error " + err.Error())
	}
	_ = pkg.WriteCode(builder.ProtoMessagePath(), protoMsgCode)

	if err := pkg.GenProtoGRpc(builder.SrvRoot, builder.ProtoMessagePath()); err != nil {
		fmt.Println("🚫 generate proto message error " + err.Error())
	}

	// 生成proto service
	protoSrvCode, err := proto.GenService()
	if err != nil {
		panic(err)
	}
	_ = pkg.WriteCode(builder.ProtoServicePath(), protoSrvCode)

	if err := pkg.GenProtoGRpc(builder.SrvRoot, builder.ProtoServicePath()); err != nil {
		fmt.Println("🚫 generate proto service error " + err.Error())
	}

	// go 代码构造
	gocode := gocode.NewBuilder(builder)

	// 生成types代码
	goTypesCode, err := gocode.GenTypes()
	if err != nil {
		fmt.Println("🚫 generate proto error error " + err.Error())
	}
	_ = pkg.WriteCode(builder.GoTypesPath(), goTypesCode)

	// 生成entity代码
	goEntityCode, err := gocode.GenEntity()
	if err != nil {
		fmt.Println("🚫 generate proto error error " + err.Error())
	}
	_ = pkg.WriteCode(builder.GoEntityPath(), goEntityCode)

	// 生成entity代码
	goRepoCode, err := gocode.GenRepo()
	if err != nil {
		fmt.Println("🚫 generate proto error error " + err.Error())
	}
	_ = pkg.WriteCode(builder.GoRepoPath(), goRepoCode)

	// 生成entity代码
	goDbsCode, err := gocode.GenDbs()
	if err != nil {
		fmt.Println("🚫 generate proto error error " + err.Error())
	}
	_ = pkg.WriteCode(builder.GoDbsPath(), goDbsCode)

	// 生成entity代码
	goSrvCode, err := gocode.GenService()
	if err != nil {
		fmt.Println("🚫 generate proto error error " + err.Error())
	}
	_ = pkg.WriteCode(builder.GoServicePath(), goSrvCode)

	// 生成entity代码
	goAppCode, err := gocode.GenApp()
	if err != nil {
		fmt.Println("🚫 generate proto error error " + err.Error())
	}
	_ = pkg.WriteCode(builder.GoAppPath(), goAppCode)

	// 生成entry代码
	goAppEntryCode, err := gocode.GenAppEntry()
	if err != nil {
		fmt.Println("🚫 generate proto error error " + err.Error())
	}
	_ = pkg.WriteCode(builder.GoAppEntryPath(), goAppEntryCode)

	// 生成ts代码
	// tsBuilder := web.NewTsBuilder(builder)
	// webTsCode, err := tsBuilder.GenTypeScript(builder.ProtoMessagePath(), builder.ProtoServicePath())
	// _ = pkg.WriteCode(builder.ProtoServicePath(), goAppEntryCode)
}
