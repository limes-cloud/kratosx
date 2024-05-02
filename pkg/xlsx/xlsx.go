package xlsx

import "github.com/xuri/excelize/v2"

type Xlsx interface {
	Reader() (Reader, error)
	Writer() Writer
}

type xlsx struct {
	path string
}

func New(path string) Xlsx {
	return &xlsx{path: path}
}

func (x *xlsx) Reader() (Reader, error) {
	xf, err := excelize.OpenFile(x.path)
	if err != nil {
		return nil, err
	}
	return &reader{
		xlsx: xf,
	}, nil
}

func (x *xlsx) Writer() Writer {
	return &writer{path: x.path, xlsx: excelize.NewFile(), sheet: "Sheet1"}
}
