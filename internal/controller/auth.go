package controller

import (
	"blog-service-v3/internal/dto"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	secretKey []byte
}

func NewAuthController(router fiber.Router, sk []byte) *AuthController {
	ctrl := AuthController{secretKey: sk}

	router.Post("/login", ctrl.Login)

	return &ctrl
}

func (ac *AuthController) Login(ctx *fiber.Ctx) error {
	// ------ Hardcoded user ------
	allowedUser := dto.User{
		ID:       1,
		Username: "Mhmd",
		Password: "1234",
	}
	// ----------------------------

	var user dto.User
	if err := ctx.BodyParser(&user); err != nil {
		return err
	}

	if user.Username != allowedUser.Username || user.Password != allowedUser.Password {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	// can go to the service layer {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"username":  user.Username,
		"expiresAt": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(ac.secretKey)
	if err != nil {
		return err
	}
	// }

	result := dto.AuthToken{
		Token: tokenString,
	}

	return ctx.JSON(result)
}
