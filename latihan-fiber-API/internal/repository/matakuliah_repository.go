package repository

import (
	"gorm.io/gorm"
	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/model"
)

func AmbilSemuaMataKuliah() ([]model.MataKuliah, error) {
	var mataKuliah []model.MataKuliah

	err := config.DB.Preload("ProgramStudi").
		Order("id ASC").
		Find(&mataKuliah).Error

	return mataKuliah, err
}

func AmbilMataKuliahPerId(id uint) (model.MataKuliah, error) {
	var mataKuliah model.MataKuliah

	err := config.DB.Preload("ProgramStudi").
		First(&mataKuliah, id).Error

	return mataKuliah, err
}

func SimpanMataKuliah(mataKuliah *model.MataKuliah) error {
	return config.DB.Create(mataKuliah).Error
}

func PerbaruiMataKuliah(mataKuliah *model.MataKuliah) error {
	return config.DB.Save(mataKuliah).Error
}

func HapusMataKuliah(id uint) error {
	hasil := config.DB.Delete(&model.MataKuliah{}, id)

	if hasil.Error != nil {
		return hasil.Error
	}

	if hasil.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
