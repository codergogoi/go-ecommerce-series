package rest

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RestHandler struct {
	App *fiber.App
	DB  *gorm.DB
}
