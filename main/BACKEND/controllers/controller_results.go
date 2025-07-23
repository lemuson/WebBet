package controllers

import (
	"Web-Bet/main/BACKEND/models"
	"Web-Bet/main/BACKEND/services"
	"net/http"
)

type ResultController struct {
	service *services.ResultService
}

// Создание контроллера
func Create_ResultController(service *services.ResultService) *ResultController {
	return &ResultController{service: service}
}

// ПОЛУЧЕНИЕ ВСЕХ ЗАПИСЕЙ
// GET /Web-Bet/api/sports
func (controller *ResultController) GetAll(response http.ResponseWriter, request *http.Request) {
	HandleGetAll(
		response,
		request,
		controller.service.GetAll,
		func(results []models.Result) any {
			return results_JSON(results)
		})
}

// ПРЕОБРАЗОВАНИЕ
// JSON Result ~ ResultResponse
func result_JSON(result models.Result) models.ResultResponse {
	return models.ResultResponse{
		ID:       result.ID,
		Название: result.Название,
	}
}

// ПРЕОБРАЗОВАНИЕ
// JSON []Result ~ []ResultResponse
func results_JSON(results []models.Result) []models.ResultResponse {
	responses := make([]models.ResultResponse, len(results))
	for i, result := range results {
		responses[i] = result_JSON(result)
	}
	return responses
}
