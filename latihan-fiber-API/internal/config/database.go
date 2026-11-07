package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func HubungkanBasisData(konf Konfigurasi) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		konf.DbUser,
		konf.DbSandi,
		konf.DbHost,
		konf.DbPort,
		konf.DbNama,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal menghubungkan database:", err)
	}

	DB = db

	log.Println("Berhasil terhubung ke database MySQL!")
}
