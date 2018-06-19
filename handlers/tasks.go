package handlers

import (
    "fmt"
    "database/sql"
    "net/http"
    _ "strconv"
    "github.com/labstack/echo"
    "../models"
)

type H map[string]interface{}

// GetTasks endpoint
func GetTasks(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.JSON(http.StatusOK, models.GetTasks(db))
    }
}


// // Handler
// func(c echo.Context) (err error) {
//   u := new(User)
//   if err = c.Bind(u); err != nil {
//     return
//   }
//   return c.JSON(http.StatusOK, u)
// }

func PutTasks(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {

      // Instantiate a new task
      var task models.Task

      // Map imcoming JSON body to the new Task
      c.Bind(&task)

      // @TODO Here, if task.Name is missing or empty error

      // Add a task using our new model
      id, err := models.PutTasks(db, task.Name)

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
