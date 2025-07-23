package services

import (
	"Web-Bet/main/BACKEND/models"

	"gorm.io/gorm"
)

type BetService struct {
	database *gorm.DB
}

// Создание BetService
func Create_BetService(database *gorm.DB) *BetService {
	return &BetService{database: database}
}

// SELECT *
func (service *BetService) GetAll() ([]models.Bet, error) {
	var bets []models.Bet

	result := service.database.
		Preload("Прогноз").
		Preload("Прогноз.Матч.Команда1").
		Preload("Прогноз.Матч.Команда2").
		Preload("Прогноз.Результат").
		Preload("СтатусСтавки").
		Find(&bets)
	if result.Error != nil {
		return nil, result.Error
	}

	return bets, nil
}

// SELECT {id}
func (service *BetService) GetByID(ID uint) (*models.Bet, error) {
	var bet models.Bet

	result := service.database.
		Preload("Прогноз").
		Preload("Прогноз.Матч.Команда1").
		Preload("Прогноз.Матч.Команда2").
		Preload("Прогноз.Результат").
		Preload("СтатусСтавки").
		First(&bet, "ID_Ставка = ?", ID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &bet, nil
}

// INSERT
func (service *BetService) Create(input *models.Bet) (*models.Bet, error) {
	if input.IDСтатусСтавки == 0 {
		input.IDСтатусСтавки = 1
	}

	if err := service.database.Create(input).Error; err != nil {
		return nil, err
	}

	var fullBet models.Bet
	if err := service.database.
		Preload("Прогноз").
		Preload("Прогноз.Матч.Команда1").
		Preload("Прогноз.Матч.Команда2").
		Preload("Прогноз.Результат").
		Preload("СтатусСтавки").
		First(&fullBet, input.ID).Error; err != nil {
		return nil, err
	}

	return &fullBet, nil
}

func (service *BetService) GetUser(ID uint) (*models.User, error) {
	var user models.User

	result := service.database.
		Preload("ДанныеПользователя").
		Preload("Ставки").
		Preload("Ставки.Прогноз").
		Preload("Ставки.Прогноз.Матч.Команда1").
		Preload("Ставки.Прогноз.Матч.Команда2").
		Preload("Ставки.Прогноз.Результат").
		Preload("Ставки.СтатусСтавки").
		First(&user, "ID_Пользователь = ?", ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
