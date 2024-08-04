package proto

import "github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen"

type Proto struct {
	*Error
	*Message
	*Service
}

type Builder interface {
	GenError() (string, error)
	GenMessage() (string, error)
	GenService() (string, error)
}

func NewBuilder(builder *gen.Builder) Builder {
	return &Proto{
		Error:   NewErrorBuilder(builder),
		Message: NewMessageBuilder(builder),
		Service: NewServiceBuilder(builder),
	}
}
