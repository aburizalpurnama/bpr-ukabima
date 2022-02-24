package nasabahPeroranganPersonal

import (
	"bprukabima/nasabah"
	"time"
)

type NasabahPeroranganPersonal struct {
	ID                           int `gorm:"column:id_nasabah_perorangan_personal"`
	NasabahId                    int
	Nasabah                      nasabah.Nasabah
	JenisKelamin                 int
	IdPropinsiLahir              int
	IdKotaKabupatenLahir         int
	TglLahir                     time.Time
	NamaIbuKandung               string
	IdAgama                      int
	IdPendidikan                 int
	StatusPerkawinan             int
	NoIdentitasPasangan          string
	TglLahirPasangan             time.Time
	IdPropinsiLahirPasangan      int
	IdKotaKabupatenLahirPasangan int
	NamaPasangan                 string
	PerjanjianPisahHarta         int
	JumlahAnak                   int
}
