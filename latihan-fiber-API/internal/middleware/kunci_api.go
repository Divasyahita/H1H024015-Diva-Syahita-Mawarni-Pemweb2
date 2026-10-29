
package middleware

import (
	"os"

	"github.com/gofiber/fiber/v3"
)

// KunciApi memeriksa API Key yang dikirim melalui header
func KunciApi() fiber.Handler {

	return func(c fiber.Ctx) error {

		// Mengambil API Key yang benar dari file .env
		kunciBenar := os.Getenv("API_KEY")

		// Mengambil API Key yang dikirim oleh pengguna
		kunciPengguna := c.Get("X-Api-Key")

		// Jika API Key belum dikonfigurasi di server
		if kunciBenar == "" {
			return fiber.NewError(
				fiber.StatusInternalServerError,
				"API Key server belum dikonfigurasi",
			)
		}

		// Jika API Key kosong atau tidak sesuai
		if kunciPengguna == "" || kunciPengguna != kunciBenar {
			return fiber.NewError(
				fiber.StatusUnauthorized,
				"API Key tidak valid atau tidak diberikan",
			)
		}

		// Jika API Key benar, lanjutkan ke handler
		return c.Next()
	}
}
