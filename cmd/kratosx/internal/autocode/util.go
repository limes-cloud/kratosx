package autocode

import (
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func toUpperCamelCase(s string) string {
	s = strings.ReplaceAll(s, "_", " ")
	s = cases.Title(language.Und, cases.NoLower).String(s)
	return strings.ReplaceAll(s, " ", "")
}

func toLowerCamelCase(s string) string {
	s = strings.ReplaceAll(s, "_", " ")
	s = cases.Title(language.Und, cases.NoLower).String(s)
	s = strings.ReplaceAll(s, " ", "")
	prefix := strings.ToLower(string(s[0]))
	return prefix + s[1:]
}

func toLowerCase(s string) string {
	return strings.ToLower(toUpperCamelCase(s))
}

func pluralize(word string) string {
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

func uniqueStrings(input []string) []string {
	uniqueMap := make(map[string]struct{}) // 使用空结构体作为值类型，因为空结构体不占用内存。
	var result []string

	for _, str := range input {
		if _, exists := uniqueMap[str]; !exists {
			uniqueMap[str] = struct{}{}
			result = append(result, str)
		}
	}

	return result
}

func isExistFolder(folderPath string) bool {
	fileInfo, err := os.Stat(folderPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return fileInfo.IsDir()
}

func autoMkDir(fp string) {
	if strings.LastIndex(fp, "/") < strings.LastIndex(fp, ".") {
		fp = fp[:strings.LastIndex(fp, "/")]
	}

	if isExistFolder(fp) {
		return
	}
	_ = os.MkdirAll(fp, 0777)
}
