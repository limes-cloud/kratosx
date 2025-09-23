package tasker

import (
	"fmt"
	"github.com/limes-cloud/kratosx/library/logger"
	"time"
)

type Tasker interface {
	// BeforeStart 注册启动前的处理函数
	BeforeStart(name string, f func())

	// AfterStart 注册启动后的处理函数
	AfterStart(name string, f func())

	// WaitBeforeStart 等待启动前的函数完成
	WaitBeforeStart()

	// WaitAfterStart 等待启动后的函数完成
	WaitAfterStart()

	// BeforeStop 注册停止前的处理函数
	BeforeStop(name string, f func())

	// AfterStop 注册停止后的处理函数
	AfterStop(name string, f func())

	// WaitBeforeStop 等待停止前的函数完成
	WaitBeforeStop()

	// WaitAfterStop 等待停止后的函数完成
	WaitAfterStop()
}

type item struct {
	name string
	f    func()
}

type stop struct {
	afterStart  []item
	beforeStart []item
	afterStop   []item
	beforeStop  []item
}

var ins = &stop{
	beforeStart: make([]item, 0),
	afterStart:  make([]item, 0),
	beforeStop:  make([]item, 0),
	afterStop:   make([]item, 0),
}

func Instance() Tasker {
	return ins
}

func (s *stop) BeforeStart(name string, f func()) {
	s.beforeStart = append(s.beforeStart, item{
		name: name,
		f:    f,
	})
}

func (s *stop) AfterStart(name string, f func()) {
	s.afterStart = append(s.beforeStart, item{
		name: name,
		f:    f,
	})
}

func (s *stop) BeforeStop(name string, f func()) {
	s.beforeStop = append(s.beforeStop, item{
		name: name,
		f:    f,
	})
}

func (s *stop) AfterStop(name string, f func()) {
	s.afterStop = append(s.afterStop, item{
		name: name,
		f:    f,
	})
}

func (s *stop) WaitBeforeStart() {
	for _, item := range s.beforeStart {
		t := time.Now().UnixMilli()
		logger.Instance().Info("wait beforeStart start", logger.F("name", item.name))
		item.f()
		item.f = nil
		logger.Instance().Info("wait beforeStart finish",
			logger.F("name", item.name),
			logger.F("time", fmt.Sprintf("%dms", time.Now().UnixMilli()-t)),
		)
	}
}

func (s *stop) WaitAfterStart() {
	for _, item := range s.afterStart {
		t := time.Now().UnixMilli()
		logger.Instance().Info("wait afterStart start", logger.F("name", item.name))
		item.f()
		item.f = nil
		logger.Instance().Info("wait beforeStart finish",
			logger.F("name", item.name),
			logger.F("time", fmt.Sprintf("%dms", time.Now().UnixMilli()-t)),
		)
	}
}

func (s *stop) WaitBeforeStop() {
	for _, item := range s.beforeStop {
		t := time.Now().UnixMilli()
		logger.Instance().Info("wait beforeStop start", logger.F("name", item.name))
		item.f()
		item.f = nil
		logger.Instance().Info("wait beforeStart finish",
			logger.F("name", item.name),
			logger.F("time", fmt.Sprintf("%dms", time.Now().UnixMilli()-t)),
		)
	}
}

func (s *stop) WaitAfterStop() {
	for _, item := range s.beforeStop {
		t := time.Now().UnixMilli()
		logger.Instance().Info("wait afterStop start", logger.F("name", item.name))
		item.f()
		item.f = nil
		logger.Instance().Info("wait beforeStart finish",
			logger.F("name", item.name),
			logger.F("time", fmt.Sprintf("%dms", time.Now().UnixMilli()-t)),
		)
	}
}
