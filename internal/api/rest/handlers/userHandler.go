package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/internal/service"
	"log"
	"net/http"
	"strconv"
)

type UserHandler struct {
	// svc UserService
	svc service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {

	app := rh.App

	// create an instance of user service & inject to handler
	svc := service.UserService{
		Repo:   repository.NewUserRepository(rh.DB),
		CRepo:  repository.NewCatalogRepository(rh.DB),
		Auth:   rh.Auth,
		Config: rh.Config,
	}
	handler := UserHandler{
		svc: svc,
	}

	pubRoutes := app.Group("/")
	// Public endpoints
	pubRoutes.Post("/register", handler.Register)
	pubRoutes.Post("/login", handler.Login)

	pvtRoutes := pubRoutes.Group("/users", rh.Auth.Authorize)

	// Private endpoint
	pvtRoutes.Get("/verify", handler.GetVerificationCode)
	pvtRoutes.Post("/verify", handler.Verify)

	pvtRoutes.Post("/profile", handler.CreateProfile)
	pvtRoutes.Get("/profile", handler.GetProfile)
	pvtRoutes.Patch("/profile", handler.UpdateProfile)

	pvtRoutes.Post("/cart", handler.AddToCart)
	pvtRoutes.Get("/cart", handler.GetCart)

	pvtRoutes.Get("/order", handler.GetOrders)
	pvtRoutes.Get("/order/:id", handler.GetOrder)

	pvtRoutes.Post("/become-seller", handler.BecomeSeller)

}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	user := dto.UserSignup{}
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide valid inputs",
		})
	}

	token, err := h.svc.Signup(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "error on signup",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "register",
		"token":   token,
	})
}
func (h *UserHandler) Login(ctx *fiber.Ctx) error {

	loginInput := dto.UserLogin{}
	err := ctx.BodyParser(&loginInput)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide valid inputs",
		})
	}

	token, err := h.svc.Login(loginInput.Email, loginInput.Password)

	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "please provide correct user id password",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login",
		"token":   token,
	})
}
func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {

	user := h.svc.Auth.GetCurrentUser(ctx)
	log.Println(user)
	// create verification code and update to user profile in DB
	err := h.svc.GetVerificationCode(user)
	log.Println(err)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "unable to generate verification code",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get verification code",
	})

}
func (h *UserHandler) Verify(ctx *fiber.Ctx) error {

	user := h.svc.Auth.GetCurrentUser(ctx)

	// request
	var req dto.VerificationCodeInput

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide a valid input",
		})
	}

	err := h.svc.VerifyCode(user.ID, req.Code)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "verified successfully",
	})
}
func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {

	user := h.svc.Auth.GetCurrentUser(ctx)
	req := dto.ProfileInput{}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide a valid input",
		})
	}
	log.Printf("User %v", user)
	// create profile

	err := h.svc.CreateProfile(user.ID, req)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "unable to create profile",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "profile created successfully",
	})
}
func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {

	user := h.svc.Auth.GetCurrentUser(ctx)
	log.Println(user)

	// call user service and perform get profile
	profile, err := h.svc.GetProfile(user.ID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "unable to get profile",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get profile",
		"profile": profile,
	})
}

func (h *UserHandler) UpdateProfile(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)
	req := dto.ProfileInput{}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide a valid input",
		})
	}

	err := h.svc.UpdateProfile(user.ID, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "unable to update profile",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "profile updated successfully",
	})
}

func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {

	req := dto.CreateCartRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide a valid product and qty",
		})
	}

	user := h.svc.Auth.GetCurrentUser(ctx)

	// call user service and perform create cart
	cartItems, err := h.svc.CreateCart(req, user)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "cart created successfully", cartItems)

}
func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)
	cart, _, err := h.svc.FindCart(user.ID)
	if err != nil {
		return rest.InternalError(ctx, errors.New("cart does not exist"))
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get cart",
		"cart":    cart,
	})
}

func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)

	orders, err := h.svc.GetOrders(user)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get orders",
		"orders":  orders,
	})
}
func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	orderId, _ := strconv.Atoi(ctx.Params("id"))
	user := h.svc.Auth.GetCurrentUser(ctx)

	order, err := h.svc.GetOrderById(uint(orderId), user.ID)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get order by id",
		"order":   order,
	})
}
func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {

	user := h.svc.Auth.GetCurrentUser(ctx)

	req := dto.SellerInput{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"message": "request parameters are not valid",
		})
	}

	token, err := h.svc.BecomeSeller(user.ID, req)

	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "fail to become seller",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "become seller",
		"token":   token,
	})
}
