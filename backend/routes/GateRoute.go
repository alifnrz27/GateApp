package routes

import (
	"GateApp/backend/modules/gate/http"
	"GateApp/backend/modules/gate/repository"
	"GateApp/backend/modules/gate/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GateRouter(router fiber.Router, db *gorm.DB) {
	gateRepo := repository.NewGateRepository(db)

	gateService := service.NewGateService(gateRepo)

	gateHandler := http.NewGateHandler(gateService)

	http.GateRoutes(router, gateHandler)

}
