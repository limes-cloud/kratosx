package gen

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/base"
	"github.com/spf13/cobra"
)

// CmdClient represents the source command.
var CmdClient = &cobra.Command{
	Use:   "gen",
	Short: "Generate the proto grpc code",
	Long:  "Generate the proto grpc code. Example: kratos proto client helloworld.proto",
	Run:   cRun,
}

var (
	protoPath string
	outPath   string
)

func init() {
	if protoPath = os.Getenv("KRATOS_PROTO_PATH"); protoPath == "" {
		protoPath = "./third_party"
	}

	if outPath = os.Getenv("KRATOS_OUT_PATH"); outPath == "" {
		outPath = "."
	}

	CmdClient.Flags().StringVarP(&protoPath, "proto_path", "p", protoPath, "proto path")
	CmdClient.Flags().StringVarP(&outPath, "out_path", "o", outPath, "out path")
}

func cRun(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Please enter the proto file or directory")
		return
	}
	var (
		err   error
		proto = strings.TrimSpace(args[0])
	)
	if err = look("protoc-gen-go", "protoc-gen-go-grpc", "protoc-gen-go-http", "protoc-gen-go-errors", "protoc-gen-openapi"); err != nil {
		// update the kratos plugins
		cmd := exec.Command("kratos", "upgrade")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Dir = "."
		if err = cmd.Run(); err != nil {
			fmt.Println(err)
			return
		}
	}

	// if outPath == "." {
	//	//lastIndex := strings.LastIndex(proto, "/")
	//	//outPath = proto[:lastIndex]
	// }

	if strings.HasSuffix(proto, ".proto") {
		err = generate(proto, args)
	} else {
		err = walk(proto, args)
	}
	if err != nil {
		fmt.Println(err)
	}
}

func look(name ...string) error {
	for _, n := range name {
		if _, err := exec.LookPath(n); err != nil {
			return err
		}
	}
	return nil
}

func walk(dir string, args []string) error {
	if dir == "" {
		dir = "."
	}
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if ext := filepath.Ext(path); ext != ".proto" || strings.HasPrefix(path, "third_party") {
			return nil
		}
		return generate(path, args)
	})
}

// generate is used to execute the generate command for the specified proto file
func generate(proto string, args []string) error {
	input := []string{
		"--proto_path=.",
		"--proto_path=" + outPath,
	}
	if pathExists(protoPath) {
		input = append(input, "--proto_path="+protoPath)
	}

	fmt.Println(outPath)
	inputExt := []string{
		"--proto_path=" + base.KratosMod(),
		"--proto_path=" + filepath.Join(base.KratosMod(), "third_party"),
		"--go_out=paths=source_relative:" + outPath,
		"--go-grpc_out=paths=source_relative:" + outPath,
		"--go-http_out=paths=source_relative:" + outPath,
		"--go-errors_out=paths=source_relative:" + outPath,
		"--openapi_out=paths=source_relative:" + outPath,
		"--validate_out=lang=go,paths=source_relative:" + outPath,

		// "--go_out=" + outPath,
		// "--go-grpc_out=" + outPath,
		// "--go-http_out=" + outPath,
		// "--go-errors_out=" + outPath,
		// "--openapi_out=" + outPath,
		// "--validate_out=lang=go:" + outPath,
	}
	input = append(input, inputExt...)
	protoBytes, err := os.ReadFile(proto)
	if err == nil && len(protoBytes) > 0 {
		if ok, _ := regexp.Match(`\n[^/]*(import)\s+"validate/validate.proto"`, protoBytes); ok {
			input = append(input, "--validate_out=lang=go,paths=source_relative:"+outPath)
		}
	}
	input = append(input, proto)
	for _, a := range args {
		if strings.HasPrefix(a, "-") {
			input = append(input, a)
		}
	}
	fd := exec.Command("protoc", input...)
	fd.Stdout = os.Stdout
	fd.Stderr = os.Stderr
	fd.Dir = "."
	if err := fd.Run(); err != nil {
		return err
	}
	fmt.Printf("proto: %s\n", proto)
	return nil
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
