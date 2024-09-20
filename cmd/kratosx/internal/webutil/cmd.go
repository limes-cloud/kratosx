package webutil

import (
	"log"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/spf13/cobra"

	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/router"
)

var CmdWebUtil = &cobra.Command{
	Use:   "webutil",
	Short: "webutil start",
	Long:  "webutil the kratosx tools. Example: kratosx webutil 8080",
	Run:   Run,
}

func init() {
	// CmdAutoCode.Flags().StringVarP(&port, "port", "p", "9000", "server port")
}

func Run(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		panic("example: kratosx autocode 8080")
	}

	srv := NewServer(args[0])

	router.NewRouter(srv)

	app := kratos.New(
		kratos.Name("webutil"),
		kratos.Version("v1.0.0"),
		kratos.Server(),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

func NewServer(port string) *http.Server {
	return http.NewServer(
		http.Address(":"+port),
		http.Middleware(),
	)
}
