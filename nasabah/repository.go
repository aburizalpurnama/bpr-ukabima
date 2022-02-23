package nasabah

import (
	"log"

	"gorm.io/gorm"
)

type Repository interface {
	Save(nasabah Nasabah) (Nasabah, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(nasabah Nasabah) (Nasabah, error) {
	err := r.db.Create(&nasabah).Error

	if err != nil {
		log.Fatal(err.Error())
		return nasabah, err
	}

	return nasabah, nil
}
