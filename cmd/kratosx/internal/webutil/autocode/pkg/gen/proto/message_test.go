package proto

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen/types"
)

func initTable() *types.Table {
	var table types.Table
	content, err := os.ReadFile("internal/webutil/autocode/pkg/gen/auto.json")
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(content, &table); err != nil {
		panic(err)
	}
	return &table
}

func TestMessage(t *testing.T) {
	builder := NewMessageBuilder(gen.NewBuilder(nil, initTable()))
	code, err := builder.ScanMessage()
	if err != nil {
		panic(err)
	}
	fmt.Println(code)
}

func TestMessage_MakeMessage(t *testing.T) {
	builder := NewMessageBuilder(gen.NewBuilder(nil, initTable()))
	data, err := builder.GenMessage()
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
