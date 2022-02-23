package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestConnectionToDB(t *testing.T) {
	dsn := "host=localhost user=rizal password=qwerty123 dbname=bpr port=5432 sslmode=disable TimeZone=UTC"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	assert.Nil(t, err)
}
