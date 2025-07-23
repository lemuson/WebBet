package controllers

import (
	"Web-Bet/main/BACKEND/models"
	"Web-Bet/main/BACKEND/services"
	"net/http"
)

type SportController struct {
	service *services.SportService
}

// Создание контроллера
func Create_SportController(service *services.SportService) *SportController {
	return &SportController{service: service}
}

// ПОЛУЧЕНИЕ ВСЕХ ЗАПИСЕЙ
// GET /Web-Bet/api/sports
func (controller *SportController) GetAll(response http.ResponseWriter, request *http.Request) {
	HandleGetAll(
		response,
		request,
		controller.service.GetAll,
		func(sport []models.Sport) any {
			return sports_JSON(sport)
		})
}

// ПОЛУЧЕНИЕ ЗАПИСИ ПО ID
// GET /Web-Bet/api/sports/{id}
func (controller *SportController) GetByID(response http.ResponseWriter, request *http.Request) {
	ID, ok := GetID(response, request)
	if !ok {
		return
	}

	HandleGetByID(
		response,
		request,
		ID,
		controller.service.GetByID,
		func(sport *models.Sport) any {
			return sport_JSON(*sport)
		})
}

// СОЗДАНИЕ НОВОЙ ЗАПИСИ
// POST /Web-Bet/api/sports
func (controller *SportController) Create(response http.ResponseWriter, request *http.Request) {
	HandleCreate(
		response,
		request,
		controller.service.Create,
		func(sport *models.Sport) any {
			return sport_JSON(*sport)
		})
}

// ОБНОВЛЕНИЕ ЗАПИСИ ПО ID
// PUT /Web-Bet/api/sports/{id}
func (controller *SportController) Update(response http.ResponseWriter, request *http.Request) {
	ID, ok := GetID(response, request)
	if !ok {
		return
	}

	HandleUpdate(
		response,
		request,
		ID,
		controller.service.Update,
		func(sport *models.Sport) any {
			return sport_JSON(*sport)
		})
}

// УДАЛЕНИЕ ЗАПИСИ ПО ID
// DELETE /Web-Bet/api/sports/{id}
func (controller *SportController) Delete(response http.ResponseWriter, request *http.Request) {
	ID, ok := GetID(response, request)
	if !ok {
		return
	}

	HandleDelete(
		response,
		request,
		ID,
		controller.service.Delete,
	)
}

// ПРЕОБРАЗОВАНИЕ
// JSON Sport ~ SportResponse
func sport_JSON(sport models.Sport) models.SportResponse {
	return models.SportResponse{
		ID:          sport.ID,
		Название:    sport.Название,
		Изображение: sport.Изображение,
	}
}

// ПРЕОБРАЗОВАНИЕ
// JSON []Sport ~ []SportResponse
func sports_JSON(teams []models.Sport) []models.SportResponse {
	responses := make([]models.SportResponse, len(teams))
	for i, team := range teams {
		responses[i] = sport_JSON(team)
	}
	return responses
}
