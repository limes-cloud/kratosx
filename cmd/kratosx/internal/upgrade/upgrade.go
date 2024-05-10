package upgrade

import (
	"fmt"

	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/base"
	"github.com/spf13/cobra"
)

// CmdUpgrade represents the upgrade command.
var CmdUpgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the kratosx tools",
	Long:  "Upgrade the kratosx tools. Example: kratosx upgrade",
	Run:   Run,
}

// Run upgrade the kratos tools.
func Run(_ *cobra.Command, _ []string) {
	err := base.GoInstall(
		"github.com/limes-cloud/kratosx/cmd/kratosx@latest",
		"github.com/limes-cloud/kratosx/cmd/protoc-gen-go-httpx@latest",
		"github.com/limes-cloud/kratosx/cmd/protoc-gen-go-errorsx@latest",
		"google.golang.org/protobuf/cmd/protoc-gen-go@latest",
		"google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest",
		"github.com/google/gnostic/cmd/protoc-gen-openapi@latest",
		"github.com/envoyproxy/protoc-gen-validate@latest",
	)
	if err != nil {
		fmt.Println(err)
	}
}
