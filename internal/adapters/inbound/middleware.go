package inbound

import (
	"seven-solutions-challenge/internal/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
)

func NewLoggerMiddleware(lcfg domain.LoggerConfig) fiber.Handler {
	return logger.New(
		logger.Config{
			Format:     lcfg.Format,
			TimeFormat: lcfg.TimeFormat,
			TimeZone:   lcfg.TimeZone,
		},
	)
}

func NewAuthMiddleware(appCfg domain.AppConfig) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(appCfg.SecretKey),
	})
}
