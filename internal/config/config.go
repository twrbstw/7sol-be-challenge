package config

import (
	"os"
	"seven-solutions-challenge/internal/domain"
)

func LoadDefaultConfig() domain.Configs {
	return domain.Configs{
		DbConfig: domain.DbConfig{
			Uri:  os.Getenv("MONGO_URI"),
			Name: os.Getenv("MONGO_NAME"),
		},
		LoggerConfig: domain.LoggerConfig{
			Format:     os.Getenv("LOGGER_FORMAT"),
			TimeFormat: os.Getenv("LOGGER_TIME_FORMAT"),
			TimeZone:   os.Getenv("LOGGER_TIME_ZONE"),
		},
		AppConfig: domain.AppConfig{
			TokenTimeout: os.Getenv("APP_TOKEN_TIMEOUT"),
			SecretKey:    os.Getenv("APP_SECRET_KEY"),
		},
	}
}
