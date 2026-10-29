
package router

import (
	"github.com/gofiber/fiber/v3"

	"latihan-fiber-api/internal/handler"
	"latihan-fiber-api/internal/middleware"
)

// Daftarkan digunakan untuk mendaftarkan seluruh endpoint API
func Daftarkan(app *fiber.App) {

	// Membuat grup endpoint /api
	api := app.Group("/api")

	// Endpoint status tetap bisa diakses tanpa API Key
	api.Get("/status", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"sukses": true,
			"pesan":  "API Fiber aktif",
		})
	})

	mahasiswa := api.Group("/mahasiswa")

	mahasiswa.Use(middleware.KunciApi())

	mahasiswa.Get("/", handler.DaftarMahasiswa)

	mahasiswa.Get("/nim/:nim", handler.CariMahasiswaNim)
	mahasiswa.Get("/:id", handler.DetailMahasiswa)
}
