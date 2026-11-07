package handler

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"latihan-fiber-api/internal/model"
	"latihan-fiber-api/internal/repository"
)

// GET /api/mahasiswa
func DaftarMahasiswa(c fiber.Ctx) error {
	halaman := 1
	perHalaman := 10

	if nilai := c.Query("halaman"); nilai != "" {
		hasil, err := strconv.Atoi(nilai)
		if err != nil || hasil < 1 {
			return fiber.NewError(fiber.StatusBadRequest, "Nomor halaman tidak valid")
		}
		halaman = hasil
	}

	if nilai := c.Query("per_halaman"); nilai != "" {
		hasil, err := strconv.Atoi(nilai)
		if err != nil || hasil < 1 || hasil > 100 {
			return fiber.NewError(fiber.StatusBadRequest, "Jumlah data per halaman harus 1-100")
		}
		perHalaman = hasil
	}

	angkatan := 0
	if nilai := c.Query("angkatan"); nilai != "" {
		hasil, err := strconv.Atoi(nilai)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Angkatan tidak valid")
		}
		angkatan = hasil
	}

	data, total, err := repository.Ambil(repository.FilterMahasiswa{
		Cari:       c.Query("cari"),
		Angkatan:   angkatan,
		Urut:       c.Query("urut"),
		Halaman:    halaman,
		PerHalaman: perHalaman,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data mahasiswa")
	}

	totalHalaman := (int(total) + perHalaman - 1) / perHalaman

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   data,
		"pagination": fiber.Map{
			"halaman":       halaman,
			"per_halaman":   perHalaman,
			"total_data":    total,
			"total_halaman": totalHalaman,
		},
	})
}

// GET /api/mahasiswa/:id
func DetailMahasiswa(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "ID mahasiswa tidak valid")
	}

	mahasiswa, err := repository.AmbilPerId(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusNotFound, "Mahasiswa tidak ditemukan")
	}
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data mahasiswa")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   mahasiswa,
	})
}

// POST /api/mahasiswa
func TambahMahasiswa(c fiber.Ctx) error {
	var permintaan PermintaanMahasiswa

	if err := c.Bind().Body(&permintaan); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Format data tidak valid")
	}

	if permintaan.ProgramStudiID == 0 ||
		permintaan.Nim == "" ||
		permintaan.Nama == "" ||
		permintaan.Email == "" ||
		permintaan.Angkatan <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Data mahasiswa belum lengkap")
	}

	aktif := true
	if permintaan.Aktif != nil {
		aktif = *permintaan.Aktif
	}

	mahasiswa := model.Mahasiswa{
		ProgramStudiID: permintaan.ProgramStudiID,
		Nim:            permintaan.Nim,
		Nama:           permintaan.Nama,
		Email:          permintaan.Email,
		Angkatan:       permintaan.Angkatan,
		Ipk:            permintaan.Ipk,
		Aktif:          aktif,
	}

	if err := repository.Simpan(&mahasiswa); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal menyimpan data mahasiswa")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Data mahasiswa berhasil ditambahkan",
		"data":   mahasiswa,
	})
}

// PUT /api/mahasiswa/:id
func PerbaruiMahasiswa(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "ID mahasiswa tidak valid")
	}

	mahasiswa, err := repository.AmbilPerId(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusNotFound, "Mahasiswa tidak ditemukan")
	}
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data mahasiswa")
	}

	var permintaan PermintaanMahasiswa
	if err := c.Bind().Body(&permintaan); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Format data tidak valid")
	}

	if permintaan.ProgramStudiID == 0 ||
		permintaan.Nim == "" ||
		permintaan.Nama == "" ||
		permintaan.Email == "" ||
		permintaan.Angkatan <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Data mahasiswa belum lengkap")
	}

	mahasiswa.ProgramStudiID = permintaan.ProgramStudiID
	mahasiswa.Nim = permintaan.Nim
	mahasiswa.Nama = permintaan.Nama
	mahasiswa.Email = permintaan.Email
	mahasiswa.Angkatan = permintaan.Angkatan
	mahasiswa.Ipk = permintaan.Ipk

	if permintaan.Aktif != nil {
		mahasiswa.Aktif = *permintaan.Aktif
	}

	if err := repository.Perbarui(&mahasiswa); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal memperbarui data mahasiswa")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Data mahasiswa berhasil diperbarui",
		"data":   mahasiswa,
	})
}

// DELETE /api/mahasiswa/:id
func HapusMahasiswa(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "ID mahasiswa tidak valid")
	}

	err = repository.Hapus(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Mahasiswa tidak ditemukan")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal menghapus data mahasiswa")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Data mahasiswa berhasil dihapus",
	})
}

// GET /api/mahasiswa/nim/:nim
func CariMahasiswaNim(c fiber.Ctx) error {
	nim := c.Params("nim")

	mahasiswa, err := repository.AmbilPerNim(nim)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusNotFound, "Mahasiswa tidak ditemukan")
	}

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mencari data mahasiswa")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   mahasiswa,
	})
}

// GET /api/program-studi
func DaftarProgramStudi(c fiber.Ctx) error {
	data, err := repository.AmbilProgramStudiDenganJumlah()

	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"Gagal mengambil data program studi",
		)
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   data,
	})
}

// POST /api/mahasiswa/:id/pulihkan
func PulihkanMahasiswa(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "ID mahasiswa tidak valid")
	}

	err = repository.Pulihkan(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Mahasiswa tidak ditemukan")
		}

		if errors.Is(err, gorm.ErrInvalidData) {
			return fiber.NewError(fiber.StatusBadRequest, "Mahasiswa belum dihapus")
		}

		return fiber.NewError(fiber.StatusInternalServerError, "Gagal memulihkan data mahasiswa")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Data mahasiswa berhasil dipulihkan",
	})
}
