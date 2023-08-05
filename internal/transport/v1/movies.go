package v1

import (
	"github.com/SenselessA/w2w_backend/internal/middleware"
	"github.com/SenselessA/w2w_backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

type MoviesHandler struct {
	services *services.Services
}

func (h *MoviesHandler) initMoviesRouter(api *fiber.Router, middleware *middleware.Middleware) {
	router := *api
	rating := router.Group("/movies")
	{
		rating.Get("/:id", h.getByMovieId)
	}
}

func (h *MoviesHandler) getByMovieId(c *fiber.Ctx) error {
	movieId := c.Params("id")

	movie, err := h.services.Movies.GetByMovieId(movieId)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		err := c.SendString("Ошибка сервера")
		if err != nil {
			return err
		}
		return err
	}

	c.Status(fiber.StatusOK)
	err = c.JSON(movie)

	if err != nil {
		return err
	}

	return nil
}
