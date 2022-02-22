package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestConnectionToDB(t *testing.T) {
	dsn := "host=localhost user=rizal password=qwerty123 dbname=bpr_2 port=5432 sslmode=disable TimeZone=UTC"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	assert.Nil(t, err)
}
