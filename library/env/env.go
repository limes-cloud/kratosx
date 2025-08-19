package env

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
)

const (
	runMode    = "APP_RUN_MODE"
	appName    = "APP_NAME"
	appVersion = "APP_VERSION"
	appRoot    = "APP_ROOT"
)

type Env interface {
	// Get 获取应用名称
	Get(name string) string

	// RunMode 获取运行模式
	RunMode() string

	// SetRunUnitTest 标记为单元测试执行
	SetRunUnitTest()

	// RunUnitTest 是否执行单元测试
	RunUnitTest() bool

	// SetRunMode 设置运行模式
	SetRunMode(val string)

	// AppName 获取应用名称
	AppName() string

	// SetAppName 设置应用名称
	SetAppName(name string)

	// AppVersion 获取应用版本
	AppVersion() string

	// SetAppVersion 设置应用版本
	SetAppVersion(version string)

	// RootDir 获取项目根目录
	RootDir() string
}

type env struct {
	runMode    string
	appName    string
	appVersion string
	appRoot    string
	runUnit    bool
	set        sync.Map
}

var (
	ins *env

	once sync.Once
)

func Instance() Env {
	return ins
}

// Load 加载项目环境
func Load() {
	once.Do(func() {
		ins = &env{}
		ins.loadEnv()
	})
}

// loadEnv 加载环境变量
func (env *env) loadEnv() {
	// 获取当前目录
	path, _ := os.Getwd()
	if path == "" {
		return
	}

	// 设置项目根目录
	env.appRoot = path

	// 存在环境文件则直接退出查找
	if _, err := os.Stat(filepath.Join(path, ".env")); err != nil {
		return
	}

	// 加载环境变量
	if err := godotenv.Load(filepath.Join(path, ".env")); err != nil {
		log.Printf("load env error %s", err.Error())
	}
}

// Get 获取指定的env信息
func (env *env) Get(name string) string {
	value, ok := env.set.Load(name)
	if !ok {
		val := os.Getenv(name)
		env.set.Store(name, val)
		return val
	}
	v, _ := value.(string)
	return v
}

// RunMode 运行模式
func (env *env) RunMode() string {
	if env.runMode == "" {
		env.runMode = os.Getenv(runMode)
	}
	return env.runMode
}

// SetRunMode 设置运行模式
func (env *env) SetRunMode(val string) {
	env.runMode = val
	_ = os.Setenv(runMode, val)
}

// AppName 获取应用名称
func (env *env) AppName() string {
	if env.appName == "" {
		env.appName = os.Getenv(appName)
	}
	return env.appName
}

// SetAppName 设置应用名称
func (env *env) SetAppName(val string) {
	env.appName = val
	_ = os.Setenv(appName, val)
}

// AppVersion 获取应用版本
func (env *env) AppVersion() string {
	if env.appVersion == "" {
		env.appVersion = os.Getenv(appVersion)
	}
	return os.Getenv(appVersion)
}

// SetAppVersion 设置应用版本
func (env *env) SetAppVersion(val string) {
	env.appVersion = val
	_ = os.Setenv(appVersion, val)
}

// RootDir 获取项目根目录
func (env *env) RootDir() string {
	return env.appRoot
}

// SetRunUnitTest 设置为单元测试运行
func (env *env) SetRunUnitTest() {
	env.runUnit = true
}

// RunUnitTest 是否为单元测试运行
func (env *env) RunUnitTest() bool {
	return env.runUnit
}
