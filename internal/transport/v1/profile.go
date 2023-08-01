package v1

import (
	"github.com/SenselessA/w2w_backend/internal/middleware"
	"github.com/SenselessA/w2w_backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

type ProfileHandler struct {
	services *services.Services
}

func (h *ProfileHandler) initProfileRouter(api *fiber.Router, middleware *middleware.Middleware) {
	router := *api
	rating := router.Group("/profiles")
	{
		rating.Get("/:id", h.getByProfileId)
	}
}

func (h *ProfileHandler) getByProfileId(c *fiber.Ctx) error {
	userId := c.Params("id")

	movies, err := h.services.Ratings.GetByUserId(userId)
	if err != nil {
		return err
	}

	favorites, err := h.services.Favorites.GetByUserId(userId)
	if err != nil {
		return err
	}

	user, err := h.services.Users.GetUserById(userId)
	if err != nil {
		return err
	}

	c.Status(fiber.StatusOK)
	err = c.JSON(fiber.Map{"movies": movies, "favorites": favorites, "user": user})
	if err != nil {
		return err
	}

	return nil
}
