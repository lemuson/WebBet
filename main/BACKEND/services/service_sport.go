package services

import (
	"Web-Bet/main/BACKEND/models"

	"gorm.io/gorm"
)

type SportService struct {
	database *gorm.DB
}

// Создание SportSevice
func Create_SportService(database *gorm.DB) *SportService {
	return &SportService{database: database}
}

// SELECT *
func (service *SportService) GetAll() ([]models.Sport, error) {
	var sports []models.Sport

	result := service.database.Find(&sports)
	if result.Error != nil {
		return nil, result.Error
	}

	return sports, nil
}

// SELECT {id}
func (service *SportService) GetByID(ID uint) (*models.Sport, error) {
	var sport models.Sport

	result := service.database.First(&sport, "ID_ВидСпорта = ?", ID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &sport, nil
}

// INSERT
func (service *SportService) Create(sport *models.Sport) (*models.Sport, error) {
	result := service.database.Create(sport)
	if result.Error != nil {
		return nil, result.Error
	}
	return sport, nil
}

// UPDATE {ID}
func (service *SportService) Update(id uint, updated *models.Sport) (*models.Sport, error) {
	var sport models.Sport

	if err := service.database.First(&sport, "ID_ВидСпорта = ?", id).Error; err != nil {
		return nil, err
	}

	sport.Название = updated.Название
	sport.Изображение = updated.Изображение
	if err := service.database.Save(&sport).Error; err != nil {
		return nil, err
	}

	return &sport, nil
}

// DELETE {ID}
func (service *SportService) Delete(id uint) error {
	var sport models.Sport
	if err := service.database.First(&sport, "ID_ВидСпорта = ?", id).Error; err != nil {
		return err
	}

	return service.database.Delete(&sport).Error
}

// // SELECT {name}
// func (service *SportService) GetByName(name string) (*models.Sport, error) {
// 	var sport models.Sport

// 	result := service.database.First(&sport, "Название = ?", name)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return &sport, nil
// }
