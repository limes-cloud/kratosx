package add

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/base"
)

// CmdAdd represents the add command.
var CmdAdd = &cobra.Command{
	Use:   "add",
	Short: "Add a proto API template",
	Long:  "Add a proto API template. Example: kratosx proto add helloworld/v1/hello.proto",
	Run:   run,
}

func run(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Please enter the proto file or directory")
		return
	}
	input := args[0]
	n := strings.LastIndex(input, "/")
	if n == -1 {
		fmt.Println("The proto path needs to be hierarchical.")
		return
	}
	path := input[:n]
	fileName := input[n+1:]
	// api/bxxx/xxx
	pkgName := strings.ReplaceAll(path, "/", ".")
	modName := base.ModName(path)
	modArr := strings.Split(modName, "/")
	for _, v := range modArr {
		if strings.HasPrefix(pkgName, v) {
			pkgName = strings.ReplaceAll(pkgName, v+".", "")
		}
	}
	pkgName = strings.Join(modArr, ".") + "." + pkgName

	p := &Proto{
		Name:        fileName,
		Path:        path,
		Package:     pkgName,
		GoPackage:   goPackage(path),
		JavaPackage: javaPackage(pkgName),
		Service:     serviceName(fileName),
	}
	if err := p.Generate(); err != nil {
		log.Fatal(err)
		return
	}
}

//	func modName() string {
//		modBytes, err := os.ReadFile("go.mod")
//		if err != nil {
//			if modBytes, err = os.ReadFile("../go.mod"); err != nil {
//				return ""
//			}
//		}
//		return modfile.ModulePath(modBytes)
//	}
func goPackage(path string) string {
	dir, _ := os.Getwd()
	s := strings.Split(dir+"/"+path, "/")
	abs := "."
	pkg := s[len(s)-1]
	return abs + ";" + pkg
}

func javaPackage(name string) string {
	return name
}

func serviceName(name string) string {
	return toUpperCamelCase(strings.Split(name, ".")[0])
}

func toUpperCamelCase(s string) string {
	s = strings.ReplaceAll(s, "_", " ")
	s = cases.Title(language.Und, cases.NoLower).String(s)
	return strings.ReplaceAll(s, " ", "")
}
