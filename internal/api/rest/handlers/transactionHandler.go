package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/internal/service"
	"go-ecommerce-app/pkg/payment"
	"gorm.io/gorm"
	"net/http"
)

type TransactionHandler struct {
	Svc           service.TransactionService
	UserSvc       service.UserService
	PaymentClient payment.PaymentClient
	Config        config.AppConfig
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
	useSvc := service.UserService{
		Repo:   repository.NewUserRepository(as.DB),
		CRepo:  repository.NewCatalogRepository(as.DB),
		Auth:   as.Auth,
		Config: as.Config,
	}

	handler := TransactionHandler{
		Svc:           svc,
		PaymentClient: as.Pc,
		UserSvc:       useSvc,
		Config:        as.Config,
	}

	secRoute := app.Group("/buyer", as.Auth.Authorize)
	secRoute.Get("/payment", handler.MakePayment)
	secRoute.Get("/verify", handler.VerifyPayment)

	sellerRoute := app.Group("/seller", as.Auth.AuthorizeSeller)
	sellerRoute.Get("/orders", handler.GetOrders)
	sellerRoute.Get("/orders/:id", handler.GetOrderDetails)
}

func (h *TransactionHandler) MakePayment(ctx *fiber.Ctx) error {

	//1 grab authorized user
	user := h.Svc.Auth.GetCurrentUser(ctx)

	pubKey := h.Config.PubKey

	// 2. Check if user has an active payment session then return the payment url
	activePayment, err := h.Svc.GetActivePayment(user.ID)
	if activePayment.ID > 0 {
		return ctx.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "create payment",
			"pubKey":  pubKey,
			"secret":  activePayment.ClientSecret,
		})
	}

	//3. call user service get cart data to aggregate the total amount and collect payment
	_, amount, err := h.UserSvc.FindCart(user.ID)

	orderId, err := helper.RandomNumbers(8)
	if err != nil {
		return rest.InternalError(ctx, errors.New("error generating order id"))
	}

	// 4. Create a new payment session on stripe
	paymentResult, err := h.PaymentClient.CreatePayment(amount, user.ID, orderId)

	//5. Store payment session in db to create to store payment info
	err = h.Svc.StoreCreatedPayment(dto.CreatePaymentRequest{
		UserId:       user.ID,
		Amount:       amount,
		ClientSecret: paymentResult.ClientSecret,
		PaymentId:    paymentResult.ID,
		OrderId:      orderId,
	})

	if err != nil {
		return ctx.Status(400).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "create payment",
		"pubKey":  pubKey,
		"secret":  paymentResult.ClientSecret,
	})
}

func (h *TransactionHandler) VerifyPayment(ctx *fiber.Ctx) error {

	// grab authorized user
	user := h.Svc.Auth.GetCurrentUser(ctx)

	// do we have active payment session to verify?
	activePayment, err := h.Svc.GetActivePayment(user.ID)
	if err != nil || activePayment.ID == 0 {
		return ctx.Status(400).JSON(errors.New("no active payment exist"))
	}

	// fetch payment status from stripe
	paymentRes, err := h.PaymentClient.GetPaymentStatus(activePayment.PaymentId)
	paymentJson, _ := json.Marshal(paymentRes)
	paymentLogs := string(paymentJson)
	paymentStatus := "failed"

	// if payment then create order
	if paymentRes.Status == "succeeded" {
		// create Order
		paymentStatus = "success"
		err = h.UserSvc.CreateOrder(user.ID, activePayment.OrderId, activePayment.PaymentId, activePayment.Amount)
	}

	if err != nil {
		return rest.InternalError(ctx, err)
	}

	// update payment status
	h.Svc.UpdatePayment(user.ID, paymentStatus, paymentLogs)

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message":  "create payment",
		"response": paymentRes,
	})
}

func (h *TransactionHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("success")
}

func (h *TransactionHandler) GetOrderDetails(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("success")
}
