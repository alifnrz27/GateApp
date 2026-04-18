package http

import (
	"GateApp/backend/gpio"
	"GateApp/backend/modules/gate/service"
	"GateApp/backend/modules/gate/validation"
	"GateApp/backend/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
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

	fmt.Println("Trigger relay...")

	err = gpio.TriggerRelay(17, 2)
	if err != nil {
		fmt.Println("Trigger error:", err)
		return nil
	}

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

	gate, err := h.service.TriggerGate(req.GateUUID, req.Trigger)
	if err != nil {
		response := utils.APIResponse("Error Trigger", fiber.StatusBadRequest, "error", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := utils.APIResponse("Gate triggered successful", fiber.StatusOK, "success", gate)
	return c.Status(fiber.StatusOK).JSON(response)
}
