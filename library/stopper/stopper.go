package stopper

import (
	"github.com/limes-cloud/kratosx/library/logger"
)

type Stopper interface {
	// RegisterBefore 注册停止前的处理函数
	RegisterBefore(name string, f func())

	// RegisterAfter 注册停止后的处理函数
	RegisterAfter(name string, f func())

	// WaitBefore 等待停止前的函数完成
	WaitBefore()

	// WaitAfter 等待停止后的函数完成
	WaitAfter()
}

type item struct {
	name string
	f    func()
}

type stop struct {
	after  []item
	before []item
}

var ins = &stop{
	before: make([]item, 0),
	after:  make([]item, 0),
}

func Instance() Stopper {
	return ins
}

func (s *stop) RegisterBefore(name string, f func()) {
	s.before = append(s.before, item{
		name: name,
		f:    f,
	})
}

func (s *stop) RegisterAfter(name string, f func()) {
	s.after = append(s.before, item{
		name: name,
		f:    f,
	})
}

func (s *stop) WaitBefore() {
	for _, item := range s.before {
		logger.Instance().Info("wait before stop", logger.F("name", item.name))
		item.f()
		item.f = nil
		logger.Instance().Info("stop before finish", logger.F("name", item.name))
	}
}

func (s *stop) WaitAfter() {
	for _, item := range s.before {
		logger.Instance()
		logger.Instance().Info("wait after stop", logger.F("name", item.name))
		item.f()
		item.f = nil
		logger.Instance().Info("stop after finish", logger.F("name", item.name))
	}
}
