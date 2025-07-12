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
	}
}
