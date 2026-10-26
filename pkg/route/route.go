// Package route 从 proto 描述符构建 gRPC operation → HTTP(method,path) 的运行时映射。
// 用途：抹平 HTTP 与 gRPC 两种寻址——gRPC 请求按 operation 反查其 HTTP method+path，
// 使白名单 / 鉴权只需按 HTTP 维度登记一次即可同时覆盖两种传输。
package route

import (
	"sync"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// HTTPRoute 一个 RPC 对应的 HTTP 入口（取自 proto 的 google.api.http 注解）。
type HTTPRoute struct {
	Method string // GET / POST / PUT / DELETE / PATCH
	Path   string // 如 /v1/resource/store
}

var (
	index map[string]HTTPRoute // key = gRPC operation "/pkg.Service/Method"
	once  sync.Once
)

// Lookup 按 gRPC operation 反查对应的 HTTP method+path（operation 即 transport.Operation()）。
func Lookup(operation string) (HTTPRoute, bool) {
	once.Do(build)
	r, ok := index[operation]
	return r, ok
}

// build 遍历已注册的 proto 描述符，建立 operation → HTTPRoute 映射。
func build() {
	index = make(map[string]HTTPRoute)
	protoregistry.GlobalFiles.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		svcs := fd.Services()
		for i := 0; i < svcs.Len(); i++ {
			svc := svcs.Get(i)
			ms := svc.Methods()
			for j := 0; j < ms.Len(); j++ {
				m := ms.Get(j)
				r, ok := httpRule(m)
				if !ok {
					continue
				}
				op := "/" + string(svc.FullName()) + "/" + string(m.Name())
				index[op] = r
			}
		}
		return true
	})
}

// httpRule 从方法的 google.api.http 注解提取 HTTP method+path。
func httpRule(m protoreflect.MethodDescriptor) (HTTPRoute, bool) {
	opts := m.Options()
	if opts == nil {
		return HTTPRoute{}, false
	}
	rule, ok := proto.GetExtension(opts, annotations.E_Http).(*annotations.HttpRule)
	if !ok || rule == nil {
		return HTTPRoute{}, false
	}
	switch p := rule.Pattern.(type) {
	case *annotations.HttpRule_Get:
		return HTTPRoute{Method: "GET", Path: p.Get}, true
	case *annotations.HttpRule_Post:
		return HTTPRoute{Method: "POST", Path: p.Post}, true
	case *annotations.HttpRule_Put:
		return HTTPRoute{Method: "PUT", Path: p.Put}, true
	case *annotations.HttpRule_Delete:
		return HTTPRoute{Method: "DELETE", Path: p.Delete}, true
	case *annotations.HttpRule_Patch:
		return HTTPRoute{Method: "PATCH", Path: p.Patch}, true
	case *annotations.HttpRule_Custom:
		return HTTPRoute{Method: p.Custom.Kind, Path: p.Custom.Path}, true
	}
	return HTTPRoute{}, false
}
