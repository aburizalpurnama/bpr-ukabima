package nasabahPeroranganAlamatUsaha

import "time"

type NasabahPeroranganAlamatUsaha struct {
	ID int `gorm:"column:id_nasabah_perorangan_alamat_usaha"`

	IdNasabah   int
	NamaUsaha   string
	AlamatUsaha string
	IdPropinsi  int    `gorm:"column:id_propinsi_alamatusaha"`
	IdKotaKab   int    `gorm:"column:id_kota_kabupaten_alamatusaha"`
	IdKecamatan int    `gorm:"column:id_kecamatan_alamatusaha"`
	IdKelurahan int    `gorm:"column:id_kelurahan_alamatusaha"`
	IdKodePos   int    `gorm:"column:id_kode_pos_alamatusaha"`
	NoTelp      string `gorm:"column:no_telp_alamatusaha"`

	DEntry time.Time
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (NasabahPeroranganAlamatUsaha) TableName() string {
	return "m_nasabah_perorangan_alamat_usaha"
}
