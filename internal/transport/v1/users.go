package v1

import (
	"github.com/SenselessA/w2w_backend/internal/middleware"
	"github.com/SenselessA/w2w_backend/internal/models"
	"github.com/SenselessA/w2w_backend/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type UsersHandler struct {
	services *services.Services
}

func (h *UsersHandler) initAuthRouter(api *fiber.Router, middleware *middleware.Middleware) {
	router := *api
	users := router.Group("/auth")
	{
		//users.Get("/", h.get)
		//users.Delete("/:id", h.delete)
		//users.Put("/:id", h.update)
		users.Post("/register", h.create)
		users.Post("/login", h.login)
	}
}

// @Summary create
// @Tags products
// @ID create-product
// @Param input body types.CreateInput true " "
// @Success 200 {integer} integer 1
// @Failure 400 {object} types.ErrorResponse
// @Router /api/v1/products/ [post]
func (h *UsersHandler) create(c *fiber.Ctx) error {
	input := models.UserCreateInput{}

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	userId, err := h.services.Users.Create(input)
	if err != nil {
		return err
	}

	claims := jwt.MapClaims{
		"id":    userId,
		"email": input.Email,
		"admin": false,
		"exp":   time.Now().Add(time.Hour * 300).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("salt"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// http response
	return c.JSON(fiber.Map{"token": t})
}

func (h *UsersHandler) login(c *fiber.Ctx) error {
	input := models.UserLoginInput{}

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	user, err := h.services.Users.Login(input)
	if err != nil {
		return err
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
		"name":  user.Username,
		"admin": false,
		"exp":   time.Now().Add(time.Hour * 300).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("salt"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

// @Summary get csv
// @Tags products
// @ID get-csv-products
// @Success 200 "OK"
// @Failure 400 {object} types.ErrorResponse
// @Router /api/v1/products/ [get]
//func (h *Handler) get(c *gin.Context) {
//	products, err := h.services.Products.Get(c)
//	if err != nil {
//		newErrorResponse(c, http.StatusBadRequest, err)
//		return
//	}
//
//	headers := map[string]string{
//		"Content-Disposition": `attachment; filename="products.csv"`,
//	}
//
//	c.DataFromReader(http.StatusOK, -1, "text/html; charset=UTF-8", products, headers)
//}

// @Summary delete
// @Tags products
// @ID delete-product
// @Param id path integer true " "
// @Success 200 "OK"
// @Failure 400 {object} types.ErrorResponse
// @Router /api/v1/products/{id} [delete]
//func (h *Handler) delete(c *gin.Context) {
//	id, ok := c.Params.Get("id")
//	if !ok {
//		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("empty id"))
//		return
//	}
//
//	i, err := strconv.Atoi(id)
//	if err != nil {
//		newErrorResponse(c, http.StatusBadRequest, err)
//		return
//	}
//
//	if err = h.services.Products.Delete(c, int64(i)); err != nil {
//		newErrorResponse(c, http.StatusBadRequest, err)
//		return
//	}
//
//	c.Status(http.StatusOK)
//}

// @Summary update
// @Tags products
// @ID update-product
// @Param id path integer true " "
// @Param input body types.UpdateInput true " "
// @Success 200 "OK"
// @Failure 400 {object} types.ErrorResponse
// @Router /api/v1/products/{id} [put]
//func (h *Handler) update(c *gin.Context) {
//	id, ok := c.Params.Get("id")
//	if !ok {
//		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("empty id"))
//		return
//	}
//
//	i, err := strconv.Atoi(id)
//	if err != nil {
//		newErrorResponse(c, http.StatusBadRequest, err)
//		return
//	}
//
//	var input types.UpdateInput
//
//	if err = c.BindJSON(&input); err != nil {
//		newErrorResponse(c, http.StatusBadRequest, err)
//		return
//	}
//
//	if err = h.services.Products.Update(c, int64(i), input); err != nil {
//		newErrorResponse(c, http.StatusBadRequest, err)
//		return
//	}
//
//	c.Status(http.StatusOK)
//}
