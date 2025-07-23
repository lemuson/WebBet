package services

import (
	"Web-Bet/main/BACKEND/models"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type MatchService struct {
	database *gorm.DB
}

// Создание MatchService
func Create_MatchService(database *gorm.DB) *MatchService {
	return &MatchService{database: database}
}

// SELECT *
func (service *MatchService) GetAll() ([]models.Match, error) {
	var matches []models.Match

	result := service.database.
		Preload("Команда1").
		Preload("Команда1").
		Preload("Команда2").
		Preload("Прогнозы").
		Preload("Прогнозы.Результат").
		Where("ID_Результат IS NULL").
		Find(&matches)

	if result.Error != nil {
		return nil, result.Error
	}
	return matches, nil
}

// SELECT {id}
func (service *MatchService) GetByID(ID uint) (*models.Match, error) {
	var match models.Match

	result := service.database.
		Preload("Команда1").
		Preload("Команда2").
		Preload("Результат").
		Preload("Прогнозы").
		Preload("Прогнозы.Результат").
		First(&match, "ID_Матч = ?", ID)

	if result.Error != nil {
		return nil, result.Error
	}
	return &match, nil
}

// INSERT
func (service *MatchService) Create(input *models.Match) (*models.Match, error) {
	if err := service.database.Create(input).Error; err != nil {
		return nil, err
	}

	var createdMatch models.Match
	err := service.database.
		Preload("Команда1").
		Preload("Команда2").
		Preload("Результат").
		Preload("Прогнозы").
		Preload("Прогнозы.Результат").
		First(&createdMatch, input.ID).Error
	if err != nil {
		return nil, err
	}

	return &createdMatch, nil
}

//UPDATE {id}

//DELETE {id}

// START
func (service *MatchService) Start(ID uint) (*models.Match, error) {
	match, err := service.GetByID(ID)
	if err != nil {
		return nil, err
	}

	var results []models.Result
	if err := service.database.Find(&results).Error; err != nil || len(results) == 0 {
		return nil, err
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomResult := results[r.Intn(len(results))]
	match.Результат = randomResult
	if err := service.database.Save(&match).Error; err != nil {
		return nil, err
	}

	var bets []models.Bet
	if err := service.database.
		Preload("Прогноз").
		Preload("Пользователь").
		Preload("Пользователь.ДанныеПользователя").
		Where("ID_Прогноз IN (SELECT ID_Прогноз FROM Прогнозы WHERE ID_Матч = ?)", match.ID).
		Find(&bets).Error; err != nil {
		return nil, err
	}

	var winStatus, loseStatus models.BetStatus
	if err := service.database.Where("Название = ?", "Выигрыш").First(&winStatus).Error; err != nil {
		return nil, err
	}
	if err := service.database.Where("Название = ?", "Проигрыш").First(&loseStatus).Error; err != nil {
		return nil, err
	}

	for _, bet := range bets {
		isWin := bet.Прогноз.IDРезультат == match.Результат.ID

		if isWin {
			bet.IDСтатусСтавки = winStatus.ID
		} else {
			bet.IDСтатусСтавки = loseStatus.ID
		}

		if err := service.database.Save(&bet).Error; err != nil {
			return nil, err
		}

		if isWin {
			var userData models.UserData
			if bet.Пользователь.IDДанныеПользователя != nil {
				if err := service.database.First(&userData, *bet.Пользователь.IDДанныеПользователя).Error; err != nil {
					return nil, err
				}
				userData.Баланс += bet.Размер * bet.Коэффициент
				if err := service.database.Save(&userData).Error; err != nil {
					return nil, err
				}
			}
		}
	}

	return match, nil
}

// CUSTOM
func (service *MatchService) Custom(where string, args any, join string) ([]models.Match, error) {
	var matches []models.Match

	db := service.database
	if join != "" {
		db = db.Joins(join)
	}

	if where != "" {
		db = db.Where(where, args)
	}

	result := db.
		Preload("Команда1").
		Preload("Команда2").
		Preload("ВидСпорта").
		Preload("Результат").
		Preload("Прогнозы").
		Preload("Прогнозы.Результат").
		Find(&matches)

	if result.Error != nil {
		return nil, result.Error
	}

	return matches, nil
}
