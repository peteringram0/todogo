package main

import (
	"database/sql"
	"html/template"
	"io"

	"todogo/handlers"

	"os"

	"todogo/auth"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	_ "github.com/mattn/go-sqlite3"
	"github.com/michaeljs1990/sqlitestore"
)

var store *sqlitestore.SqliteStore

func init() {
	var err error
	store, err = sqlitestore.NewSqliteStore("./database", "sessions", "/", 3600, []byte("<SecretKey>"))
	if err != nil {
		panic(err)
	}
}

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

// todo.go
func main() {

	db := initDB("storage.db")
	migrate(db)

	// Instance of echo
	e := echo.New()

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = renderer

	// Setup sessions
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.File("/", "public/index.html")

	e.GET("/tasks", handlers.GetTasks(db))
	e.POST("/tasks", handlers.PostTask(db))
	e.PUT("/tasks/:id", handlers.PutTask(db))
	// e.DELETE("/tasks/:id", func(c echo.Context) error { return c.JSON(200, "DELETE Task "+c.Param("id")) })

	e.GET("/auth", auth.AuthHandler())
	e.GET("/login", auth.LoginHandler())

	port := os.Getenv("PORT")

	// Start as a web server
	e.Start(":" + port)

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
