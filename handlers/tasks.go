package handlers

import (
	"database/sql"
	"net/http"
	_ "strconv"

	"todogo/models"
	"todogo/structs"

	"github.com/labstack/echo"
)

type H map[string]interface{}

// GetTasks endpoint
func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

func PostTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Instantiate a new task
		var task structs.Task

		// Map imcoming JSON body to the new Task
		c.Bind(&task)

		// If name is empty return bad request
		if task.Name == "" {
			return c.JSON(http.StatusBadRequest, H{
				"error": "Please send us a name!!",
			})
		}

		// Add a task using our new model
		id, err := models.PostTask(db, task.Name)

		// Return a JSON response if successful
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
			// Handle any errors
		} else {
			return err
		}

	}
}

func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Instantiate a new task
		var task structs.Task

		// Map imcoming JSON body to the new Task
		c.Bind(&task)

		// If name is empty return bad request
		if task.Name == "" {
			return c.JSON(http.StatusBadRequest, H{
				"error": "Please send us a name!!",
			})
		}

		// Add a task using our new model
		rowsAffected, err := models.PutTask(db, c.Param("id"), task.Name)

		if rowsAffected == 0 {
			return c.JSON(http.StatusBadRequest, H{
				"error": "Cant updated that !!",
			})
		} else if rowsAffected >= 1 && err == nil {
			return c.JSON(http.StatusCreated, H{
				"updated": c.Param("id"),
			})
		} else {
			return err
		}

	}
}
