package main

import (
	"database/sql"

	"todogo/handlers"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

// todo.go
func main() {

	db := initDB("storage.db")
	migrate(db)

	// Instance of echo
	e := echo.New()

	e.File("/", "public/index.html")

	e.GET("/tasks", handlers.GetTasks(db))
	e.POST("/tasks", handlers.PostTask(db))
	e.PUT("/tasks/:id", handlers.PutTask(db))
	// e.DELETE("/tasks/:id", func(c echo.Context) error { return c.JSON(200, "DELETE Task "+c.Param("id")) })

	// Start as a web server
	e.Start(":8000")

}

// todo.go
func initDB(filepath string) *sql.DB {

	db, err := sql.Open("sqlite3", filepath)

	// Here we check for any db errors then exit
	if err != nil {
		panic(err)
	}

	// If we don't get any errors but somehow still don't get a db connection
	// we exit as well
	if db == nil {
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS tasks(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name VARCHAR NOT NULL
    );
    `

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}
