package services

import (
	"Web-Bet/main/BACKEND/models"

	"gorm.io/gorm"
)

type TeamService struct {
	database *gorm.DB
}

// Создание TeamService
func Create_TeamService(database *gorm.DB) *TeamService {
	return &TeamService{database: database}
}

// SELECT *
func (service *TeamService) GetAll() ([]models.Team, error) {
	var teams []models.Team

	result := service.database.Find(&teams)
	if result.Error != nil {
		return nil, result.Error
	}

	return teams, nil
}

// SELECT {id}
func (service *TeamService) GetByID(ID uint) (*models.Team, error) {
	var team models.Team

	result := service.database.First(&team, "ID_Команда = ?", ID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &team, nil
}

// INSERT
func (service *TeamService) Create(team *models.Team) (*models.Team, error) {
	if err := service.database.Create(team).Error; err != nil {
		return nil, err
	}
	return team, nil
}

// UPDATE {id}
func (service *TeamService) Update(id uint, newTeam *models.Team) (*models.Team, error) {
	var updateTeam models.Team

	if err := service.database.First(&updateTeam, "ID_Команда = ?", id).Error; err != nil {
		return nil, err
	}

	updateTeam.Название = newTeam.Название
	updateTeam.Изображение = newTeam.Изображение

	if err := service.database.Save(&updateTeam).Error; err != nil {
		return nil, err
	}

	return &updateTeam, nil
}

// DELETE {id}
func (service *TeamService) Delete(id uint) error {
	var team models.Team
	if err := service.database.First(&team, "ID_Команда = ?", id).Error; err != nil {
		return err
	}

	if err := service.database.Delete(&team).Error; err != nil {
		return err
	}

	return nil
}
