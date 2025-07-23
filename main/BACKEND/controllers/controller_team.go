package controllers

import (
	"Web-Bet/main/BACKEND/models"
	"Web-Bet/main/BACKEND/services"
	"net/http"
)

type TeamController struct {
	service *services.TeamService
}

// Создание контроллера
func Create_TeamController(service *services.TeamService) *TeamController {
	return &TeamController{service: service}
}

// ПОЛУЧЕНИЕ ВСЕХ ЗАПИСЕЙ
// GET /Web-Bet/api/teams
func (controller *TeamController) GetAll(response http.ResponseWriter, request *http.Request) {
	HandleGetAll(
		response,
		request,
		controller.service.GetAll,
		func(team []models.Team) any {
			return teams_JSON(team)
		})
}

// ПОЛУЧЕНИЕ ЗАПИСИ ПО ID
// GET /Web-Bet/api/teams/{id}
func (controller *TeamController) GetByID(response http.ResponseWriter, request *http.Request) {
	ID, ok := GetID(response, request)
	if !ok {
		return
	}

	HandleGetByID(
		response,
		request,
		ID,
		controller.service.GetByID,
		func(team *models.Team) any {
			return team_JSON(*team)
		})
}

// СОЗДАНИЕ НОВОЙ ЗАПИСИ
// POST /Web-Bet/teams
func (controller *TeamController) Create(response http.ResponseWriter, request *http.Request) {
	HandleCreate(
		response,
		request,
		controller.service.Create,
		func(team *models.Team) any {
			return team_JSON(*team)
		})
}

// ОБНОВЛЕНИЕ ЗАПИСИ ПО ID
// PUT /Web-Bet/teams/{id}
func (controller *TeamController) Update(response http.ResponseWriter, request *http.Request) {
	ID, ok := GetID(response, request)
	if !ok {
		return
	}

	HandleUpdate(
		response,
		request,
		ID,
		controller.service.Update,
		func(team *models.Team) any {
			return team_JSON(*team)
		})
}

// УДАЛЕНИЕ ЗАПИСИ ПО ID
// DELETE /Web-Bet/teams/{id}
func (controller *TeamController) Delete(response http.ResponseWriter, request *http.Request) {
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

// JSON Team ~ TeamResponse
func team_JSON(team models.Team) models.TeamResponse {
	return models.TeamResponse{
		ID:          team.ID,
		Название:    team.Название,
		Изображение: team.Изображение,
	}
}

// JSON []Team ~ []TeamResponse
func teams_JSON(teams []models.Team) []models.TeamResponse {
	responses := make([]models.TeamResponse, len(teams))
	for i, team := range teams {
		responses[i] = team_JSON(team)
	}
	return responses
}
