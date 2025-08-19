package config

import (
	"path/filepath"

	kconfig "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

type LoadRequest struct {
	Type string
	Path string
}

func Load(req *LoadRequest) kconfig.Source {
	switch req.Type {
	case "file":
		return file.NewSource(filepath.Join(req.Path))
		// case "configure":
		//	return New()
		// case "consul":
		//	return NewConsul()
	}
	return nil
}
