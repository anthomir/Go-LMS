package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateJWT(userID uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": userID,
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
