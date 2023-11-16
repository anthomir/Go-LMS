package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/anthomir/GoProject/enums"
	"github.com/anthomir/GoProject/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func RequireUserAuth(c *gin.Context) {
	tokenString, err := extractTokenString(c)
	if err != nil {
		handleUnauthorized(c)
		return
	}

	token, err := validateToken(tokenString)
	if err != nil {
		handleUnauthorized(c)
		return
	}

	claims, ok := getTokenClaims(token)
	if !ok {
		handleUnauthorized(c)
		return
	}

	if isTokenExpired(claims) {
		handleUnauthorized(c)
		return
	}

	userId, err := getSub(claims)
	if err != nil {
		handleUnauthorized(c)
		return
	}
	userService := services.NewUserService(&gorm.DB{})
	user, err := userService.FindById(userId, []string{})
	if err != nil || user == nil {
		handleUnauthorized(c)
		return
	}

	c.Set("user", user)
	fmt.Println(user);
	c.Next()
}

func RequireTeacherAuth(c *gin.Context) {
	tokenString, err := extractTokenString(c)
	if err != nil {
		handleUnauthorized(c)
		return
	}

	token, err := validateToken(tokenString)
	if err != nil {
		handleUnauthorized(c)
		return
	}

	claims, ok := getTokenClaims(token)
	if !ok {
		handleUnauthorized(c)
		return
	}

	if isTokenExpired(claims) {
		handleUnauthorized(c)
		return
	}

	userId, err := getSub(claims)
	if err != nil {
		handleUnauthorized(c)
		return
	}
	userService := services.NewUserService(&gorm.DB{})
	user, err := userService.FindById(userId, []string{})
	if err != nil || user == nil || user.Role == enums.RoleUser {
		handleUnauthorized(c)
		return
	}

	c.Set("user", user)
	c.Next()
}

func RequireAdminAuth(c *gin.Context) {
	tokenString, err := extractTokenString(c)
	if err != nil {
		handleUnauthorized(c)
		return
	}

	token, err := validateToken(tokenString)
	if err != nil {
		handleUnauthorized(c)
		return
	}

	claims, ok := getTokenClaims(token)
	if !ok {
		handleUnauthorized(c)
		return
	}

	if isTokenExpired(claims) {
		handleUnauthorized(c)
		return
	}

	userId, err := getSub(claims)
	if err != nil {
		handleUnauthorized(c)
		return
	}
	userService := services.NewUserService(&gorm.DB{})
	user, err := userService.FindById(userId, []string{})
	if err != nil || user == nil || user.Role != enums.RoleAdmin {
		handleUnauthorized(c)
		return
	}

	c.Set("user", user)
	c.Next()
}

/////////////////////////////////////////
//                                     //
//            Helper functions         //   
//                                     //
/////////////////////////////////////////

func extractTokenString(c *gin.Context) (string, error) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return "", errors.New("missing authorization token")
	}
	return tokenString[7:], nil
}

func validateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || method.Alg() != "HS256" {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}

func getTokenClaims(token *jwt.Token) (jwt.MapClaims, bool) {
	claims, ok := token.Claims.(jwt.MapClaims)
	return claims, ok
}

func isTokenExpired(claims jwt.MapClaims) bool {
	exp, ok := claims["exp"].(float64)
	if !ok {
		return true
	}
	return float64(time.Now().Unix()) > exp
}

func getSub(claims jwt.MapClaims) (uuid.UUID, error) {
	subClaim, ok := claims["sub"].(string)
	if !ok {
		return uuid.UUID{}, errors.New("sub extraction failed")
	}

	userID, err := uuid.Parse(subClaim)
	if err != nil {
		return uuid.UUID{}, errors.New("sub extraction failed")

	}

	return userID, nil
}

func handleUnauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	c.Abort()
}
