package http

import (
	"github.com/gofiber/fiber/v2"
)

func GateRoutes(router fiber.Router, handler *GateHandler) {
	authRouter := router.Group("/gates") // middleware.JWTMiddleware,
	// middleware.JWTUserContextMiddleware(),

	authRouter.Get("/", handler.GetAllGates)
	authRouter.Get("/:uuid", handler.GetGateByUuid)
	authRouter.Post("/trigger", handler.Trigger)
}
