package router

import (
	"github.com/gofiber/fiber/v3"
	"latihan-fiber-api/internal/handler"
	"latihan-fiber-api/internal/middleware"
)

func Daftarkan(app *fiber.App) {

	// Endpoint status API
	app.Get("/api/status", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"sukses": true,
			"pesan":  "API Fiber aktif",
		})
	})

	// Endpoint Mahasiswa
	mahasiswa := app.Group("/api/mahasiswa")
	mahasiswa.Use(middleware.KunciApi())

	mahasiswa.Get("/", handler.DaftarMahasiswa)
	mahasiswa.Get("/nim/:nim", handler.CariMahasiswaNim)
	mahasiswa.Post("/", handler.TambahMahasiswa)
	mahasiswa.Get("/:id", handler.DetailMahasiswa)
	mahasiswa.Put("/:id", handler.PerbaruiMahasiswa)
	mahasiswa.Delete("/:id", handler.HapusMahasiswa)

	// Endpoint pemulihan mahasiswa
	mahasiswa.Post("/:id/pulihkan", handler.PulihkanMahasiswa)

	// Endpoint Mata Kuliah
	mataKuliah := app.Group("/api/mata-kuliah")
	mataKuliah.Use(middleware.KunciApi())

	mataKuliah.Get("/", handler.DaftarMataKuliah)
	mataKuliah.Post("/", handler.TambahMataKuliah)
	mataKuliah.Get("/:id", handler.DetailMataKuliah)
	mataKuliah.Put("/:id", handler.PerbaruiMataKuliah)
	mataKuliah.Delete("/:id", handler.HapusMataKuliah)

	// Endpoint Program Studi
	app.Get("/api/program-studi", middleware.KunciApi(), handler.DaftarProgramStudi)
}
