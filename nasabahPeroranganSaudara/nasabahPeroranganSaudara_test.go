package nasabahPeroranganSaudara

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

	nSaudara := NasabahPeroranganSaudara{}

	nSaudara.DEntry = time.Now()
	nSaudara.IdKecamatan = 3312150
	nSaudara.IdKelurahan = 3312150008
	nSaudara.IdKodepos = 39615
	nSaudara.IdKotaKab = 3312
	nSaudara.IdNasabah = 183887
	nSaudara.IdPropinsi = 33
	nSaudara.NamaSaudara = "Ujang"
	nSaudara.NoTelp = "082299918291"

	repository := NewRepository(db)
	n, repositoryErr := repository.Save(nSaudara)

	if repositoryErr != nil {
		log.Fatal(repositoryErr.Error())
	}

	fmt.Println(n)

}
