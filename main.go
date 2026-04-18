package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"GateApp/backend/gpio"
)

func main() {
	// 🔹 Load ENV
	_ = godotenv.Load()

	// 🔹 Init GPIO (sekali saja)
	if err := gpio.Init(); err != nil {
		log.Fatal("GPIO init error:", err)
	}

	// 🔹 Fiber app
	app := fiber.New()

	// 🔹 Routes
	app.Get("/trigger", func(c *fiber.Ctx) error {
		err := gpio.TriggerRelay(17, 2)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.SendString("Relay triggered")
	})

	// 🔹 Graceful shutdown
	go func() {
		port := os.Getenv("APP_PORT")
		if port == "" {
			port = "8006"
		}

		if err := app.Listen(port); err != nil {
			log.Fatal(err)
		}
	}()

	// 🔹 Handle CTRL+C
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down...")

	// 🔹 Close GPIO
	gpio.Close()

	// 🔹 Shutdown Fiber
	if err := app.Shutdown(); err != nil {
		log.Println("Server shutdown error:", err)
	}

	log.Println("Server exited")
}
