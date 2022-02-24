package nasabahPeroranganPersonal

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

	nPersonal := NasabahPeroranganPersonal{}

	nPersonal.IdAgama = 94
	nPersonal.IdKotaKabupatenLahir = 3312
	nPersonal.IdKotaKabupatenLahirPasangan = 3312
	nPersonal.IdPendidikan = 390
	nPersonal.IdPropinsiLahir = 33
	nPersonal.IdPropinsiLahirPasangan = 33
	nPersonal.DEntry = time.Now()

	nPersonal.IdNasabahPerorangan = 183887

	repository := NewRepository(db)
	n, repositoryErr := repository.Save(nPersonal)

	if repositoryErr != nil {
		log.Fatal(repositoryErr.Error())
	}

	fmt.Println(n)

}
