package config

import (
	"os"
	"seven-solutions-assignment/models"
)

func LoadDefaultConfig() models.Configs {
	return models.Configs{
		DbConfig: models.DbConfig{
			Uri:  os.Getenv("MONGO_URI"),
			Name: os.Getenv("MONGO_NAME"),
		},
	}
}
