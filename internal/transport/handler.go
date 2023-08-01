package transport

import (
	"github.com/SenselessA/w2w_backend/internal/middleware"
	"github.com/SenselessA/w2w_backend/internal/services"
	v1 "github.com/SenselessA/w2w_backend/internal/transport/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/golang-jwt/jwt/v5"
)

type Handler struct {
	services    *services.Services
	middlewares *middleware.Middleware
}

func NewHandler(services *services.Services, middleware *middleware.Middleware) *Handler {
	return &Handler{
		services:    services,
		middlewares: middleware,
	}
}

func (h *Handler) Init() *fiber.App {
	app := fiber.New()

	// swagger initialization

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
	}))

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong 👋!")
	})

	h.initAPI(app)

	return app
}

func (h *Handler) initAPI(app *fiber.App) {
	handlerV1 := v1.NewHandler(h.services)
	api := app.Group("/api")

	handlerV1.Init(&api, h.middlewares)

	// Restricted Routes
	app.Get("/protected", h.middlewares.AuthenticateJWT(), restricted)
}

func restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["email"].(string)
	return c.SendString("Welcome " + name)
}
