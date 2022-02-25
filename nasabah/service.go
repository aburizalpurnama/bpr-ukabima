package nasabah

import (
	"bprukabima/nasabahPeroranganAlamatUsaha"
	"bprukabima/nasabahPeroranganInfo"
	"bprukabima/nasabahPeroranganPekerjaan"
	"bprukabima/nasabahPeroranganPersonal"
	"bprukabima/nasabahPeroranganSaudara"
	"log"
	"strconv"
	"time"
)

type NasabahService interface {
	RegisterNasabah()
}

type nasabahService struct {
	nGeneralRepo     Repository
	nAlamatUsahaRepo nasabahPeroranganAlamatUsaha.Repository
	nInfoRepo        nasabahPeroranganInfo.Repository
	nPersonalRepo    nasabahPeroranganPersonal.Repository
	nPekerjaanRepo   nasabahPeroranganPekerjaan.Repository
	nSaudaraRepo     nasabahPeroranganSaudara.Repository
}

func NewService(
	nGeneralRepo Repository,
	nAlamatUsahaRepo nasabahPeroranganAlamatUsaha.Repository,
	nInfoRepo nasabahPeroranganInfo.Repository,
	nPersonalRepo nasabahPeroranganPersonal.Repository,
	nPekerjaanRepo nasabahPeroranganPekerjaan.Repository,
	nSaudaraRepo nasabahPeroranganSaudara.Repository) *nasabahService {

	nasabahService := &nasabahService{}
	nasabahService.nGeneralRepo = nGeneralRepo
	nasabahService.nAlamatUsahaRepo = nAlamatUsahaRepo
	nasabahService.nInfoRepo = nInfoRepo
	nasabahService.nPersonalRepo = nPersonalRepo
	nasabahService.nPekerjaanRepo = nPekerjaanRepo
	nasabahService.nSaudaraRepo = nSaudaraRepo

	return nasabahService

}

func (s *nasabahService) RegisterNasabah(input InputRegisterNasabah) (NasabahRegisterResponse, error) {

	// create response object
	response := NasabahRegisterResponse{}

	// insert nasabahGeneral
	nGeneral := Nasabah{}

	nGeneral.DEntry = time.Now()
	nGeneral.IdEntry = "BimaApps"
	nGeneral.IsBlacklist = false
	nGeneral.StatusApprove = true
	nGeneral.IdBpr = input.IdKodeBpr
	nGeneral.BprCabang = input.IdKodeCabang
	nGeneral.NoIdentitas = input.Nik
	nGeneral.NoMobile = input.Mobile
	nGeneral.NoTelp = input.Mobile
	nGeneral.NamaNasabah = input.Nama
	nGeneral.NamaIdentitas = input.Nama
	nGeneral.AlamatEmail = input.Email
	nGeneral.AlamatDomisili = input.Alamat
	nGeneral.AlamatIdentitas = input.Alamat
	nGeneral.TipeIdentitas = 0
	nGeneral.DomisiliSesuaiIdentitas = true
	nGeneral.IdPropinsiDomisili = input.IdPropinsi
	nGeneral.IdPropinsiIdentitas = input.IdPropinsi
	nGeneral.IdKotaKabupatenDomisili = input.IdKotaKab
	nGeneral.IdKotaKabupatenIdentitas = input.IdKotaKab
	nGeneral.IdKecamatanDomisili = input.IdKecamatan
	nGeneral.IdKecamatanIdentitas = input.IdKecamatan
	nGeneral.IdKelurahanDomisili = input.IdKelurahan
	nGeneral.IdKelurahanIdentitas = input.IdKelurahan
	nGeneral.IdKodePosDomisili = input.IdKodePos
	nGeneral.IdKodePosIdentitas = input.IdKodePos
	nGeneral.IdGolDeb = 557
	nGeneral.NomorRegisterNasabah = "999999999" // Harusnya ada bussines proses tersendiri untuk dapar nomor_register_nasabah

	savedNGeneral, generalRepoErr := s.nGeneralRepo.Save(nGeneral)

	if generalRepoErr != nil {
		log.Fatal(generalRepoErr.Error())

		response.code = "404"
		response.status = "Gagal menyimpan nasabah_perorangan_umum"
		return response, generalRepoErr
	}

	// insert nasabahPersonal
	nPersonal := nasabahPeroranganPersonal.NasabahPeroranganPersonal{}

	nPersonal.DEntry = time.Now()
	nPersonal.IdEntry = "BimaApps"
	nPersonal.IdNasabahPerorangan = savedNGeneral.ID
	nPersonal.IdPendidikan = 393 // Harusnya ambil dari tabel m_pendidikan
	nPersonal.IdAgama = 94       // harusnya ambil dari tabel m_agama

	// convert jenis kelamin
	if input.Kelamin == "Laki-laki" {
		nPersonal.JenisKelamin = 0
	} else {
		nPersonal.JenisKelamin = 1
	}

	nPersonal.IdPropinsiLahir = savedNGeneral.IdPropinsiIdentitas
	nPersonal.IdPropinsiLahirPasangan = savedNGeneral.IdPropinsiIdentitas
	nPersonal.IdKotaKabupatenLahir = savedNGeneral.IdKotaKabupatenIdentitas
	nPersonal.IdKotaKabupatenLahirPasangan = savedNGeneral.IdKotaKabupatenIdentitas

	_, personalRepoErr := s.nPersonalRepo.Save(nPersonal)

	if personalRepoErr != nil {
		log.Fatal(personalRepoErr.Error())

		response.code = "404"
		response.status = "Gagal menyimpan nasabah_perorangan_personal"
		return response, personalRepoErr
	}

	// insert nasabahPekerajaan
	nPekerjaan := nasabahPeroranganPekerjaan.NasabahPeroranganPekerjaan{}

	nPekerjaan.DEntry = time.Now()
	nPekerjaan.IdEntry = "BimaApps"
	nPekerjaan.IdNasabah = savedNGeneral.ID
	nPekerjaan.TempatPekerjaan = "BimaApps"
	nPekerjaan.RangePenghasilanPerbulan = 3
	nPekerjaan.IdGroupAnggota = 4233 // Harusnya ambil dari m_group_instansi_anggota
	nPekerjaan.IdBidangUsaha = 13
	nPekerjaan.IdPekerjaan = 216

	_, pekerjaanRepoErr := s.nPekerjaanRepo.Save(nPekerjaan)

	if pekerjaanRepoErr != nil {
		log.Fatal(pekerjaanRepoErr.Error())

		response.code = "404"
		response.status = "Gagal menyimpan nasabah_perorangan_pekerjaan"
		return response, pekerjaanRepoErr
	}

	// insert nasabahAlamatUsaha
	nAlamatUsaha := nasabahPeroranganAlamatUsaha.NasabahPeroranganAlamatUsaha{}

	nAlamatUsaha.DEntry = time.Now()
	nAlamatUsaha.IdEntry = "BimaApps"
	nAlamatUsaha.IdNasabah = savedNGeneral.ID
	nAlamatUsaha.IdPropinsi = savedNGeneral.IdPropinsiIdentitas
	nAlamatUsaha.IdKecamatan = savedNGeneral.IdKecamatanIdentitas
	nAlamatUsaha.IdKelurahan = savedNGeneral.IdKelurahanIdentitas
	nAlamatUsaha.IdKotaKab = savedNGeneral.IdKotaKabupatenIdentitas
	nAlamatUsaha.IdKodePos = savedNGeneral.IdKodePosIdentitas

	_, alamatURepoErr := s.nAlamatUsahaRepo.Save(nAlamatUsaha)

	if alamatURepoErr != nil {
		log.Fatal(alamatURepoErr.Error())

		response.code = "404"
		response.status = "Gagal menyimpan nasabah_perorangan_alamat_usaha"
		return response, alamatURepoErr
	}

	// insert nasabahSaudara
	nSaudara := nasabahPeroranganSaudara.NasabahPeroranganSaudara{}

	nSaudara.DEntry = time.Now()
	nSaudara.IdEntry = "BimaApps"
	nSaudara.IdNasabah = savedNGeneral.ID
	nSaudara.NamaSaudara = savedNGeneral.NamaNasabah
	nSaudara.IdKecamatan = savedNGeneral.IdKecamatanIdentitas
	nSaudara.IdKelurahan = savedNGeneral.IdKelurahanIdentitas
	nSaudara.IdKotaKab = savedNGeneral.IdKotaKabupatenIdentitas
	nSaudara.IdPropinsi = savedNGeneral.IdPropinsiIdentitas
	nSaudara.IdKodepos = savedNGeneral.IdKodePosIdentitas
	nSaudara.NoMobile = savedNGeneral.NoMobile
	nSaudara.NoTelp = savedNGeneral.NoTelp

	_, saudaraRepoErr := s.nSaudaraRepo.Save(nSaudara)

	if saudaraRepoErr != nil {
		log.Fatal(saudaraRepoErr.Error())

		response.code = "404"
		response.status = "Gagal menyimpan nasabah_perorangan_saudara"
		return response, saudaraRepoErr
	}

	// insert nasabahInformasi
	nInfo := nasabahPeroranganInfo.NasabahPeroranganInfo{}

	nInfo.DEntry = time.Now()
	nInfo.IdEntry = "BimaApps"
	nInfo.NamaAliasNasabah = savedNGeneral.NamaIdentitas
	nInfo.IdNasabah = savedNGeneral.ID
	nInfo.IdHubunganPelapor = 19
	nInfo.IdNegara = 62
	nInfo.IdReferensiNasabah = 31655

	_, infoRepoErr := s.nInfoRepo.Save(nInfo)

	if infoRepoErr != nil {
		log.Fatal(infoRepoErr.Error())

		response.code = "404"
		response.status = "Gagal menyimpan nasabah_perorangan_info_lainnya"
		return response, infoRepoErr
	}

	// insert response model
	response.savedIdNasabahPerorangan = strconv.Itoa(savedNGeneral.ID)
	response.savedCif = savedNGeneral.NomorRegisterNasabah
	response.savedNamaNasabah = savedNGeneral.NamaNasabah
	response.savedBpr = strconv.Itoa(savedNGeneral.BprCabang)
	response.code = "200"
	response.status = "BERHASIL"

	return response, nil
}
