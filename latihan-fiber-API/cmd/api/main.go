package main



import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"

	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/middleware"
	router "latihan-fiber-api/internal/router"
)


func main() {
	
	konf := config.Muat()

	app := fiber.New(fiber.Config{
		AppName:      konf.NamaAplikasi,
		ErrorHandler: penanganGalat,
	})

	app.Use(requestid.New()) // Memberikan ID unik untuk request
	app.Use(logger.New())    // Mencatat aktivitas request
	app.Use(recoverer.New()) // Menangani panic agar server tidak berhenti

	// Memasang middleware buatan sendiri
	app.Use(middleware.WaktuProses())

	router.Daftarkan(app)

	app.Use(func(c fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Endpoint tidak tersedia",
		})
	})

	log.Fatal(app.Listen(":" + konf.Port))

}

// Fungsi untuk menangani error secara terpusat
func penanganGalat(c fiber.Ctx, err error) error {

	// Nilai default jika terjadi error server
	kode := fiber.StatusInternalServerError
	pesan := "Terjadi kesalahan pada server"

	var galatFiber *fiber.Error

	// Memeriksa apakah error berasal dari Fiber
	if errors.As(err, &galatFiber) {
		kode = galatFiber.Code
		pesan = galatFiber.Message
	}

	// Mengirim respons error dalam format JSON
	return c.Status(kode).JSON(fiber.Map{
		"sukses": false,
		"pesan":  pesan,
	})
}
