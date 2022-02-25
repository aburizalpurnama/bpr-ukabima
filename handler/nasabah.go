package handler

import (
	"bprukabima/nasabah"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Nasabah struct {
	nasabahService nasabah.NasabahService
}

func NewNasabahHandler(service nasabah.NasabahService) *Nasabah {
	return &Nasabah{service}
}

func (h *Nasabah) RegisterNasabah(c *gin.Context) {
	input := nasabah.InputRegisterNasabah{}

	err := c.ShouldBindJSON(&input)

	if err != nil {
		log.Fatal(err.Error())
		c.JSON(http.StatusBadRequest, input)
	}

	savedNasabah, registErr := h.nasabahService.RegisterNasabah(input)

	if err != nil {
		log.Fatal(registErr.Error)
		c.JSON(http.StatusBadRequest, input)
	}

	log.Println(savedNasabah)

	c.JSON(http.StatusOK, savedNasabah)
}
