package utils

import (
	"seven-solutions-challenge/src/models"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwt(name, email string, appCfg models.AppConfig) (string, error) {
	timeout, err := strconv.Atoi(appCfg.TokenTimeout)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"name":  name,
		"email": email,
		"exp":   time.Now().Add(time.Minute * time.Duration(timeout)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(appCfg.SecretKey))
}
