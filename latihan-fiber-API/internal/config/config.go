package config

import (
	"os" // Membaca environment variable

	"github.com/joho/godotenv" // Membaca file .env
)

// Struct untuk menyimpan konfigurasi aplikasi
type Konfigurasi struct {
	NamaAplikasi string
	Port         string
	Lingkungan   string
}

// Fungsi untuk memuat konfigurasi dari file .env
func Muat() Konfigurasi {
	godotenv.Load()

	return Konfigurasi{
		NamaAplikasi: ambil("APP_NAME", "Fiber API"),
		Port:         ambil("APP_PORT", "3000"),
		Lingkungan:   ambil("APP_ENV", "local"),
	}
}

// Fungsi untuk mengambil nilai konfigurasi
func ambil(kunci string, bawaan string) string {
	nilai := os.Getenv(kunci)

	// Jika nilai kosong, gunakan nilai bawaan
	if nilai == "" {
		return bawaan
	}

	return nilai
}
