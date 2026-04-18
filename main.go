package main

import (
	"embed"
	"log"

	"GateApp/backend/config"
	"GateApp/backend/gpio"

	"github.com/joho/godotenv"
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
