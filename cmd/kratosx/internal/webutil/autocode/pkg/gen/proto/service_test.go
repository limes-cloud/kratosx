package proto

import (
	"fmt"
	"testing"

	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen"
)

func TestServer(t *testing.T) {
	builder := NewMessageBuilder(gen.NewBuilder(nil, initTable()))
	code, err := builder.ScanMessage()
	if err != nil {
		panic(err)
	}
	fmt.Println(code)
}

func TestServer_MakeServer(t *testing.T) {
	builder := NewServiceBuilder(gen.NewBuilder(nil, initTable()))
	data, err := builder.GenService()
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
