package main

import (
	"bprukabima/nasabah"
	"bprukabima/utils"
	"log"
	"time"
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

	// create instance of repository
	nasabahRepository := nasabah.NewRepository(db)

	nasabah := nasabah.NewNasabah()
	nasabah.IdNasabahPerorangan = 92002
	nasabah.DEntry = time.Now()
	nasabah.IsBlacklist = false
	nasabah.NamaIdentitas = "Test"
	nasabah.IdGolDeb = 557
	nasabah.IdPropinsiIdentitas = 33
	nasabah.IdKotaKabupatenIdentitas = 3312
	nasabah.IdKecamatanIdentitas = 3312150
	nasabah.IdKelurahanIdentitas = 3312150008
	nasabah.IdKodePosIdentitas = 39615
	nasabah.IdPropinsiDomisili = 33
	nasabah.IdKotaKabupatenDomisili = 3312
	nasabah.IdKecamatanDomisili = 3312150
	nasabah.IdKelurahanDomisili = 3312150008
	nasabah.IdKodePosDomisili = 39615
	nasabah.IdBpr = 2
	nasabah.BprCabang = 1

	nasabahRepository.Save(*nasabah)
}
