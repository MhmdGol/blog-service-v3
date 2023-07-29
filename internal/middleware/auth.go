package middleware

import (
	"blog-service-v3/internal/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func RequireAuth(ctx *fiber.Ctx) error {
	tokenString := ctx.GetReqHeaders()["Authorization"]
	//there is another way to do this

	conf, err := config.Load()
	if err != nil {
		return err
	}

	if tokenString == "" {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return ([]byte)(conf.SecretKey), nil
	})

	if err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	if !token.Valid {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	return ctx.Next()
}
