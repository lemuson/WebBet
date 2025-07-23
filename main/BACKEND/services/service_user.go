package services

import (
	"Web-Bet/main/BACKEND/models"
	"fmt"

	"gorm.io/gorm"
)

type UserService struct {
	database *gorm.DB
}

// Создание UserService
func Create_UserService(database *gorm.DB) *UserService {
	return &UserService{database: database}
}

// SELECT *
func (service *UserService) GetAll() ([]models.User, error) {
	var users []models.User

	result := service.database.
		Preload("ДанныеПользователя").
		Preload("Ставки").
		Preload("Ставки.Прогноз").
		Preload("Ставки.Прогноз.Матч.Команда1").
		Preload("Ставки.Прогноз.Матч.Команда2").
		Preload("Ставки.Прогноз.Результат").
		Preload("Ставки.СтатусСтавки").
		Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

// SELECT {id}
func (service *UserService) GetByID(ID uint) (*models.User, error) {
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

// SELECT {login, password}
func (service *UserService) GetByLoginAndPassword(login string, password string) (*models.User, error) {
	var user models.User

	err := service.database.
		Preload("ДанныеПользователя").
		Preload("Роль").
		Preload("Ставки").
		Preload("Ставки.Прогноз").
		Preload("Ставки.Прогноз.Матч.Команда1").
		Preload("Ставки.Прогноз.Матч.Команда2").
		Preload("Ставки.Прогноз.Результат").
		Preload("Ставки.СтатусСтавки").
		Where("Логин = ? AND Пароль = ?", login, password).
		First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

// INSERT
func (service *UserService) Create(input *models.User) (*models.User, error) {
	var existingUser models.User
	if err := service.database.First(&existingUser, "Логин = ?", input.Логин).Error; err == nil {
		return nil, fmt.Errorf("пользователь с таким логином уже существует")
	}

	if input.ДанныеПользователя != nil {
		input.ДанныеПользователя.Баланс = 1000
		if err := service.database.Create(input.ДанныеПользователя).Error; err != nil {
			return nil, err
		}
		input.IDДанныеПользователя = &input.ДанныеПользователя.ID
	}

	input.IDРоль = 1
	if err := service.database.Create(input).Error; err != nil {
		return nil, err
	}

	return input, nil
}

// UPDATE {id}
func (service *UserService) Update(id uint, input *models.User) (*models.User, error) {
	var user models.User

	if err := service.database.Preload("ДанныеПользователя").First(&user, "ID_Пользователь = ?", id).Error; err != nil {
		return nil, err
	}

	if input.Логин != "" {
		user.Логин = input.Логин
	}
	if input.Пароль != "" {
		user.Пароль = input.Пароль
	}

	if input.ДанныеПользователя != nil && user.ДанныеПользователя != nil {
		if input.ДанныеПользователя.Имя != "" {
			user.ДанныеПользователя.Имя = input.ДанныеПользователя.Имя
		}
		if input.ДанныеПользователя.Телефон != "" {
			user.ДанныеПользователя.Телефон = input.ДанныеПользователя.Телефон
		}
		if err := service.database.Save(&user).Error; err != nil {
			return nil, err
		}
	}
	if err := service.database.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// DELETE {id}
func (service *UserService) Delete(ID uint) error {
	var user models.User

	if err := service.database.First(&user, "ID_Пользователь = ?", ID).Error; err != nil {
		return err
	}

	if err := service.database.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (service *UserService) UpdateUserBalance(userID uint, delta float64) error {
	var user models.User
	if err := service.database.Preload("ДанныеПользователя").First(&user, userID).Error; err != nil {
		return err
	}

	if user.ДанныеПользователя == nil {
		return fmt.Errorf("у пользователя нет связанных данных")
	}

	if user.ДанныеПользователя.Баланс+delta < 0 {
		return fmt.Errorf("недостаточно средств")
	}

	user.ДанныеПользователя.Баланс += delta
	return service.database.Save(user.ДанныеПользователя).Error
}
