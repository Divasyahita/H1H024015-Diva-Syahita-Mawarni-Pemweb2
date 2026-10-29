
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

	// Membuat grup route mahasiswa
	mahasiswa := api.Group("/mahasiswa")

	// Menampilkan semua mahasiswa
	mahasiswa.Get("/", handler.DaftarMahasiswa)

	// Menampilkan detail mahasiswa berdasarkan ID
	mahasiswa.Get("/:id", handler.DetailMahasiswa)
}
