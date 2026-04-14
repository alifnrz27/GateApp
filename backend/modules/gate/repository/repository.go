package repository

import (
	"GateApp/backend/models"

	"gorm.io/gorm"
)

type gateRepository struct {
	db *gorm.DB
}

type GateRepository interface {
	GetAllGates() ([]models.Gate, error)
	FindByUUID(uuid string) (*models.Gate, error)
}
