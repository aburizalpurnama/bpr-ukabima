package nasabahPeroranganPekerjaan

import (
	"bprukabima/utils"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestSaveFromRepository(t *testing.T) {
	db, err := utils.GetDb()

	if err != nil {
		log.Fatal(err.Error())
	}

	nPekerjaan := NasabahPeroranganPekerjaan{}

	nPekerjaan.IdBidangUsaha = 13
	nPekerjaan.IdNasabah = 183887
	nPekerjaan.IdPekerjaan = 216
	nPekerjaan.RangePenghasilanPerbulan = 4
	nPekerjaan.NamaPekerjaan = "WIRASWASTA"
	nPekerjaan.IdGroupAnggota = 4233
	nPekerjaan.DEntry = time.Now()

	// nPersonal.IdNasabahPerorangan = 183887

	repository := NewRepository(db)
	n, repositoryErr := repository.Save(nPekerjaan)

	if repositoryErr != nil {
		log.Fatal(repositoryErr.Error())
	}

	fmt.Println(n)

}
