package nasabahPeroranganInfo

import "time"

type NasabahPeroranganInfo struct {
	ID int `gorm:"column:id_nasabah_perorangan_infoalainnya"`

	IdNasabah          int
	NamaAliasNasabah   string
	StatusHubungan     string
	IdNegara           int `gorm:"column:id_negara_lainnya"`
	IdHubunganPelapor  int
	IdReferensiNasabah int

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
func (NasabahPeroranganInfo) TableName() string {
	return "m_nasabah_perorangan_informasi_lainnya"
}
