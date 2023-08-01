package v1

import (
	"database/sql"
	"github.com/SenselessA/w2w_backend/internal/middleware"
	"github.com/SenselessA/w2w_backend/internal/models"
	"github.com/SenselessA/w2w_backend/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type RatingHandler struct {
	services *services.Services
}

func (h *RatingHandler) initRatingRouter(api *fiber.Router, middleware *middleware.Middleware) {
	router := *api
	rating := router.Group("/ratings")
	{
		rating.Post("/", middleware.AuthenticateJWT(), h.create)
		rating.Put("/", middleware.AuthenticateJWT(), h.update)
		rating.Delete("/", middleware.AuthenticateJWT(), h.delete)
		rating.Get("/:id", h.getByUserId)
		rating.Get("/ratedByUser/:id", middleware.AuthenticateJWT(), h.getRateByUserId)
	}
}

func (h *RatingHandler) create(c *fiber.Ctx) error {
	input := models.RatingCreateInput{}

	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	input.UserID = int64(claims["id"].(float64))

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	err := h.services.Ratings.Create(input)
	if err != nil {
		return err
	}

	c.Status(fiber.StatusOK)

	return nil
}

func (h *RatingHandler) update(c *fiber.Ctx) error {
	input := models.RatingUpdateInput{}

	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	input.UserID = int64(claims["id"].(float64))

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	updatedRating, err := h.services.Ratings.Update(input)
	if err != nil {
		return err
	}

	c.Status(fiber.StatusOK)
	err = c.JSON(updatedRating)

	if err != nil {
		return err
	}

	return nil
}

func (h *RatingHandler) delete(c *fiber.Ctx) error {
	input := models.RatingDeleteInput{}

	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	input.UserID = int64(claims["id"].(float64))

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	err := h.services.Ratings.Delete(input)
	if err != nil {
		return err
	}

	c.Status(fiber.StatusOK)

	return nil
}

func (h *RatingHandler) getRateByUserId(c *fiber.Ctx) error {
	input := models.RatingGetByUserIdInput{}

	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	input.UserID = int64(claims["id"].(float64))
	input.MovieID = c.Params("id")

	rating, err := h.services.Ratings.GetRateByUserId(input)
	if err != nil {
		if err == sql.ErrNoRows {
			c.Status(fiber.StatusOK)
			err = c.JSON(fiber.Map{"rating": nil})
			return nil
		}
		return err
	}

	c.Status(fiber.StatusOK)
	err = c.JSON(rating)

	if err != nil {
		return err
	}

	return nil
}

func (h *RatingHandler) getByUserId(c *fiber.Ctx) error {
	userId := c.Params("id")

	ratings, err := h.services.Ratings.GetByUserId(userId)
	if err != nil {
		return err
	}

	c.Status(fiber.StatusOK)
	err = c.JSON(ratings)

	if err != nil {
		return err
	}

	return nil
}
