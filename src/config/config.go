package config

import (
	"os"
	"seven-solutions-challenge/src/models"
)

func LoadDefaultConfig() models.Configs {
	return models.Configs{
		DbConfig: models.DbConfig{
			Uri:  os.Getenv("MONGO_URI"),
			Name: os.Getenv("MONGO_NAME"),
		},
		LoggerConfig: models.LoggerConfig{
			Format:     os.Getenv("LOGGER_FORMAT"),
			TimeFormat: os.Getenv("LOGGER_TIME_FORMAT"),
			TimeZone:   os.Getenv("LOGGER_TIME_ZONE"),
		},
		AppConfig: models.AppConfig{
			TokenTimeout: os.Getenv("APP_TOKEN_TIMEOUT"),
			SecretKey:    os.Getenv("APP_SECRET_KEY"),
		},
	}
}
