package v1

import (
	"github.com/SenselessA/w2w_backend/internal/middleware"
	"github.com/SenselessA/w2w_backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	services         *services.Services
	UsersHandler     *UsersHandler
	FavoritesHandler *FavoritesHandler
	RatingHandler    *RatingHandler
	MoviesHandler    *MoviesHandler
	ProfileHandler   *ProfileHandler
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Init(api *fiber.Router, middleware *middleware.Middleware) {
	router := *api

	AuthHandler := &UsersHandler{services: h.services}
	FavoritesHandler := &FavoritesHandler{services: h.services}
	RatingHandler := &RatingHandler{services: h.services}
	MoviesHandler := &MoviesHandler{services: h.services}
	ProfileHandler := &ProfileHandler{services: h.services}

	v1 := router.Group("/v1")
	{
		AuthHandler.initAuthRouter(&v1, middleware)
		FavoritesHandler.initFavoritesRouter(&v1, middleware)
		RatingHandler.initRatingRouter(&v1, middleware)
		MoviesHandler.initMoviesRouter(&v1, middleware)
		ProfileHandler.initProfileRouter(&v1, middleware)
	}
}
