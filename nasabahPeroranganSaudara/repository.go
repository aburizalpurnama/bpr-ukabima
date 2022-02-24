package nasabahPeroranganSaudara

import (
	"log"

	"gorm.io/gorm"
)

type Repository interface {
	Save(nasabahPeroranganSaudara NasabahPeroranganSaudara) (NasabahPeroranganSaudara, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(nasabahPeroranganSaudara NasabahPeroranganSaudara) (NasabahPeroranganSaudara, error) {
	err := r.db.Save(&nasabahPeroranganSaudara).Error

	if err != nil {
		log.Fatal(err.Error())
		return nasabahPeroranganSaudara, err
	}

	return nasabahPeroranganSaudara, nil
}
