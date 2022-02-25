package nasabah

import (
	"bprukabima/nasabahPeroranganAlamatUsaha"
	"bprukabima/nasabahPeroranganInfo"
	"bprukabima/nasabahPeroranganPekerjaan"
	"bprukabima/nasabahPeroranganPersonal"
	"bprukabima/nasabahPeroranganSaudara"
	"bprukabima/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterFromServiceLayer(t *testing.T) {

	input := InputRegisterNasabah{}

	input.Alamat = "Test Bekasi"
	input.Email = "test@mail.com"
	input.IdKecamatan = 3312150
	input.IdKelurahan = 3312150008
	input.IdKodeBpr = 2
	input.IdKodeCabang = 1
	input.IdKodePos = 39615
	input.IdKotaKab = 3312
	input.IdPropinsi = 33
	input.Kelamin = "Laki-laki"
	input.Mobile = "08123456789"
	input.Nama = "Test insert from service layer"
	input.Nik = "092938917283761"

	db, dbErr := utils.GetDb()

	assert.Nil(t, dbErr)

	nGeneralRepo := NewRepository(db)
	nAlamatUsahaRepo := nasabahPeroranganAlamatUsaha.NewRepository(db)
	nInfoRepo := nasabahPeroranganInfo.NewRepository(db)
	nPersonalRepo := nasabahPeroranganPersonal.NewRepository(db)
	nPekerjaanRepo := nasabahPeroranganPekerjaan.NewRepository(db)
	nSaudaraRepo := nasabahPeroranganSaudara.NewRepository(db)

	nService := NewService(nGeneralRepo, nAlamatUsahaRepo, nInfoRepo, nPersonalRepo, nPekerjaanRepo, nSaudaraRepo)

	response, registErr := nService.RegisterNasabah(input)

	assert.Nil(t, registErr)

	fmt.Println(response)
}
