package auth

import (
	"crypto/rand"
	"log"
	"net/http"

	"github.com/labstack/echo"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"os"
)

type H map[string]interface{}

var cred Credentials
var conf *oauth2.Config

type Credentials struct {
	Cid     string `json:"cid"`
	Csecret string `json:"csecret"`
}

func getLoginURL(state string) string {
	return conf.AuthCodeURL(state)
}

// RandToken generates a random @l length token.
func RandToken(l int) string {
	b := make([]byte, l)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func init() {
	file, err := ioutil.ReadFile("./creds.json")
	if err != nil {
		log.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	json.Unmarshal(file, &cred)

	conf = &oauth2.Config{
		ClientID:     cred.Cid,
		ClientSecret: cred.Csecret,
		RedirectURL:  "http://127.0.0.1:8000/auth",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email", // https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		},
		Endpoint: google.Endpoint,
	}

}

func LoginHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		state := RandToken(32)
		link := getLoginURL(state)

		session, _ := session.Get("session", c)
		session.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		session.Values["state"] = state
		session.Save(c.Request(), c.Response())

		return c.Render(http.StatusOK, "auth.html", map[string]interface{}{
			"link": link,
		})

	}
}

func AuthHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.JSON(http.StatusCreated, H{
			"message": "AuthHandler",
		})

	}
}
