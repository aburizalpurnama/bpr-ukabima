package nasabahPeroranganSaudara

import "time"

type NasabahPeroranganSaudara struct {
	ID                 int `gorm:"column:id_nasabah_perorangan_saudara"`
	IdNasabah          int
	NamaSaudara        string
	AlamatSaudara      string
	NoIdentitasSaudara string
	IdPropinsi         int    `gorm:"column:id_propinsi_saudara"`
	IdKotaKab          int    `gorm:"column:id_kota_kabupaten_saudara"`
	IdKecamatan        int    `gorm:"column:id_kecamatan_saudara"`
	IdKelurahan        int    `gorm:"column:id_kelurahan_saudara"`
	IdKodepos          int    `gorm:"column:id_kode_pos_saudara"`
	NoTelp             string `gorm:"column:no_telp_saudara"`
	NoMobile           string `gorm:"column:no_mobile_saudara"`
	AhliWarisSaudara   string
	HubunganSaudara    string
	DEntry             time.Time
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (NasabahPeroranganSaudara) TableName() string {
	return "m_nasabah_perorangan_saudara"
}
