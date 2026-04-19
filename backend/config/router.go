package config

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"GateApp/backend/routes"

	"gorm.io/gorm"
)

func Route(db *gorm.DB) {

	// ===============================
	// MAIN APP
	// ===============================
	app := fiber.New(fiber.Config{
		BodyLimit:     50 * 1024 * 1024,
		StrictRouting: false,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			return c.Status(code).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		},
	})

	// ===============================
	// GLOBAL MIDDLEWARE
	// ===============================
	app.Use(cors.New(cors.Config{
		AllowOrigins:  "*",
		AllowMethods:  "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:  "Origin, Content-Type, Accept, Authorization, X-Requested-With",
		ExposeHeaders: "Content-Length, Content-Type",
		MaxAge:        86400,
	}))

	// ===============================
	// API APP
	// ===============================
	api := fiber.New(fiber.Config{
		BodyLimit:     50 * 1024 * 1024,
		StrictRouting: false,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			return c.Status(code).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
			})
		},
	})

	// ===============================
	// API MIDDLEWARE
	// ===============================
	api.Use(cors.New(cors.Config{
		AllowOrigins:  "*",
		AllowMethods:  "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:  "Origin, Content-Type, Accept, Authorization, X-Requested-With",
		ExposeHeaders: "Content-Length, Content-Type",
		MaxAge:        86400,
	}))

	routes.GateRouter(api, db)

	// ===============================
	// MOUNT API
	// ===============================
	app.Mount("/api/v1", api)

	// ===============================
	// RUN APP
	// ===============================
	port := os.Getenv("PORT")
	if port == "" {
		port = "8006"
	}

	log.Println("Server running on port:", port)
	log.Fatalln(app.Listen(":" + port))
}
