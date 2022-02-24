package nasabahPeroranganPersonal

import (
	"log"

	"gorm.io/gorm"
)

type Repository interface {
	Save(nasabahPeroranganPersonal NasabahPeroranganPersonal) (NasabahPeroranganPersonal, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(nasabahPeroranganPersonal NasabahPeroranganPersonal) (NasabahPeroranganPersonal, error) {
	err := r.db.Save(&nasabahPeroranganPersonal).Error

	if err != nil {
		log.Fatal(err.Error())
		return nasabahPeroranganPersonal, err
	}

	return nasabahPeroranganPersonal, nil
}
