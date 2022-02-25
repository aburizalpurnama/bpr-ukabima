package nasabahPeroranganPersonal

import (
	"time"
)

type NasabahPeroranganPersonal struct {
	ID                  int `gorm:"column:id_nasabah_perorangan_personal"`
	IdNasabahPerorangan int
	// Nasabah                      nasabah.Nasabah `gorm:"foreignKey:IdNasabahPerorangan"`
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

	// Base entity
	IdEntry  string
	IdUpdate string
	DEntry   time.Time
	DUpdate  time.Time
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (NasabahPeroranganPersonal) TableName() string {
	return "m_nasabah_perorangan_personal"
}
