package gormtranserror

import (
	"errors"
	"gorm.io/gorm"
)

type GormError struct {
	oriError error
	e        string
}

func NewError(db *gorm.DB, ori error, text string) error {
	if errTranslator, ok := db.Dialector.(gorm.ErrorTranslator); ok {
		return &GormError{
			oriError: errTranslator.Translate(ori),
			e:        text,
		}
	}
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

	return errors.Is(src, dst)
}

func (ge *GormError) Error() string {
	return ge.e
}
