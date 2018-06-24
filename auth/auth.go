package auth

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"todogo/helper"
	"todogo/models"
	"todogo/structs"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type H map[string]interface{}

var cred structs.Credentials
var conf *oauth2.Config

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
		RedirectURL:  helper.GetEnv("REDIRECT_URL", "http://127.0.0.1:8000/api/v1/auth"),
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

		sess, _ := session.Get("session", c)
		sess.Values["state"] = state
		sess.Save(c.Request(), c.Response())

		return c.JSON(http.StatusOK, H{
			"link": link,
		})

	}
}

func LogoutHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, H{
			"logout": "logout trigger",
		})
	}
}

func AuthHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Get the session from session - WORKING
		sess, _ := session.Get("session", c)

		retrievedState := sess.Values["state"]
		queryState := c.QueryParam("state")

		// If our current state is different to the receved state !!
		if retrievedState != queryState {
			log.Printf("Invalid session state: retrieved: %s; Param: %s", retrievedState, queryState)
			return c.JSON(http.StatusUnauthorized, H{
				"error": "Invalid session state",
			})
		}

		code := c.QueryParam("code")
		tok, err := conf.Exchange(oauth2.NoContext, code)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusUnauthorized, H{
				"error": "Login failed",
			})
		}

		client := conf.Client(oauth2.NoContext, tok)
		userinfo, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
		if err != nil {
			log.Println(err)
			return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
		}

		defer userinfo.Body.Close()
		data, _ := ioutil.ReadAll(userinfo.Body)
		u := structs.User{}
		if err = json.Unmarshal(data, &u); err != nil {
			log.Println(err)
			return c.JSON(http.StatusUnauthorized, H{
				"error": "Unmarshal failed",
			})
		}

		sess.Values["user-email"] = u.Email
		sess.Values["user-name"] = u.Name
		sess.Values["user-picture"] = u.Picture
		sess.Save(c.Request(), c.Response())

		_, err = models.GetUser(db, u.Email)

		if err != nil {
			models.CreateUser(db, u)
		}

		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["email"] = u.Email
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		return c.Redirect(http.StatusSeeOther, "/?token="+t)

	}
}
