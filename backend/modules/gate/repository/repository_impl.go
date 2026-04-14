package repository

import (
	"GateApp/backend/models"

	"gorm.io/gorm"
)

func NewGateRepository(db *gorm.DB) GateRepository {
	return &gateRepository{db}
}

func (r *gateRepository) GetAllGates() ([]models.Gate, error) {
	var gates []models.Gate

	query := r.db.Model(&models.Gate{})

	err := query.Find(&gates).Error
	return gates, err
}

func (r *gateRepository) FindByUUID(uuid string) (*models.Gate, error) {
	var gate models.Gate
	err := r.db.Where("uuid = ?", uuid).First(&gate).Error
	if err != nil {
		return nil, err
	}
	return &gate, nil
}
