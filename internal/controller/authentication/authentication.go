package authentication

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("Blogloglog")

type User struct {
	ID       int
	Username string `json:"user"`
	Password string `json:"pass"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// ------ Hardcoded user ------
	allowedUser := User{
		ID:       1,
		Username: "Mhmd",
		Password: "1234",
	}
	// ----------------------------

	// login check
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)

	if user.Username != allowedUser.Username || user.Password != allowedUser.Password {
		fmt.Fprintf(w, "Not Authenticated!")
		return
	}

	// here we generate a token for the user
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"username":  user.Username,
		"expiresAt": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		panic(err)
	}

	// give the token to the user
	w.Write([]byte(tokenString))
}

func CheckAuthentication(r *http.Request) bool {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return false
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return false
	}

	if !token.Valid {
		return false
	}

	return true
}
