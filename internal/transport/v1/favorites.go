package v1

import (
	"database/sql"
	"github.com/SenselessA/w2w_backend/internal/middleware"
	"github.com/SenselessA/w2w_backend/internal/models"
	"github.com/SenselessA/w2w_backend/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type FavoritesHandler struct {
	services *services.Services
}

func (h *FavoritesHandler) initFavoritesRouter(api *fiber.Router, middleware *middleware.Middleware) {
	router := *api
	favorites := router.Group("/favorites")
	{
		favorites.Post("/", middleware.AuthenticateJWT(), h.create)
		favorites.Delete("/:id", middleware.AuthenticateJWT(), h.delete)
		favorites.Get("/profile/:id", h.getByUserId)
		favorites.Get("/:id", middleware.AuthenticateJWT(), h.getFavoriteByUserId)
	}
}

func (h *FavoritesHandler) create(c *fiber.Ctx) error {
	input := models.FavoritesCreateInput{}

	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	input.UserID = int64(claims["id"].(float64))

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	err := h.services.Favorites.Create(input)
	if err != nil {
		return err
	}

	c.Status(fiber.StatusOK)

	return nil
}

func (h *FavoritesHandler) delete(c *fiber.Ctx) error {
	input := models.FavoritesCreateInput{}
	input.MovieID = c.Params("id")

	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	input.UserID = int64(claims["id"].(float64))

	err := h.services.Favorites.Delete(input)
	if err != nil {
		return err
	}

	c.Status(fiber.StatusOK)

	return nil
}

func (h *FavoritesHandler) getByUserId(c *fiber.Ctx) error {
	userId := c.Params("id")

	favorites, err := h.services.Favorites.GetByUserId(userId)
	if err != nil {
		return err
	}

	c.Status(fiber.StatusOK)
	err = c.JSON(favorites)

	if err != nil {
		return err
	}

	return nil
}

func (h *FavoritesHandler) getFavoriteByUserId(c *fiber.Ctx) error {
	input := models.FavoriteByUserIdInput{}
	input.MovieID = c.Params("id")

	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	input.UserID = int64(claims["id"].(float64))

	_, err := h.services.Favorites.GetFavoriteByUserId(input)
	if err != nil {
		if err == sql.ErrNoRows {
			c.Status(fiber.StatusOK)
			err = c.JSON(fiber.Map{"isFavorite": false})
			return nil
		}
		return err
	}

	c.Status(fiber.StatusOK)
	err = c.JSON(fiber.Map{"isFavorite": true})

	if err != nil {
		return err
	}

	return nil
}
