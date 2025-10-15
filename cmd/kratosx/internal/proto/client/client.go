package client

import (
	"fmt"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/base"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/pkg"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cobra"
)

// CmdClient represents the source command.
var CmdClient = &cobra.Command{
	Use:   "client",
	Short: "Generate the proto client code",
	Long:  "Generate the proto client code. Example: kratos proto client helloworld.proto",
	Run:   run,
}

var (
	protoPath     string
	outPath       string
	projectPath   string
	filterModPath = []string{"api", "sdk", "rpc"}
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

func projectDir(path string) string {
	curDir, _ := os.Getwd()
	path = curDir + "/" + path

	// 判断当前是否存在环境文件
	for {
		// 出现go.mod 认为在根目录
		if _, err := os.Stat(filepath.Join(path, "go.mod")); err == nil {

			// 兼容内部包含api模式
			if !pkg.InList(filterModPath, filepath.Base(path)) {
				return path
			}
		}

		if path == "" || path == "/" {
			wp, _ := os.Getwd()
			return wp
		}

		// 往上移动一个目录
		path = filepath.Dir(path)
	}
}

func filterOverlap(path1, path2 string) string {
	// 移除路径末尾的斜杠
	path1 = strings.TrimSuffix(path1, "/")
	path2 = strings.TrimSuffix(path2, "/")

	// 如果path1是path2的前缀，则直接返回path2中不重合的部分
	if strings.HasPrefix(path2, path1) {
		return strings.TrimPrefix(path2, path1)
	}

	// 分割路径为部分
	parts1 := strings.Split(path1, "/")
	parts2 := strings.Split(path2, "/")

	// 查找path1与path2重叠的部分
	overlapIndex := 0
	for i := len(parts1) - 1; i >= 0; i-- {
		// 尝试将path1从当前部分开始与path2进行前缀匹配
		prefix := strings.Join(parts1[i:], "/")
		if strings.HasPrefix(path2, prefix) {
			overlapIndex = len(parts1) - i
		}
	}

	// 如果存在重叠，提取path2中不重叠的部分
	if overlapIndex != -1 {
		uniqueParts := parts2[overlapIndex:]
		return strings.Join(uniqueParts, "/")
	}

	// 如果没有重叠，返回完整的path2
	return path2
}

func run(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Please enter the proto file or directory")
		return
	}
	var (
		err   error
		proto = strings.TrimSpace(args[0])
	)
	if err = look("protoc-gen-go", "protoc-gen-go-grpc", "protoc-gen-go-httpx", "protoc-gen-go-errorsx", "protoc-gen-openapi"); err != nil {
		// update the kratos plugins
		cmd := exec.Command("kratosx", "upgrade")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Dir = "."
		if err = cmd.Run(); err != nil {
			fmt.Println(err)
			return
		}
	}

	//if outPath == "." {
	//	lastIndex := strings.LastIndex(proto, "/")
	//	outPath = proto[:lastIndex]
	//}
	// 移除尾部的path
	//outPath = strings.TrimSuffix(outPath, "/proto")
	//outPath = strings.TrimSuffix(outPath, "/pb")

	projectPath = projectDir(proto)
	if projectPath != "" {
		outPath = filepath.Dir(projectPath)
		proto = filterOverlap(projectPath, proto)
	}

	if strings.HasSuffix(proto, ".proto") {
		err = generate(proto, args)
	} else {
		err = walk(proto, args)
	}
	if err != nil {
		log.Fatal(err)
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

func findProtoFiles(root string) []string {
	protoFiles := map[string]struct{}{}

	_ = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".proto") && !strings.Contains(path, "third_party") {
			paths := strings.Split(path, "/")
			if len(paths) > 1 {
				paths = paths[:len(paths)-1]
				protoFiles[strings.Join(paths, "/")] = struct{}{}
			}
		}
		return nil
	})

	var list []string
	for path := range protoFiles {
		list = append(list, path)
	}
	return list
}

// generate is used to execute the generate command for the specified proto file
func generate(proto string, args []string) error {
	fmt.Println("projectPath", projectPath, outPath)
	input := []string{
		"--proto_path=.",
		"--proto_path=" + projectPath,
	}
	if pathExists(protoPath) {
		input = append(input, "--proto_path="+protoPath)
	}

	protoDirs := findProtoFiles(".")
	for _, path := range protoDirs {
		input = append(input, "--proto_path="+path)
	}

	inputExt := []string{
		"--proto_path=" + base.KratosMod(projectPath),
		"--proto_path=" + filepath.Join(base.KratosMod(projectPath), "third_party"),
		"--proto_path=" + base.KratosxMod(projectPath),
		"--proto_path=" + filepath.Join(base.KratosxMod(projectPath), "third_party"),
		"--go_out=" + outPath,
		"--go-grpc_out=" + outPath,
		"--go-httpx_out=" + outPath,
		"--go-errorsx_out=" + outPath,
		"--openapi_out=" + outPath,
	}

	input = append(input, inputExt...)
	protoBytes, err := os.ReadFile(proto)
	if err == nil && len(protoBytes) > 0 {
		if ok, _ := regexp.Match(`\n[^/]*(import)\s+"validate/validate.proto"`, protoBytes); ok {
			input = append(input, "--validate_out=lang=go:"+outPath)
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
	fd.Dir = projectPath

	if err := fd.Run(); err != nil {
		return err
	}
	fmt.Printf("proto: %s\n", proto)
	return nil
}

//func generate(proto string, args []string) error {
//	input := []string{
//		"--proto_path=.",
//	}
//	if pathExists(protoPath) {
//		input = append(input, "--proto_path="+protoPath)
//	}
//	inputExt := []string{
//		"--proto_path=" + base.KratosMod(),
//		"--proto_path=" + filepath.Join(base.KratosMod(), "third_party"),
//		"--go_out=paths=source_relative:.",
//		"--go-grpc_out=paths=source_relative:.",
//		"--go-http_out=paths=source_relative:.",
//		"--go-errors_out=paths=source_relative:.",
//		"--openapi_out=paths=source_relative:.",
//	}
//	input = append(input, inputExt...)
//	protoBytes, err := os.ReadFile(proto)
//	if err == nil && len(protoBytes) > 0 {
//		if ok, _ := regexp.Match(`\n[^/]*(import)\s+"validate/validate.proto"`, protoBytes); ok {
//			input = append(input, "--validate_out=lang=go,paths=source_relative:.")
//		}
//	}
//	input = append(input, proto)
//	for _, a := range args {
//		if strings.HasPrefix(a, "-") {
//			input = append(input, a)
//		}
//	}
//	fd := exec.Command("protoc", input...)
//	fd.Stdout = os.Stdout
//	fd.Stderr = os.Stderr
//	fd.Dir = "."
//	if err := fd.Run(); err != nil {
//		return err
//	}
//	fmt.Printf("proto: %s\n", proto)
//	return nil
//}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
