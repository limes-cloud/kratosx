package proto

import (
	"fmt"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/autocode/pkg/gen"
	"testing"
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
	builder := NewServerBuilder(gen.NewBuilder(nil, initTable()))
	data, err := builder.GenServer()
	if err != nil {
		panic(err)
	}
	fmt.Println(data)

}
