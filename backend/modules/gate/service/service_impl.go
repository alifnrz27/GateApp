// service/gate_service.go
package service

import (
	"GateApp/backend/gpio"
	"GateApp/backend/models"
	"GateApp/backend/modules/gate/repository"
	"errors"
)

func NewGateService(repo repository.GateRepository) GateService {
	return &gateService{repo}
}

func (s *gateService) GetAllGates() ([]models.Gate, error) {
	gates, err := s.repo.GetAllGates()
	if err != nil {
		return nil, errors.New("Gates not found")
	}

	return gates, nil
}

func (s *gateService) FindByUUID(uuid string) (*models.Gate, error) {
	gate, err := s.repo.FindByUUID(uuid)
	if err != nil {
		return nil, errors.New("Gate not found")
	}

	return gate, nil
}

func (s *gateService) TriggerGate(uuid string, trigger string) (*models.Gate, error) {

	if trigger != "open" {
		return nil, errors.New("Trigger is not valid.")
	}

	gate, err := s.repo.FindByUUID(uuid)
	if err != nil {
		return nil, errors.New("Gate not found")
	}

	err = gpio.TriggerRelay(gate.Pin)
	if err != nil {
		return nil, errors.New("Failed to trigger relay")
	}

	return gate, nil
}
