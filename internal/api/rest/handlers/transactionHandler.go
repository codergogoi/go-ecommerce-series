package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/internal/service"
	"gorm.io/gorm"
)

type TransactionHandler struct {
	svc service.TransactionService
}

func initializeTransactionService(db *gorm.DB, auth helper.Auth) service.TransactionService {
	return service.TransactionService{
		Repo: repository.NewTransactionRepository(db),
		Auth: auth,
	}
}

func SetupTransactionRoutes(as *rest.RestHandler) {

	app := as.App
	svc := initializeTransactionService(as.DB, as.Auth)

	handler := TransactionHandler{
		svc: svc,
	}

	secRoute := app.Group("/", as.Auth.Authorize)
	secRoute.Get("/payment", handler.MakePayment)

	sellerRoute := app.Group("/seller", as.Auth.AuthorizeSeller)
	sellerRoute.Get("/orders", handler.GetOrders)
	sellerRoute.Get("/orders/:id", handler.GetOrderDetails)
}

func (h *TransactionHandler) MakePayment(ctx *fiber.Ctx) error {
	payload := struct {
		message string `json:"message"`
	}{
		message: "success",
	}
	return ctx.Status(200).JSON(payload)
}

func (h *TransactionHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("success")
}

func (h *TransactionHandler) GetOrderDetails(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("success")
}
