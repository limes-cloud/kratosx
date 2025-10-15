package base

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
)

// ModulePath returns go module path.
func ModulePath(filename string) (string, error) {
	modBytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return modfile.ModulePath(modBytes), nil
}

// ModuleVersion returns module version.
func ModuleVersion(proDir string, path string) (string, error) {
	stdout := &bytes.Buffer{}
	fd := exec.Command("go", "mod", "graph")
	fd.Stdout = stdout
	fd.Stderr = stdout
	fd.Dir = proDir
	if err := fd.Run(); err != nil {
		return "", err
	}
	rd := bufio.NewReader(stdout)
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			return "", err
		}
		str := string(line)
		i := strings.Index(str, "@")
		if strings.Contains(str, path+"@") && i != -1 {
			return path + str[i:], nil
		}
	}
}

// KratosMod returns kratos mod.
func KratosMod(proDir string) string {
	// go 1.15+ read from env GOMODCACHE
	cacheOut, _ := exec.Command("go", "env", "GOMODCACHE").Output()
	cachePath := strings.Trim(string(cacheOut), "\n")
	pathOut, _ := exec.Command("go", "env", "GOPATH").Output()
	gopath := strings.Trim(string(pathOut), "\n")
	if cachePath == "" {
		cachePath = filepath.Join(gopath, "pkg", "mod")
	}
	if path, err := ModuleVersion(proDir, "github.com/go-kratos/kratos/v2"); err == nil {
		// $GOPATH/pkg/mod/github.com/limes-cloud/kratos@v2
		return filepath.Join(cachePath, path)
	}
	// $GOPATH/src/github.com/limes-cloud/kratos
	return filepath.Join(gopath, "src", "github.com", "go-kratos", "kratos")
}

// KratosxMod returns kratos mod.
func KratosxMod(proDir string) string {
	// go 1.15+ read from env GOMODCACHE
	cacheOut, _ := exec.Command("go", "env", "GOMODCACHE").Output()
	cachePath := strings.Trim(string(cacheOut), "\n")
	pathOut, _ := exec.Command("go", "env", "GOPATH").Output()
	gopath := strings.Trim(string(pathOut), "\n")
	if cachePath == "" {
		cachePath = filepath.Join(gopath, "pkg", "mod")
	}

	if path, err := ModuleVersion(proDir, "github.com/limes-cloud/kratosx"); err == nil {
		// $GOPATH/pkg/mod/github.com/limes-cloud/kratos@v2
		return filepath.Join(cachePath, path)
	}
	// $GOPATH/src/github.com/limes-cloud/kratos
	return filepath.Join(gopath, "src", "github.com", "limes-cloud", "kratosx")
}

// KratosxCliMod returns kratos mod.
func KratosxCliMod() string {
	path := os.Getenv("AUTOCODE_TEMP_PATH")
	if path != "" {
		return path
	}
	// go 1.15+ read from env GOMODCACHE
	cacheOut, _ := exec.Command("go", "env", "GOMODCACHE").Output()
	cachePath := strings.Trim(string(cacheOut), "\n")
	pathOut, _ := exec.Command("go", "env", "GOPATH").Output()
	gopath := strings.Trim(string(pathOut), "\n")
	if cachePath == "" {
		cachePath = filepath.Join(gopath, "pkg", "mod")
	}
	cliPath := cachePath + "/github.com/limes-cloud/kratosx/cmd"

	files, err := os.ReadDir(cliPath)
	if err != nil {
		return filepath.Join(gopath, "src", "github.com", "limes-cloud", "kratosx", "cmd", "kratosx")
	}

	var lastKratosxDir string
	for _, file := range files {
		if file.IsDir() {
			dirName := file.Name()
			if strings.HasPrefix(dirName, "kratosx@") {
				lastKratosxDir = dirName
			}
		}
	}

	return cliPath + "/" + lastKratosxDir
}

// ModName 获取mod名称
func ModName(p string) string {
	path, _ := os.Getwd()
	if path == "" {
		return ""
	}

	// 绝对路径，移除前缀path
	if filepath.IsAbs(p) {
		p = strings.TrimPrefix(p, path)
	}

	path = filepath.Join(path, p)

	// 判断当前是否存在环境文件
	for {
		// 出现go.mod 认为在根目录
		if _, err := os.Stat(filepath.Join(path, "go.mod")); err == nil {
			break
		}

		if path == "" || path == "/" {
			break
		}

		// 往上移动一个目录
		path = filepath.Dir(path)
	}
	mod, err := ModulePath(path + "/go.mod")
	if err != nil {
		return ""
	}
	return mod
}
