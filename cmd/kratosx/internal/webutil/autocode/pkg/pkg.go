package pkg

import (
	"os"
	"os/exec"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// IsExistFile 判断是否存在文件
func IsExistFile(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return !fileInfo.IsDir()
}

// IsExistFolder 判断是否存在目录
func IsExistFolder(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return fileInfo.IsDir()
}

// AutoMakeFolder 自动创建目录
func AutoMakeFolder(fp string) {
	if strings.LastIndex(fp, "/") < strings.LastIndex(fp, ".") {
		fp = fp[:strings.LastIndex(fp, "/")]
	}

	if IsExistFolder(fp) {
		return
	}
	_ = os.MkdirAll(fp, 0777)
}

// In 对比接口
type In[T ListType] interface {
	Has(T) bool
}

type _comparable[T ListType] struct {
	m map[T]struct{}
}

type ListType interface {
	~string | ~int | ~uint32 | ~rune | ~float64 | ~int64 | ~float32
}

// InList 判断是否存在数组内
func InList[ListType comparable](list []ListType, val ListType) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}

func New[T ListType](list []T) In[T] {
	m := make(map[T]struct{})
	for _, item := range list {
		m[item] = struct{}{}
	}
	return &_comparable[T]{
		m: m,
	}
}

func (c *_comparable[T]) Has(t T) bool {
	_, ok := c.m[t]
	return ok
}

// ToPluralize 转换为复数形式
func ToPluralize(word string) string {
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

// ToSnake 转换为下划线
func ToSnake(s string) string {
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

// ToHump 转换为驼峰
func ToHump(s string) string {
	s = strings.ReplaceAll(s, "_", " ")
	s = cases.Title(language.Und, cases.NoLower).String(s)
	return strings.ReplaceAll(s, " ", "")
}

// ToLowerHump 转换首字母大写的驼峰
func ToLowerHump(s string) string {
	s = ToHump(s)
	prefix := strings.ToLower(string(s[0]))
	return prefix + s[1:]
}

// ToUpperHump 转换首字母大写的驼峰
func ToUpperHump(s string) string {
	s = ToHump(s)
	prefix := strings.ToUpper(string(s[0]))
	return prefix + s[1:]
}

func VariableName(name string, rule string) string {
	if rule == "hump" {
		return ToLowerHump(name)
	}
	return ToSnake(name)
}

func UniqueArray[T ListType](arr []T) []T {
	var (
		result []T
		bucket = make(map[T]struct{})
	)

	for _, str := range arr {
		if _, exists := bucket[str]; !exists {
			bucket[str] = struct{}{}
			result = append(result, str)
		}
	}

	return result
}

func WriteCode(path string, code string) error {
	AutoMakeFolder(path)
	return os.WriteFile(path, []byte(code), 0666)
}

func GenProtoGRpc(work string, path string) error {
	path = strings.TrimPrefix(path, work)
	path = strings.TrimPrefix(path, "/")
	fd := exec.Command("kratosx", "proto", "client", path)
	fd.Stdout = os.Stdout
	fd.Stderr = os.Stderr
	fd.Dir = work
	return fd.Run()
}
