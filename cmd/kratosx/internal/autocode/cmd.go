package autocode

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// CmdAutoCode run project command.
var CmdAutoCode = &cobra.Command{
	Use:   "autocode",
	Short: "Auto Create Code",
	Long:  "Auto Create Code: kratosx autocode -f/-s",
	Run:   Run,
}

var targetFile string

func init() {
	CmdAutoCode.Flags().StringVarP(&targetFile, "file", "f", "", "target file path")
}

func splitArgs(cmd *cobra.Command, args []string) (cmdArgs, programArgs []string) {
	dashAt := cmd.ArgsLenAtDash()
	if dashAt >= 0 {
		return args[:dashAt], args[dashAt:]
	}
	return args, []string{}
}

// Run run project.
func Run(cmd *cobra.Command, args []string) {
	wd, _ := os.Getwd()
	if !strings.HasSuffix(wd, "/") {
		wd += "/"
	}

	var path string
	cmdArgs, _ := splitArgs(cmd, args)
	if len(cmdArgs) > 0 {
		path = wd + strings.TrimSpace(cmdArgs[0])
	}
	if path == "" {
		fmt.Println("🚫 path not exist")
		return
	}

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("🚫 " + err.Error())
		return
	}

	reply, err := GenByJson(data)
	if err != nil {
		fmt.Println("🚫 generate code error " + err.Error())
		return
	}
	var lastProto []string
	for p, text := range reply {
		autoMkDir(p)
		if err := os.WriteFile(p, []byte(text), os.ModePerm); err != nil {
			fmt.Println("🚫 save generate code error " + err.Error())
			return
		}
		if strings.HasSuffix(p, "service.proto") {
			lastProto = append(lastProto, p)
			continue
		}
		if strings.HasSuffix(p, ".proto") {
			fd := exec.Command("kratosx", "proto", "client", p)
			fd.Stdout = os.Stdout
			fd.Stderr = os.Stderr
			fd.Dir = "."
			if err = fd.Run(); err != nil {
				fmt.Println("🚫 generate proto code error " + err.Error())
			}
		}

	}
	for _, p := range lastProto {
		fd := exec.Command("kratosx", "proto", "client", p)
		fd.Stdout = os.Stdout
		fd.Stderr = os.Stderr
		fd.Dir = "."
		if err = fd.Run(); err != nil {
			fmt.Println("🚫 generate proto code error " + err.Error())
		}
	}
	fmt.Println("✅ generate success!!!")
}
