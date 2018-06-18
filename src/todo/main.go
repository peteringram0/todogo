package main

import (
    "github.com/labstack/echo"
)

// todo.go
func main() {

    // Instance of echo
    e := echo.New()

    e.GET("/tasks", func(c echo.Context) error { return c.JSON(200, "GET Tasks") })
    e.PUT("/tasks", func(c echo.Context) error { return c.JSON(200, "PUT Tasks") })
    e.DELETE("/tasks/:id", func(c echo.Context) error { return c.JSON(200, "DELETE Task "+c.Param("id")) })

    // Start as a web server
    e.Start(":8000")

}
