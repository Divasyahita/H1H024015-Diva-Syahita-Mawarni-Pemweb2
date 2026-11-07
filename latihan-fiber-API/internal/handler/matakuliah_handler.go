package handler

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"latihan-fiber-api/internal/model"
	"latihan-fiber-api/internal/repository"
)

// GET: Menampilkan semua mata kuliah
func DaftarMataKuliah(c fiber.Ctx) error {
	data, err := repository.AmbilSemuaMataKuliah()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal mengambil data mata kuliah",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   data,
	})
}

// GET: Menampilkan detail mata kuliah
func DetailMataKuliah(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "ID mata kuliah tidak valid",
		})
	}

	data, err := repository.AmbilMataKuliahPerId(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Mata kuliah tidak ditemukan",
		})
	}
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal mengambil data mata kuliah",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   data,
	})
}

// POST: Menambahkan mata kuliah
func TambahMataKuliah(c fiber.Ctx) error {
	var permintaan PermintaanMataKuliah

	if err := c.Bind().Body(&permintaan); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Format data tidak valid",
		})
	}

	if permintaan.ProgramStudiID == 0 ||
		permintaan.Kode == "" ||
		permintaan.Nama == "" ||
		permintaan.SKS <= 0 ||
		permintaan.Semester <= 0 {
		return c.Status(400).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Semua field wajib diisi dengan benar",
		})
	}

	data := model.MataKuliah{
		ProgramStudiID: permintaan.ProgramStudiID,
		Kode:           permintaan.Kode,
		Nama:           permintaan.Nama,
		SKS:            permintaan.SKS,
		Semester:       permintaan.Semester,
	}

	if err := repository.SimpanMataKuliah(&data); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal menyimpan data mata kuliah",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Data mata kuliah berhasil ditambahkan",
		"data":   data,
	})
}

// PUT: Memperbarui mata kuliah
func PerbaruiMataKuliah(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "ID mata kuliah tidak valid",
		})
	}

	data, err := repository.AmbilMataKuliahPerId(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Mata kuliah tidak ditemukan",
		})
	}
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal mengambil data mata kuliah",
		})
	}

	var permintaan PermintaanMataKuliah
	if err := c.Bind().Body(&permintaan); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Format data tidak valid",
		})
	}

	if permintaan.ProgramStudiID == 0 ||
		permintaan.Kode == "" ||
		permintaan.Nama == "" ||
		permintaan.SKS <= 0 ||
		permintaan.Semester <= 0 {
		return c.Status(400).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Semua field wajib diisi dengan benar",
		})
	}

	data.ProgramStudiID = permintaan.ProgramStudiID
	data.Kode = permintaan.Kode
	data.Nama = permintaan.Nama
	data.SKS = permintaan.SKS
	data.Semester = permintaan.Semester

	if err := repository.PerbaruiMataKuliah(&data); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal memperbarui data mata kuliah",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Data mata kuliah berhasil diperbarui",
		"data":   data,
	})
}

// DELETE: Menghapus mata kuliah
func HapusMataKuliah(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "ID mata kuliah tidak valid",
		})
	}

	err = repository.HapusMataKuliah(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Mata kuliah tidak ditemukan",
		})
	}
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal menghapus data mata kuliah",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Data mata kuliah berhasil dihapus",
	})
}
