package env

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// RootDir 获取项目根目录
func RootDir() string {
	path, _ := os.Getwd()
	if path == "" {
		return ""
	}

	// 判断当前是否存在环境文件
	for {
		// 出现go.mod 认为在根目录
		if _, err := os.Stat(filepath.Join(path, "go.mod")); err == nil {
			return path
		}

		if path == "" || path == "/" {
			return ""
		}

		// 往上移动一个目录
		path = filepath.Dir(path)
	}
}

// Load 加载项目环境
func Load() {
	path, _ := os.Getwd()
	if path == "" {
		return
	}

	// 判断当前是否存在环境文件
	for {
		// 存在环境文件则直接退出查找
		if _, err := os.Stat(filepath.Join(path, ".env")); err == nil {
			break
		}

		// 直到到达根目录还没找到，直接返回
		if _, err := os.Stat(filepath.Join(path, "go.mod")); err == nil {
			return
		}

		if path == "" || path == "/" {
			return
		}

		// 往上移动一个目录
		path = filepath.Dir(path)
	}

	if err := godotenv.Load(filepath.Join(path, ".env")); err != nil {
		log.Printf("load env error %s", err.Error())
	}
}
