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

func GetUID(c echo.Context) (int64, error) {

	sess, _ := session.Get("session", c)
	uid := sess.Values["user-uid"].(int64)

	// @TODO check the type, if string then convert it, if int64 carry on

	// log.Printf("We are using the id NEWWW %f", uid)

	return uid, nil

	// if missing != nil {
	// 	return 0, errors.New("No session info !!")
	// }
	//
	// uidString, err := strconv.ParseInt(uid.(string), 10, 64)
	//
	// if err == nil {
	// 	return uidString, nil
	// } else {
	// 	return 0, errors.New("UID error")
	// }

}
