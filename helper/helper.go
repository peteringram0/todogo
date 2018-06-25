package helper

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// Get ENV key or fallback
func GetEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetUID(c echo.Context) int64 {

	claims := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)
	uid2 := claims["uid"].(float64)
	return int64(uid2)

}
