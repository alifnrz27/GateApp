package http

import (
	"GateApp/backend/modules/gate/service"
	"GateApp/backend/modules/gate/validation"
	"GateApp/backend/utils"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stianeikeland/go-rpio/v4"
)

type GateHandler struct {
	service service.GateService
}

func NewGateHandler(service service.GateService) *GateHandler {
	return &GateHandler{service: service}
}

func (h *GateHandler) GetAllGates(c *fiber.Ctx) error {

	gates, err := h.service.GetAllGates()
	if err != nil {
		response := utils.APIResponse("Error get gates", fiber.StatusBadRequest, "error", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := utils.APIResponse("Get all gate successful", fiber.StatusOK, "success", gates)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *GateHandler) GetGateByUuid(c *fiber.Ctx) error {

	uuid := c.Params("uuid")
	gates, err := h.service.FindByUUID(uuid)
	if err != nil {
		response := utils.APIResponse("Error get gates", fiber.StatusBadRequest, "error", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := utils.APIResponse("Get all gate successful", fiber.StatusOK, "success", gates)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *GateHandler) Trigger(c *fiber.Ctx) error {
	fmt.Println("Done")
	req := new(validation.TriggerGateRequest)
	if err := c.BodyParser(req); err != nil {
		response := utils.APIResponse("Invalid request", fiber.StatusBadRequest, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	errors := validation.ValidateTriggerGate(req)
	if errors != nil {
		response := utils.APIResponse("Validation error", fiber.StatusBadRequest, "error", errors)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err := c.BodyParser(&req); err != nil {
		response := utils.APIResponse("Invalid request", fiber.StatusBadRequest, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// 🔥 Open GPIO
	if err := rpio.Open(); err != nil {
		fmt.Println("Failed to open GPIO:", err)
		return nil
	}
	defer rpio.Close()

	// 🔹 Setup pins (BCM mode default)
	pinRelay := rpio.Pin(22)

	pinRelay.Output()

	fmt.Println("Starting relay cycle...")

	for {
		// 🔹 Relay 22 ON
		pinRelay.High()
		time.Sleep(1 * time.Second)
		pinRelay.Low()
	}

	gate, err := h.service.TriggerGate(req.GateUUID, req.Trigger)
	if err != nil {
		response := utils.APIResponse("Error Trigger", fiber.StatusBadRequest, "error", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := utils.APIResponse("Gate triggered successful", fiber.StatusOK, "success", gate)
	return c.Status(fiber.StatusOK).JSON(response)
}
