package nasabahPeroranganPekerjaan

import "time"

type NasabahPeroranganPekerjaan struct {
	ID        int `gorm:"column:id_nasabah_perorangan_pekerjaan"`
	IdNasabah int
	// Nasabah                      nasabah.Nasabah `gorm:"foreignKey:IdNasabahPerorangan"`
	IdPekerjaan              int
	NamaPekerjaan            string
	TempatPekerjaan          string
	AlamatPekerjaan          string
	IdBidangUsaha            int
	RangePenghasilanPerbulan int // foreign
	JumlahTanggungan         int
	NasabahBeresiko          string
	Pep                      bool `gorm:"column:id_nasabah_perorangan_pekerjaan"`
	IdGroupAnggota           int

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
func (NasabahPeroranganPekerjaan) TableName() string {
	return "m_nasabah_perorangan_pekerjaan"
}
