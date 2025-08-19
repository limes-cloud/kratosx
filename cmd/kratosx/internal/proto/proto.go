package proto

import (
	"github.com/spf13/cobra"

	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/proto/add"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/proto/client"
)

// CmdProto represents the proto command.
var CmdProto = &cobra.Command{
	Use:   "proto",
	Short: "Generate the proto files",
	Long:  "Generate the proto files.",
}

func init() {
	CmdProto.AddCommand(add.CmdAdd)
	CmdProto.AddCommand(client.CmdClient)
}
