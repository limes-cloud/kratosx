package gocode

import "github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen"

type GoCode struct {
	*Types
	*Entity
	*Dbs
	*Repo
	*Service
	*App
}

type Builder interface {
	GenTypes() (string, error)
	GenEntity() (string, error)
	GenDbs() (string, error)
	GenRepo() (string, error)
	GenService() (string, error)
	GenApp() (string, error)
	GenAppEntry() (string, error)
}

func NewBuilder(builder *gen.Builder) Builder {
	return &GoCode{
		Types:   NewTypesBuilder(builder),
		Entity:  NewEntityBuilder(builder),
		Dbs:     NewDbsBuilder(builder),
		Repo:    NewRepoBuilder(builder),
		Service: NewServiceBuilder(builder),
		App:     NewAppBuilder(builder),
	}
}
