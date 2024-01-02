package gormtranserror

import "gorm.io/gorm"

type Option func(o *options)

type Error map[error]error

type options struct {
	es                  map[error]error
	duplicatedKeyFormat string
	addForeignKeyFormat string
	delForeignKeyFormat string
	enableLoad          bool
	db                  *gorm.DB
}

func WithError(es map[error]error) Option {
	return func(o *options) {
		o.es = es
	}
}

func WithDuplicatedKeyFormat(fm string) Option {
	return func(o *options) {
		o.duplicatedKeyFormat = fm
	}
}

func WithAddForeignKeyFormat(fm string) Option {
	return func(o *options) {
		o.addForeignKeyFormat = fm
	}
}

func WithDelForeignKeyFormat(fm string) Option {
	return func(o *options) {
		o.delForeignKeyFormat = fm
	}
}

func WithEnableLoad() Option {
	return func(o *options) {
		o.enableLoad = true
	}
}

func WithGorm(db *gorm.DB) Option {
	return func(o *options) {
		o.db = db
	}
}
