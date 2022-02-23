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
	nasabah.DEntry = time.Now()
	nasabah.IsBlacklist = false
	nasabah.NamaIdentitas = "Test"

	nasabahRepository.Save(*nasabah)
}
