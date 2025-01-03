package rest

import (
	"github.com/gofiber/fiber/v2"
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/pkg/payment"
	"gorm.io/gorm"
)

type RestHandler struct {
	App    *fiber.App
	DB     *gorm.DB
	Auth   helper.Auth
	Config config.AppConfig
	Pc     payment.PaymentClient
}
