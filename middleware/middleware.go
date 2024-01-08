package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hikmahnurarifah/webbackend/util"
)

func IsAuthenticate(c *fiber.Ctx) error {
	// Ambil nilai JWT dari cookie dengan nama "jwt"
	cookie := c.Cookies("jwt")

	// Coba parse JWT menggunakan fungsi Parsejwt dari paket util
	if _, err := util.Parsejwt(cookie); err != nil {
		// Jika parsing gagal, respon dengan status 401 Unauthorized
		// dan kirimkan respons JSON yang menyatakan bahwa pengguna tidak terotentikasi
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	// Jika parsing JWT berhasil, lanjutkan ke middleware atau handler berikutnya
	return c.Next()
}
