package controllers

import (
	"Web-Bet/main/BACKEND/models"
	"Web-Bet/main/BACKEND/security"
	"Web-Bet/main/BACKEND/services"
	"net/http"
)

type BetController struct {
	service     *services.BetService
	userService *services.UserService
}

// Создание контроллера
func Create_BetController(service *services.BetService, userService *services.UserService) *BetController {
	return &BetController{service: service, userService: userService}
}

// ПОЛУЧЕНИЕ ВСЕХ ЗАПИСЕЙ
// GET /Web-Bet/api/bets
func (controller *BetController) GetAll(response http.ResponseWriter, request *http.Request) {
	HandleGetAll(
		response,
		request,
		controller.service.GetAll,
		func(bets []models.Bet) any {
			return bets_JSON(bets)
		})
}

// ПОЛУЧЕНИЕ ЗАПИСИ ПО ID
// GET /Web-Bet/api/bets/{id}
func (controller *BetController) GetByID(response http.ResponseWriter, request *http.Request) {
	ID, ok := GetID(response, request)
	if !ok {
		return
	}

	HandleGetByID(
		response,
		request,
		ID,
		controller.service.GetByID,
		func(bet *models.Bet) any {
			return bet_JSON(*bet)
		})

}

// СОЗДАНИЕ НОВОЙ ЗАПИСИ
// POST /Web-Bet/api/bets
func (controller *BetController) Create(response http.ResponseWriter, request *http.Request) {
	ID, ok := security.GetUserID_JWT(request)
	if !ok {
		http.Error(response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	HandleCreate(
		response,
		request,
		func(bet *models.Bet) (*models.Bet, error) {
			err := controller.userService.UpdateUserBalance(ID, -bet.Размер)
			if err != nil {
				return nil, err
			}
			bet.IDПользователь = ID
			return controller.service.Create(bet)
		},
		func(bet *models.Bet) any {
			return bet_JSON(*bet)
		},
	)
}

// ПРЕОБРАЗОВАНИЕ
// JSON Bet ~ BetResponse
func bet_JSON(bet models.Bet) models.BetResponse {
	return models.BetResponse{
		IDМатч:      bet.Прогноз.IDМатч,
		Матч:        bet.Прогноз.Матч.Команда1.Название + "-" + bet.Прогноз.Матч.Команда2.Название,
		Прогноз:     bet.Прогноз.Результат.Название,
		Размер:      bet.Размер,
		Коэффициент: bet.Коэффициент,
		Статус:      bet.СтатусСтавки.Название,
	}
}

// ПРЕОБРАЗОВАНИЕ
// JSON []Bet ~ []BetResponse
func bets_JSON(bets []models.Bet) []models.BetResponse {
	responses := make([]models.BetResponse, len(bets))
	for i, team := range bets {
		responses[i] = bet_JSON(team)
	}
	return responses
}
