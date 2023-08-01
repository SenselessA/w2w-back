package middleware

import (
	"github.com/SenselessA/w2w_backend/internal/config"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

type Middleware struct {
	secret config.SecretConfig
}

func InitMiddlewares(cfg config.SecretConfig) *Middleware {
	return &Middleware{secret: cfg}
}

func (m Middleware) AuthenticateJWT() func(*fiber.Ctx) error {
	jwtSecret := m.secret.Jwt

	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(jwtSecret)},
		SuccessHandler: func(ctx *fiber.Ctx) error {
			// user := ctx.Locals("user").(*jwt.Token)
			// claims := user.Claims.(jwt.MapClaims)
			// name := claims["email"].(string)
			return ctx.Next()
		},
	})
}
