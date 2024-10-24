package stop

import (
	"github.com/limes-cloud/kratosx/library/logger"
)

type Stop interface {
	// Register 注册
	Register(name string, f func())

	// Wait 等待
	Wait()
}

type stop struct {
	m map[string]func()
}

var instance *stop

func Instance() Stop {
	return instance
}

func Init() {
	instance = &stop{
		m: make(map[string]func()),
	}
}

func (s stop) Register(name string, f func()) {
	s.m[name] = f
}

func (s stop) Wait() {
	for k, f := range s.m {
		logger.Helper().Infow("msg", "wait stop", "name", k)
		f()
		logger.Helper().Infow("msg", "stop finish", "name", k)
	}
}
