package main

import (
	"embed"
	"fmt"
	"log"
	"time"

	"GateApp/backend/config"
	"GateApp/backend/gpio"

	"github.com/joho/godotenv"
	"github.com/stianeikeland/go-rpio/v4"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	// ======================
	// HTTP SERVER (BACKGROUND)
	// ======================
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using environment variables instead")
	}

	// Connect to database
	db := config.Connect()

	// Start the application
	config.Route(db)

	// 🔥 Open GPIO
	if err := rpio.Open(); err != nil {
		fmt.Println("Failed to open GPIO:", err)
		return
	}
	defer rpio.Close()

	// 🔹 Setup pins (BCM mode default)
	pin20 := rpio.Pin(20)
	pin26 := rpio.Pin(26)
	pin21 := rpio.Pin(21)

	pin20.Output()
	pin26.Output()
	pin21.Output()

	fmt.Println("Starting relay cycle...")

	for {
		// 🔹 Relay 21 ON
		pin21.High()
		time.Sleep(1 * time.Second)
		pin21.Low()

		// 🔹 Relay 20 ON
		pin20.High()
		time.Sleep(1 * time.Second)
		pin20.Low()

		// 🔹 Relay 26 ON
		pin26.High()
		time.Sleep(1 * time.Second)
		pin26.Low()
	}

	err = gpio.Init()
	if err != nil {
		log.Println("GPIO init error:", err)
	}

	// ======================
	// WAILS APP (FRONTEND)
	// ======================
	app := &App{}

	err = wails.Run(&options.App{
		Title:     "Gate Control App",
		Width:     1024,
		Height:    768,
		OnStartup: app.startup,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
