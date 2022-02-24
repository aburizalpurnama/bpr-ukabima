package nasabahPeroranganInfo

import (
	"log"

	"gorm.io/gorm"
)

type Repository interface {
	Save(nasabahPeroranganInfo NasabahPeroranganInfo) (NasabahPeroranganInfo, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(nasabahPeroranganInfo NasabahPeroranganInfo) (NasabahPeroranganInfo, error) {
	err := r.db.Save(&nasabahPeroranganInfo).Error

	if err != nil {
		log.Fatal(err.Error())
		return nasabahPeroranganInfo, err
	}

	return nasabahPeroranganInfo, nil
}
