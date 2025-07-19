package middleware

import (
	"seven-solutions-challenge/src/models"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func NewAuthMiddleware(appCfg models.AppConfig) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(appCfg.SecretKey),
	})
}
