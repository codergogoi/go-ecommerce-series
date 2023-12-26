package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-ecommerce-app/internal/api/rest"
	"net/http"
)

type UserHandler struct {
	// svc UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {

	app := rh.App

	// create an instance of user service & inject to handler
	handler := UserHandler{}

	// Public endpoints
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	// Private endpoint
	app.Get("/verify", handler.GetVerificationCode)
	app.Post("/verify", handler.Verify)
	app.Post("/profile", handler.CreateProfile)
	app.Get("/profile", handler.GetProfile)

	app.Post("/cart", handler.AddToCart)
	app.Get("/cart", handler.GetCart)
	app.Get("/order", handler.GetOrders)
	app.Get("/order/:id", handler.GetOrder)

	app.Post("/become-seller", handler.BecomeSeller)

}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "register",
	})
}
func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login",
	})
}
func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get verification code",
	})
}
func (h *UserHandler) Verify(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "verify",
	})
}
func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "create profile",
	})
}
func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get profile",
	})
}
func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "add to cart",
	})
}
func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get cart",
	})
}
func (h *UserHandler) CreateOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "create order",
	})
}
func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get orders",
	})
}
func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get order by id",
	})
}
func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "become a seller",
	})
}
