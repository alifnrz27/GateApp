package service

import (
	"GateApp/backend/models"
	"GateApp/backend/modules/gate/repository"
)

type gateService struct {
	repo repository.GateRepository
}

type GateService interface {
	GetAllGates() ([]models.Gate, error)
	FindByUUID(uuid string) (*models.Gate, error)
	TriggerGate(uuid string, trigger string) (*models.Gate, error)
}
