package nasabah

import "time"

type Nasabah struct {
	IdNasabahPerorangan        int
	DEntry                     time.Time
	NomorRegisterNasabah       string
	NamaIdentitas              string
	GelarDepan                 string
	NamaNasabah                string
	GelarBelakang              string
	IdGolDeb                   int //GolDeb
	TipeIdentitas              int
	NoIdentitas                string
	TglKadaluarsaIdentitas     time.Time `json:"tgl_kadaluarsa_identitas"`
	StatusSeumurHidupIdentitas bool
	NoNpwp                     string
	NoMobile                   string
	NoTelp                     string
	AlamatIdentitas            string
	IdPropinsiIdentitas        int // Propinsi
	IdKotaKabupatenIdentitas   int // KotaKabupaten
	IdKecamatanIdentitas       int // Kecamatan
	IdKelurahanIdentitas       int // Kelurahan
	IdKodePosIdentitas         int // KodePos
	AlamatDomisili             string
	IdPropinsiDomisili         int
	IdKotaKabupatenDomisili    int // KotaKabupaten
	IdKecamatanDomisili        int // Kecamatan
	IdKelurahanDomisili        int // Kelurahan
	IdKodePosDomisili          int // KodePos
	AlamatEmail                string
	IdBpr                      int       // Bpr
	BprCabang                  int       // BprCabang
	IsBlacklist                bool      // not null
	DBlacklist                 time.Time `json:"d_blacklist"`
	DomisiliSesuaiIdentitas    bool
	KeteranganDiblacklist      string
	SaldoTabunganDeposito      int64
	StatusApprove              bool
}

// function to custome table name
func (*Nasabah) TableName() string {
	return "m_nasabah_perorangan_umum"
}

// No args Constructor
func NewNasabah() *Nasabah {
	nasabah := new(Nasabah)
	nasabah.SaldoTabunganDeposito = 0
	return nasabah
}
