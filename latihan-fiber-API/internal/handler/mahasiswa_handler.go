
package handler

import (
	"sort"
	"strconv"

	"github.com/gofiber/fiber/v3"

	"latihan-fiber-api/internal/model"
)

// DaftarMahasiswa menampilkan daftar mahasiswa
// dengan fitur filter angkatan dan pengurutan nama atau IPK.
func DaftarMahasiswa(c fiber.Ctx) error {

	// Mengambil parameter dari URL
	angkatan := c.Query("angkatan", "")
	urut := c.Query("urut", "")

	// Menyalin data mahasiswa agar data asli tidak berubah
	daftar := make([]model.Mahasiswa, len(model.Daftar))
	copy(daftar, model.Daftar)

	// Filter mahasiswa berdasarkan angkatan
	if angkatan != "" {

		// Mengubah angkatan dari string menjadi integer
		nilaiAngkatan, err := strconv.Atoi(angkatan)

		// Jika angkatan bukan angka, tampilkan error 400
		if err != nil {
			return fiber.NewError(
				fiber.StatusBadRequest,
				"Parameter angkatan harus berupa angka",
			)
		}

		// Menyimpan mahasiswa yang sesuai dengan angkatan
		hasil := []model.Mahasiswa{}

		for _, mahasiswa := range daftar {
			if mahasiswa.Angkatan == nilaiAngkatan {
				hasil = append(hasil, mahasiswa)
			}
		}

		// Mengganti daftar dengan hasil filter
		daftar = hasil
	}

	// Mengurutkan mahasiswa berdasarkan parameter urut
	switch urut {

	case "nama":
		// Mengurutkan nama mahasiswa dari A sampai Z
		sort.Slice(daftar, func(i, j int) bool {
			return daftar[i].Nama < daftar[j].Nama
		})

	case "ipk":
		// Mengurutkan IPK dari terkecil ke terbesar
		sort.Slice(daftar, func(i, j int) bool {
			return daftar[i].Ipk < daftar[j].Ipk
		})

	case "":
		// Jika parameter urut kosong, gunakan urutan awal

	default:
		// Jika parameter urut tidak sesuai, tampilkan error 400
		return fiber.NewError(
			fiber.StatusBadRequest,
			"Parameter urut harus nama atau ipk",
		)
	}

	// Mengembalikan daftar mahasiswa dalam format JSON
	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   daftar,
	})
}

// DetailMahasiswa digunakan untuk mencari mahasiswa berdasarkan ID
func DetailMahasiswa(c fiber.Ctx) error {

	// Mengambil ID dari URL dan mengubahnya menjadi integer
	id, err := strconv.Atoi(c.Params("id"))

	// Jika ID bukan angka, tampilkan error 400
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"ID harus berupa angka",
		)
	}

	// Mencari mahasiswa berdasarkan ID
	for _, mahasiswa := range model.Daftar {

		if mahasiswa.ID == id {

			// Jika ditemukan, tampilkan data mahasiswa
			return c.JSON(fiber.Map{
				"sukses": true,
				"data":   mahasiswa,
			})
		}
	}

	// Jika ID tidak ditemukan, tampilkan error 404
	return fiber.NewError(
		fiber.StatusNotFound,
		"Mahasiswa tidak ditemukan",
	)
}

// CariMahasiswaNim digunakan untuk mencari mahasiswa berdasarkan NIM
func CariMahasiswaNim(c fiber.Ctx) error {

	// Mengambil NIM dari parameter URL
	nim := c.Params("nim")

	// Mengecek setiap mahasiswa dalam daftar
	for _, mahasiswa := range model.Daftar {

		// Membandingkan NIM mahasiswa dengan NIM dari URL
		if mahasiswa.Nim == nim {

			// Jika ditemukan, tampilkan data mahasiswa
			return c.JSON(fiber.Map{
				"sukses": true,
				"data":   mahasiswa,
			})
		}
	}

	// Jika NIM tidak ditemukan, tampilkan error 404
	return fiber.NewError(
		fiber.StatusNotFound,
		"Mahasiswa dengan NIM tersebut tidak ditemukan",
	)
}
