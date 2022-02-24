package nasabahPeroranganAlamatUsaha

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

	nAlamatUsaha := NasabahPeroranganAlamatUsaha{}

	nAlamatUsaha.DEntry = time.Now()
	nAlamatUsaha.IdKecamatan = 3312150
	nAlamatUsaha.IdKelurahan = 3312150008
	nAlamatUsaha.IdKodePos = 39615
	nAlamatUsaha.IdKotaKab = 3312
	nAlamatUsaha.IdNasabah = 183887
	nAlamatUsaha.IdPropinsi = 33
	nAlamatUsaha.NoTelp = "082299918291" // field class = int, column db = varchar
	nAlamatUsaha.NamaUsaha = "Jualan Organ"

	repository := NewRepository(db)
	n, repositoryErr := repository.Save(nAlamatUsaha)

	if repositoryErr != nil {
		log.Fatal(repositoryErr.Error())
	}

	fmt.Println(n)

}
