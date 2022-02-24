package main

import (
	"bprukabima/utils"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectionToDB(t *testing.T) {

	_, err := utils.GetDb()

	if err != nil {
		log.Fatal(err.Error())
	}
	assert.Nil(t, err)
}
