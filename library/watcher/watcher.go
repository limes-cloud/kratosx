package watcher

//
//import (
//	"github.com/go-kratos/kratos/v2/config"
//	"github.com/go-kratos/kratos/v2/log"
//)
//
//type Watcher struct {
//	watch func(string, config.Observer) error
//}
//
//func New(watch func(string, config.Observer) error) *Watcher {
//	return &Watcher{
//		watch: watch,
//	}
//}
//
//func (w *Watcher) Add(key string, dst any, handler func() error) {
//	if w == nil {
//		return
//	}
//
//	if err := w.watch(key, func(s string, value config.Value) {
//		if err := value.Scan(&dst); err != nil {
//			log.Errorf("scan %v value to dst error:%v", key, err.Error())
//			return
//		}
//		if err := handler(); err != nil {
//			log.Errorf("handle %v field value change error:%v", err.Error())
//		}
//	}); err != nil {
//		log.Errorf("watch %v field error:%v", key, err.Error())
//	}
//}
