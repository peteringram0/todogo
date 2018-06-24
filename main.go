package main

import (
	"database/sql"
	"todogo/auth"
	"todogo/handlers"
	"todogo/helper"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	"github.com/michaeljs1990/sqlitestore"
)

var store *sqlitestore.SqliteStore

func init() {
	var err error
	store, err = sqlitestore.NewSqliteStore("./sessions.db", "sessions", "/", 3600, []byte("1J9DG}m)i=ew(?wfLb0)"))
	if err != nil {
		panic(err)
	}
}

// todo.go
func main() {

	db := initDB("storage.db")
	migrate(db)

	// Instance of echo
	e := echo.New()

	// Setup sessions
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.Static("/public", "public")

	e.File("/", "public/index.html")

	v1 := e.Group("/api/v1")

	v1.GET("/auth", auth.AuthHandler(db))
	v1.GET("/login", auth.LoginHandler())

	tasks := v1.Group("/tasks")
	tasks.Use(middleware.JWT([]byte("secret")))
	tasks.GET("", handlers.GetTasks(db))
	tasks.POST("", handlers.PostTask(db))
	tasks.PUT("/:id", handlers.PutTask(db))
	// tasks.DELETE("/tasks/:id", func(c echo.Context) error { return c.JSON(200, "DELETE Task "+c.Param("id")) })

	me := v1.Group("/me")
	me.Use(middleware.JWT([]byte("secret")))
	me.GET("", handlers.GetMe(db))

	logout := v1.Group("/logout")
	logout.Use(middleware.JWT([]byte("secret")))
	logout.GET("", auth.LogoutHandler())

	// Start as a web server
	e.Start(":" + helper.GetEnv("PORT", "8000"))

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
				uid INTEGER NOT NULL,
        name VARCHAR NOT NULL
    );

		CREATE TABLE IF NOT EXISTS users(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        email VARCHAR NOT NULL,
				name VARCHAR NOT NULL,
				picture VARCHAR NOT NULL
    );
    `

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}
