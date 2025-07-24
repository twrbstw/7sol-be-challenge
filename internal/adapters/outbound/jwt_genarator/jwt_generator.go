package jwtgenarator

import (
	"seven-solutions-challenge/internal/app/ports"
	"seven-solutions-challenge/internal/domain"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtGenerator struct {
	appCfg domain.AppConfig
}

func NewJwtGenerator(appCfg domain.AppConfig) ports.IJwtGenerator {
	return &JwtGenerator{
		appCfg: appCfg,
	}
}

// GenerateJwt implements ports.IJwtGenerator.
func (j *JwtGenerator) GenerateJwt(name string, email string) (string, error) {
	timeout, err := strconv.Atoi(j.appCfg.TokenTimeout)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"name":  name,
		"email": email,
		"exp":   time.Now().Add(time.Minute * time.Duration(timeout)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.appCfg.SecretKey))
}
