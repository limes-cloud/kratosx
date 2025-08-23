package add

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"unicode"

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
	// 开源仓库特殊判断
	modArr := strings.Split(modName, "/")
	if _, err := url.Parse(modArr[0]); err == nil {
		modArr = modArr[2:]
	}

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
		RootName:    rootName(modArr[0]),
	}
	p.LService = toSnake(p.Service)
	p.LPService = toPluralize(p.LService)
	if err := p.Generate(); err != nil {
		log.Fatal(err)
		return
	}
}

func rootName(name string) string {
	rn := strings.ToLower(name)
	rn = strings.ReplaceAll(rn, "-", "")
	rn = strings.ReplaceAll(rn, "_", "")
	return rn
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

// ToPluralize 转换为复数形式
func toPluralize(word string) string {
	lastLetter := word[len(word)-1:]
	beforeLastLetter := word[len(word)-2 : len(word)-1]
	switch lastLetter {
	case "y":
		if beforeLastLetter == "a" || beforeLastLetter == "e" || beforeLastLetter == "i" || beforeLastLetter == "o" || beforeLastLetter == "u" {
			return word + "s"
		} else {
			return word[:len(word)-1] + "ies"
		}
	case "x", "s", "z", "o":
		return word + "es"
	case "h":
		if beforeLastLetter == "s" || beforeLastLetter == "c" {
			return word + "es"
		} else {
			return word + "s"
		}
	case "f":
		if beforeLastLetter == "f" {
			return word[:len(word)-2] + "ves"
		} else {
			return word[:len(word)-1] + "ves"
		}
	default:
		return word + "s"
	}
}

// toSnake 转换为下划线
func toSnake(s string) string {
	runes := []rune(s)
	length := len(runes)
	var result []rune

	for i, r := range runes {
		if unicode.IsUpper(r) && i > 0 && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
	}

	return strings.ToLower(string(result))
}
