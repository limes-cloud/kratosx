package gormtranserror

import "errors"

type GormError struct {
	oriError error
	e        string
}

func NewError(ori error, text string) error {
	return &GormError{
		oriError: ori,
		e:        text,
	}
}

func Is(src, dst error) bool {
	_src, ok := src.(*GormError)
	if ok {
		return errors.Is(_src.oriError, dst)
	}

	_dst, ok := dst.(*GormError)
	if ok {
		return errors.Is(_dst.oriError, src)
	}

	return errors.Is(src, src)
}

func (ge *GormError) Error() string {
	return ge.e
}
