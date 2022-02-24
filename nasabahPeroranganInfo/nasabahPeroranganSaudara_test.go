package nasabahPeroranganInfo

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

	nInfo := NasabahPeroranganInfo{}

	nInfo.DEntry = time.Now()
	nInfo.IdHubunganPelapor = 19
	nInfo.IdNasabah = 183887
	nInfo.IdNegara = 62
	nInfo.IdReferensiNasabah = 31655
	nInfo.NamaAliasNasabah = "Test"

	repository := NewRepository(db)
	n, repositoryErr := repository.Save(nInfo)

	if repositoryErr != nil {
		log.Fatal(repositoryErr.Error())
	}

	fmt.Println(n)

}
