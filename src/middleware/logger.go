package middleware

import (
	"seven-solutions-challenge/src/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewLoggerMiddleware(lcfg models.LoggerConfig) fiber.Handler {
	return logger.New(
		logger.Config{
			Format:     lcfg.Format,
			TimeFormat: lcfg.TimeFormat,
			TimeZone:   lcfg.TimeZone,
		},
	)
}
