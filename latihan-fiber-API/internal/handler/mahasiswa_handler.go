package handler

import (
	"strconv" 

	"github.com/gofiber/fiber/v3"

	"latihan-fiber-api/internal/model"
)

func DaftarMahasiswa(c fiber.Ctx) error {

	angkatan := c.Query("angkatan", "")

	if angkatan == "" {
		return c.JSON(fiber.Map{
			"sukses": true,
			"data":   model.Daftar,
		})
	}

	nilaiAngkatan, err := strconv.Atoi(angkatan)

	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"Parameter angkatan harus berupa angka",
		)
	}

	hasil := []model.Mahasiswa{}

	for i := 0; i < len(model.Daftar); i++ {
		if model.Daftar[i].Angkatan == nilaiAngkatan {
			hasil = append(hasil, model.Daftar[i])
		}
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   hasil,
	})
}

func DetailMahasiswa(c fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"ID harus berupa angka",
		)
	}

	for i := 0; i < len(model.Daftar); i++ {
		if model.Daftar[i].ID == id {
			return c.JSON(fiber.Map{
				"sukses": true,
				"data":   model.Daftar[i],
			})
		}
	}

	return fiber.NewError(
		fiber.StatusNotFound,
		"Mahasiswa tidak ditemukan",
	)
}


// CariMahasiswaNim digunakan untuk mencari mahasiswa berdasarkan NIM
func CariMahasiswaNim(c fiber.Ctx) error {

    // Mengambil NIM dari URL
    nim := c.Params("nim")

    // Mengecek setiap mahasiswa dalam daftar
    for _, mahasiswa := range model.Daftar {

        // Jika NIM mahasiswa sama dengan NIM yang dicari
        if mahasiswa.Nim == nim {

            // Mengembalikan data mahasiswa dalam format JSON
            return c.JSON(fiber.Map{
                "sukses": true,
                "data":   mahasiswa,
            })
        }
    }

    // Jika NIM tidak ditemukan, mengembalikan error 404
    return fiber.NewError(
        fiber.StatusNotFound,
        "Mahasiswa dengan NIM tersebut tidak ditemukan",
    )
}
