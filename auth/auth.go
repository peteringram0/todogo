package auth

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"

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

		session, err := session.Get("session", c)

		if err != nil {
			panic(err)
		}

		// session.Options = &sessions.Options{
		// Path:     "/",
		// MaxAge:   86400 * 7,
		// HttpOnly: true,
		// }
		session.Values["state"] = state
		session.Save(c.Request(), c.Response())

		return c.Render(http.StatusOK, "auth.html", map[string]interface{}{
			"link": link,
		})

	}
}

func AuthHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		// Handle the exchange code to initiate a transport.
		// session := sessions.Default(c)
		session, _ := session.Get("session", c)
		valWithOutType := session.Values["state"]
		fmt.Printf("valWithOutType: %s\n", valWithOutType)

		// queryState := c.Request.URL.Query().Get("state")
		// if retrievedState != queryState {
		// log.Printf("Invalid session state: retrieved: %s; Param: %s", retrievedState, queryState)
		// c.HTML(http.StatusUnauthorized, "error.tmpl", gin.H{"message": "Invalid session state."})
		// return
		// }
		// code := c.Request.URL.Query().Get("code")
		// tok, err := conf.Exchange(oauth2.NoContext, code)
		// if err != nil {
		// 	log.Println(err)
		// 	c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"message": "Login failed. Please try again."})
		// 	return
		// }
		//
		// client := conf.Client(oauth2.NoContext, tok)
		// userinfo, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
		// if err != nil {
		// 	log.Println(err)
		// 	c.AbortWithStatus(http.StatusBadRequest)
		// 	return
		// }
		// defer userinfo.Body.Close()
		// data, _ := ioutil.ReadAll(userinfo.Body)
		// u := structs.User{}
		// if err = json.Unmarshal(data, &u); err != nil {
		// 	log.Println(err)
		// 	c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"message": "Error marshalling response. Please try agian."})
		// 	return
		// }
		// session.Set("user-id", u.Email)
		// err = session.Save()
		// if err != nil {
		// 	log.Println(err)
		// 	c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"message": "Error while saving session. Please try again."})
		// 	return
		// }
		// seen := false
		// db := database.MongoDBConnection{}
		// if _, mongoErr := db.LoadUser(u.Email); mongoErr == nil {
		// 	seen = true
		// } else {
		// 	err = db.SaveUser(&u)
		// 	if err != nil {
		// 		log.Println(err)
		// 		c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"message": "Error while saving user. Please try again."})
		// 		return
		// 	}
		// }
		// c.HTML(http.StatusOK, "battle.tmpl", gin.H{"email": u.Email, "seen": seen})

		return c.JSON(http.StatusCreated, H{
			"message": "AuthHandler",
		})

	}
}
