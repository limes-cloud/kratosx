package tasker

import (
	"fmt"
	"sync"
	"time"

	"github.com/limes-cloud/kratosx/library/logger"
)

type Tasker interface {
	// BeforeStart 注册启动前的处理函数
	BeforeStart(name string, f func())

	// AfterStart 注册启动后的处理函数
	AfterStart(name string, f func())

	// BeforeStop 注册停止前的处理函数
	BeforeStop(name string, f func())

	// AfterStop 注册停止后的处理函数
	AfterStop(name string, f func())

	// Remove 移除指定name的回调
	Remove(name string)

	// WaitBeforeStart 等待启动前的函数完成
	WaitBeforeStart()

	// WaitAfterStart 等待启动后的函数完成
	WaitAfterStart()

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
	mu          sync.Mutex
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
	s.mu.Lock()
	defer s.mu.Unlock()
	s.beforeStart = append(s.beforeStart, item{
		name: name,
		f:    f,
	})
}

func (s *stop) AfterStart(name string, f func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.afterStart = append(s.afterStart, item{
		name: name,
		f:    f,
	})
}

func (s *stop) BeforeStop(name string, f func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.beforeStop = append(s.beforeStop, item{
		name: name,
		f:    f,
	})
}

func (s *stop) AfterStop(name string, f func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.afterStop = append(s.afterStop, item{
		name: name,
		f:    f,
	})
}

func (s *stop) Remove(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.beforeStart = removeItem(s.beforeStart, name)
	s.afterStart = removeItem(s.afterStart, name)
	s.beforeStop = removeItem(s.beforeStop, name)
	s.afterStop = removeItem(s.afterStop, name)
}

func removeItem(items []item, name string) []item {
	result := items[:0]
	for _, it := range items {
		if it.name != name {
			result = append(result, it)
		}
	}
	return result
}

func (s *stop) WaitBeforeStart() {
	s.mu.Lock()
	items := make([]item, len(s.beforeStart))
	copy(items, s.beforeStart)
	s.beforeStart = s.beforeStart[:0]
	s.mu.Unlock()

	for _, item := range items {
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
	s.mu.Lock()
	items := make([]item, len(s.afterStart))
	copy(items, s.afterStart)
	s.afterStart = s.afterStart[:0]
	s.mu.Unlock()

	for _, item := range items {
		t := time.Now().UnixMilli()
		logger.Instance().Info("wait afterStart start", logger.F("name", item.name))
		item.f()
		item.f = nil
		logger.Instance().Info("wait afterStart finish",
			logger.F("name", item.name),
			logger.F("time", fmt.Sprintf("%dms", time.Now().UnixMilli()-t)),
		)
	}
}

func (s *stop) WaitBeforeStop() {
	s.mu.Lock()
	items := make([]item, len(s.beforeStop))
	copy(items, s.beforeStop)
	s.beforeStop = s.beforeStop[:0]
	s.mu.Unlock()

	for _, item := range items {
		t := time.Now().UnixMilli()
		logger.Instance().Info("wait beforeStop start", logger.F("name", item.name))
		item.f()
		item.f = nil
		logger.Instance().Info("wait beforeStop finish",
			logger.F("name", item.name),
			logger.F("time", fmt.Sprintf("%dms", time.Now().UnixMilli()-t)),
		)
	}
}

func (s *stop) WaitAfterStop() {
	s.mu.Lock()
	items := make([]item, len(s.afterStop))
	copy(items, s.afterStop)
	s.afterStop = s.afterStop[:0]
	s.mu.Unlock()

	for _, item := range items {
		t := time.Now().UnixMilli()
		logger.Instance().Info("wait afterStop start", logger.F("name", item.name))
		item.f()
		item.f = nil
		logger.Instance().Info("wait afterStop finish",
			logger.F("name", item.name),
			logger.F("time", fmt.Sprintf("%dms", time.Now().UnixMilli()-t)),
		)
	}
}
