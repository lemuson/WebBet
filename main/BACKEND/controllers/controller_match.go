package controllers

import (
	"Web-Bet/main/BACKEND/models"
	"Web-Bet/main/BACKEND/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type MatchController struct {
	service *services.MatchService
}

// Создание контроллера
func Create_MatchController(service *services.MatchService) *MatchController {
	return &MatchController{service: service}
}

// ПОЛУЧЕНИЕ ВСЕХ ЗАПИСЕЙ
// GET /Web-Bet/api/matches
func (controller *MatchController) GetAll(response http.ResponseWriter, request *http.Request) {
	HandleGetAll(
		response,
		request,
		controller.service.GetAll,
		func(match []models.Match) any {
			return matches_JSON(match)
		})
}

// ПОЛУЧЕНИЕ ЗАПИСИ ПО ID
// GET /Web-Bet/api/matches/{id}
func (controller *MatchController) GetByID(response http.ResponseWriter, request *http.Request) {
	ID, ok := GetID(response, request)
	if !ok {
		return
	}

	HandleGetByID(
		response,
		request,
		ID,
		controller.service.GetByID,
		func(match *models.Match) any {
			return match_JSON(*match)
		})
}

// ПОЛУЧЕНИЕ ЗАПИСЕЙ ПО ВИДАМ СПОРТА
// GET /Web-Bet/api/matches/sport/{sport}
func (controller *MatchController) GetBySport(response http.ResponseWriter, request *http.Request) {
	sport, ok := mux.Vars(request)["sport"]
	if !ok {
		http.Error(response, "Не указан вид спорта", http.StatusBadRequest)
		return
	}

	HandleCustom(
		response,
		request,
		"JOIN ВидыСпорта ON ВидыСпорта.ID_ВидСпорта = Матчи.ID_ВидСпорта",
		"ВидыСпорта.Название = ? AND ID_Результат IS NULL",
		sport,
		controller.service.Custom,
		func(match []models.Match) any {
			return matches_JSON(match)
		})
}

// СОЗДАНИЕ ЗАПИСИ
// POST /Web-Bet/api/matches
func (controller *MatchController) Create(response http.ResponseWriter, request *http.Request) {
	HandleCreate(
		response,
		request,
		controller.service.Create,
		func(match *models.Match) any {
			return match_JSON(*match)
		})
}

//ОБНОВЛЕНИЕ ЗАПИСИ ПО ID
//PUT /Web-Bet/api/matches/{id}

// СТАРТ МАТЧА
// POST /Web-Bet/api/matches/start/{id}
func (controller *MatchController) Start(response http.ResponseWriter, request *http.Request) {
	ID, ok := GetID(response, request)
	if !ok {
		return
	}

	result, err := controller.service.Start(ID)
	if err != nil {
		http.Error(response, "Ошибка запроса", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(match_JSON(*result))
}

// ПРЕОБРАЗОВАНИЕ ОДНОЙ ЗАПИСИ
// JSON Match ~ MatchResponse
func match_JSON(match models.Match) models.MatchResponse {
	predictions := make([]models.PredictionResponse, len(match.Прогнозы))
	for i, prediction := range match.Прогнозы {
		predictions[i] = models.PredictionResponse{
			ID:          prediction.ID,
			Коэффициент: prediction.Коэффициент,
			Название:    prediction.Результат.Название,
		}
	}

	return models.MatchResponse{
		ID:        match.ID,
		Дата:      match.Дата,
		Результат: match.Результат.Название,
		Команда1: models.TeamResponse{
			Название:    match.Команда1.Название,
			Изображение: match.Команда1.Изображение,
		},
		Команда2: models.TeamResponse{
			Название:    match.Команда2.Название,
			Изображение: match.Команда2.Изображение,
		},
		Прогнозы: predictions,
	}
}

// ПРЕОБРАЗОВАНИЕ МАССИВА ЗАПИСЕЙ
// JSON []Match ~ []MatchResponse
func matches_JSON(matches []models.Match) []models.MatchResponse {
	responses := make([]models.MatchResponse, len(matches))
	for i, match := range matches {
		responses[i] = match_JSON(match)
	}
	return responses
}
