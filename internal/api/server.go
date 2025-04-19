package api

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/api/rest/handlers"
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/pkg/payment"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("database connection error %v\n", err)
	}

	log.Println("database connected")

	// run migration
	err = db.AutoMigrate(
		&domain.User{},
		&domain.Address{},
		&domain.BankAccount{},
		&domain.Category{},
		&domain.Product{},
		&domain.Cart{},
		&domain.Order{},
		&domain.OrderItem{},
		&domain.Payment{},
	)
	if err != nil {
		log.Fatalf("error on runing migration %v", err.Error())
	}

	log.Println("migration was successful")

	// cors configuration
	c := cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE, OPTIONS",
	})

	app.Use(c)

	app.Get("/", func(c *fiber.Ctx) error {
		return rest.SuccessResponse(c, "I am Healthy", &fiber.Map{
			"status": "ok with 200 status code",
		})
	})

	auth := helper.SetupAuth(config.AppSecret)

	paymentClient := payment.NewPaymentClient(config.StripeSecret)

	rh := &rest.RestHandler{
		App:    app,
		DB:     db,
		Auth:   auth,
		Config: config,
		Pc:     paymentClient,
	}

	setupRoutes(rh)

	app.Listen(config.ServerPort)

}

func setupRoutes(rh *rest.RestHandler) {
	// catalog
	handlers.SetupCatalogRoutes(rh)
	// user handler
	handlers.SetupUserRoutes(rh)
	// transactions
	handlers.SetupTransactionRoutes(rh)
}
