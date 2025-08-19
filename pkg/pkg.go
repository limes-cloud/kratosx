package pkg

import (
	"path/filepath"
	"strings"
)

// AppendFileSuffix 追加文件后缀
func AppendFileSuffix(path string, suffix string) string {
	newPath := filepath.Join(
		filepath.Dir(path),
		strings.TrimSuffix(
			filepath.Base(path),
			filepath.Ext(path),
		)+suffix+filepath.Ext(path),
	)
	return newPath
}
