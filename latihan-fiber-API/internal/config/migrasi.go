package config

import (
	"log"

	"latihan-fiber-api/internal/model"
)

func Migrasi() {
	err := DB.AutoMigrate(
		&model.ProgramStudi{},
		&model.Mahasiswa{},
		&model.MataKuliah{},
	)

	if err != nil {
		log.Fatal("Gagal melakukan migrasi database:", err)
	}

	log.Println("Migrasi database berhasil!")
}

func DataAwal() {
	var jumlah int64

	DB.Model(&model.ProgramStudi{}).Count(&jumlah)

	if jumlah > 0 {
		return
	}

	programStudi := []model.ProgramStudi{
		{
			Kode:    "TK",
			Nama:    "Teknik Komputer",
			Jenjang: "S1",
		},
		{
			Kode:    "IF",
			Nama:    "Informatika",
			Jenjang: "S1",
		},
		{
			Kode:    "TE",
			Nama:    "Teknik Elektro",
			Jenjang: "S1",
		},
	}

	if err := DB.Create(&programStudi).Error; err != nil {
		log.Fatal("Gagal memasukkan data awal:", err)
	}

	log.Println("Data awal program studi berhasil ditambahkan!")
}
