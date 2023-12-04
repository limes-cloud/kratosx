package config

import (
	kratosConfig "github.com/go-kratos/kratos/v2/config"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/protobuf/proto"
)

type Value interface {
	kratosConfig.Value
}

type value struct {
	kratosConfig.Value
}

func (v *value) Scan(dst any) error {
	if _, ok := dst.(proto.Message); ok {
		return v.Value.Scan(&dst)
	}

	// 序列化json
	var res any
	if err := v.Value.Scan(&res); err != nil {
		return err
	}

	dc := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           dst,
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
		),
	}
	decoder, err := mapstructure.NewDecoder(dc)
	if err != nil {
		return err
	}
	return decoder.Decode(res)
}
