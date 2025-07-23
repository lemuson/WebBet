package services

import (
	"Web-Bet/main/BACKEND/models"

	"gorm.io/gorm"
)

type ResultService struct {
	database *gorm.DB
}

// Создание BetService
func Create_ResultService(database *gorm.DB) *ResultService {
	return &ResultService{database: database}
}

// SELECT *
func (service *ResultService) GetAll() ([]models.Result, error) {
	var results []models.Result

	result := service.database.Find(&results)
	if result.Error != nil {
		return nil, result.Error
	}

	return results, nil
}
