package autocode

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CmdAutoCode = &cobra.Command{
	Use:   "autocode",
	Short: "auto generate code",
	Long:  "autocode the kratosx tools. Example: kratosx autocode 8080",
	Run:   Run,
}

var (
	port string
)

func init() {
	//CmdAutoCode.Flags().StringVarP(&port, "port", "p", "9000", "server port")
}

func Run(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		panic("example: kratosx autocode 8080")
	}

	port = args[0]

	fmt.Println(port)
}
