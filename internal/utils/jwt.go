package utils

import (
	"time"

	"github.com/djwhocodes/hostel_saas/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID, tenantID, role string) (string, error) {
	claims := jwt.MapClaims{
		"userId":   userID,
		"tenantId": tenantID,
		"role":     role,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}
