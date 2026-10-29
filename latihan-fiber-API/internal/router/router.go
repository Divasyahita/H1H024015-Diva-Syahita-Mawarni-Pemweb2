
package router

import (
	"github.com/gofiber/fiber/v3"

	"latihan-fiber-api/internal/handler"
)

// Fungsi untuk mendaftarkan semua endpoint API
func Daftarkan(app *fiber.App) {

	// Membuat grup route /api
	api := app.Group("/api")

	// Endpoint untuk mengecek status API
	api.Get("/status", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"sukses": true,
			"pesan":  "API Fiber aktif",
		})
	})

	mahasiswa := api.Group("/mahasiswa")

	mahasiswa.Get("/", handler.DaftarMahasiswa)

	mahasiswa.Get("/nim/:nim", handler.CariMahasiswaNim)

	mahasiswa.Get("/:id", handler.DetailMahasiswa)
}
