package response

import (
	"fmt"
	stdhttp "net/http"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// httpResponse 响应结构体
// nolint
type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// HTTPEncoder http 请求response
func HTTPEncoder() http.ServerOption {
	return http.ResponseEncoder(encoderResponse())
}

// ErrorEncoder err 请求response
func ErrorEncoder() http.ServerOption {
	return http.ErrorEncoder(encoderError())
}

// EncoderResponse  请求响应封装
func encoderResponse() http.EncodeResponseFunc {
	return func(w stdhttp.ResponseWriter, request *stdhttp.Request, i any) error {
		if i == nil {
			return nil
		}
		tpl := `{"code":200,"message":"success!","data":%s}`
		codec := encoding.GetCodec("json")
		data, err := codec.Marshal(i)
		if err != nil {
			return err
		}
		reply := fmt.Sprintf(tpl, string(data))
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write([]byte(reply))
		if err != nil {
			return err
		}
		return nil
	}
}

// EncoderError 错误响应封装
// nolint
func encoderError() http.EncodeErrorFunc {
	return func(w stdhttp.ResponseWriter, r *stdhttp.Request, err error) {
		if err == nil {
			return
		}
		se := &response{}
		kerr, ok := err.(*errors.Error)
		if !ok {
			se = &response{Code: stdhttp.StatusBadRequest, Message: err.Error()}
		} else {
			se = &response{
				Code:    int(kerr.Code),
				Message: kerr.Message,
			}
		}

		codec, _ := http.CodecForRequest(r, "Accept")
		body, err := codec.Marshal(se)
		if err != nil {
			w.WriteHeader(stdhttp.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(stdhttp.StatusOK)
		_, _ = w.Write(body)
	}
}
