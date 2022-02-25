package nasabahPeroranganPekerjaan

import (
	"log"

	"gorm.io/gorm"
)

type Repository interface {
	Save(nasabahPeroranganPekerajaan NasabahPeroranganPekerjaan) (NasabahPeroranganPekerjaan, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(nasabahPeroranganPekerajaan NasabahPeroranganPekerjaan) (NasabahPeroranganPekerjaan, error) {
	err := r.db.Save(&nasabahPeroranganPekerajaan).Error

	if err != nil {
		log.Fatal(err.Error())
		return nasabahPeroranganPekerajaan, err
	}

	return nasabahPeroranganPekerajaan, nil
}
