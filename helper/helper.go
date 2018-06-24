package helper

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

// Get ENV key or fallback
func GetEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetUID(c echo.Context) int64 {

	sess, _ := session.Get("session", c)
	uid := sess.Values["user-uid"].(int64)

	return uid

}
