package main

import (
	"bprukabima/handler"
	"bprukabima/nasabah"
	"bprukabima/nasabahPeroranganAlamatUsaha"
	"bprukabima/nasabahPeroranganInfo"
	"bprukabima/nasabahPeroranganPekerjaan"
	"bprukabima/nasabahPeroranganPersonal"
	"bprukabima/nasabahPeroranganSaudara"
	"bprukabima/utils"
	"log"

	"github.com/gin-gonic/gin"
)

// func main() {
// 	router := gin.Default()

// 	router.GET("/nasabah", UserHandler)

// 	router.Run()
// }

// func UserHandler(c *gin.Context) {
// 	db, err := utils.GetDb()

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	var nasabah []nasabah.Nasabah

// 	db.Find(&nasabah)

// 	c.JSON(http.StatusOK, nasabah)
// }

func main() {

	// get db connection
	db, err := utils.GetDb()

	if err != nil {
		log.Fatal(err.Error())
	}

	// Create repositories
	nGeneralRepo := nasabah.NewRepository(db)
	nAlamatUsahaRepo := nasabahPeroranganAlamatUsaha.NewRepository(db)
	nInfoRepo := nasabahPeroranganInfo.NewRepository(db)
	nPersonalRepo := nasabahPeroranganPersonal.NewRepository(db)
	nPekerjaanRepo := nasabahPeroranganPekerjaan.NewRepository(db)
	nSaudaraRepo := nasabahPeroranganSaudara.NewRepository(db)

	// Create service
	nService := nasabah.NewService(nGeneralRepo, nAlamatUsahaRepo, nInfoRepo, nPersonalRepo, nPekerjaanRepo, nSaudaraRepo)

	// Create handler
	nHandler := handler.NewNasabahHandler(nService)

	router := gin.Default()

	// API Versioning
	api := router.Group("/api/v1")

	api.POST("/create/permohonanNasabahBimaApps", nHandler.RegisterNasabah)

	router.Run()

}
