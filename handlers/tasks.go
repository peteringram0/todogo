package handlers

import (
	"database/sql"
	"net/http"

	"todogo/helper"
	"todogo/models"
	"todogo/structs"

	"github.com/labstack/echo"
)

type H map[string]interface{}

func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		tasks := models.GetTasks(db, c)
		return c.JSON(http.StatusOK, tasks)
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

		uid := helper.GetUID(c)

		// Add a task using our new model
		id, err := models.PostTask(db, task.Name, uid)

		// Return a JSON response if successful
		if err == nil {
			return c.JSON(http.StatusCreated, structs.Task{
				ID:   id,
				UID:  uid,
				Name: task.Name,
			})
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
			return c.JSON(http.StatusOK, H{
				"updated": c.Param("id"),
			})
		} else {
			return err
		}

	}
}
