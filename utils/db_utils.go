package utils

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GetDb() (*gorm.DB, error) {
	dsn := "host=localhost user=rizal password=qwerty123 dbname=bpr port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{

		// add configuration for singular table name
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db, err
}
