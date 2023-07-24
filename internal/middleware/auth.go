package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var secretKey = []byte("Blogloglog")

// type User struct {
// 	ID       int
// 	Username string `json:"user"`
// 	Password string `json:"pass"`
// }

// func LoginHandler(w http.ResponseWriter, r *http.Request) {
// 	// ------ Hardcoded user ------
// 	allowedUser := User{
// 		ID:       1,
// 		Username: "Mhmd",
// 		Password: "1234",
// 	}
// 	// ----------------------------

// 	// login check
// 	reqBody, _ := ioutil.ReadAll(r.Body)
// 	var user User
// 	json.Unmarshal(reqBody, &user)

// 	if user.Username != allowedUser.Username || user.Password != allowedUser.Password {
// 		fmt.Fprintf(w, "Not Authenticated!")
// 		return
// 	}

// 	// here we generate a token for the user
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"user_id":   user.ID,
// 		"username":  user.Username,
// 		"expiresAt": time.Now().Add(time.Hour * 24).Unix(),
// 	})

// 	tokenString, err := token.SignedString(secretKey)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// give the token to the user
// 	w.Write([]byte(tokenString))
// }

func RequireAuth(ctx *fiber.Ctx) error {
	tokenString := ctx.GetReqHeaders()["Authorization"]
	//there is another way to do this

	if tokenString == "" {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	if !token.Valid {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	return ctx.Next()
}
