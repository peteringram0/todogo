package handlers

import (
	"database/sql"
	"net/http"
	"todogo/models"
	"todogo/structs"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

func GetMe(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		sess, _ := session.Get("session", c)

		tasks := models.GetTasks(db, c)

		return c.JSON(http.StatusOK, structs.User{
			Name:    sess.Values["user-name"].(string),
			Picture: sess.Values["user-picture"].(string) + "?sz=100", // Image size from google
			Email:   sess.Values["user-email"].(string),
			Tasks:   tasks,
		})

	}
}
