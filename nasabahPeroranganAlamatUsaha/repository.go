package nasabahPeroranganAlamatUsaha

import (
	"log"

	"gorm.io/gorm"
)

type Repository interface {
	Save(nasabahPeroranganAlamatUsaha NasabahPeroranganAlamatUsaha) (NasabahPeroranganAlamatUsaha, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(nasabahPeroranganAlamatUsaha NasabahPeroranganAlamatUsaha) (NasabahPeroranganAlamatUsaha, error) {
	err := r.db.Save(&nasabahPeroranganAlamatUsaha).Error

	if err != nil {
		log.Fatal(err.Error())
		return nasabahPeroranganAlamatUsaha, err
	}

	return nasabahPeroranganAlamatUsaha, nil
}
