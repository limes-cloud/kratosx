package core

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	Gorm *gorm.DB
)

func ConnectDatabase(dial gorm.Dialector) error {
	db, err := gorm.Open(dial, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}
	Gorm = db
	return nil
}
