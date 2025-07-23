package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Получение ID
func GetID(response http.ResponseWriter, request *http.Request) (uint, bool) {
	vars := mux.Vars(request)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(response, "Не указан ID", http.StatusBadRequest)
		return 0, false
	}

	ID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(response, "Неверный формат ID", http.StatusBadRequest)
		return 0, false
	}

	return uint(ID), true
}

// ФУНКЦИЯ
// SELECT *
func HandleGetAll[R any](
	response http.ResponseWriter,
	request *http.Request,
	GET_function func() ([]R, error),
	JSON_function func([]R) any,
) {
	result, err := GET_function()
	if err != nil {
		http.Error(response, "Не найдено", http.StatusNotFound)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(JSON_function(result))
}

// ФУНКЦИЯ
// SELECT {id}
func HandleGetByID[R any](
	response http.ResponseWriter,
	request *http.Request,
	ID uint,
	GET_function func(uint) (R, error),
	JSON_function func(R) any,
) {
	result, err := GET_function(ID)
	if err != nil {
		http.Error(response, "Не найдено", http.StatusNotFound)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(JSON_function(result))
}

// ФУНКЦИЯ
// INSERT
func HandleCreate[T any, R any](
	response http.ResponseWriter,
	request *http.Request,
	CREATE_function func(*T) (R, error),
	JSON_function func(R) any,
) {
	var input T
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&input); err != nil {
		http.Error(response, "Ошибка декодирования JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	result, err := CREATE_function(&input)
	if err != nil {
		http.Error(response, "Ошибка создания", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(JSON_function(result))
}

// ФУНКЦИЯ
// UPDATE
func HandleUpdate[T any, R any](
	response http.ResponseWriter,
	request *http.Request,
	ID uint,
	UPDATE_function func(uint, *T) (R, error),
	JSON_function func(R) any,
) {
	var input T
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&input); err != nil {
		http.Error(response, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	result, err := UPDATE_function(ID, &input)
	if err != nil {
		http.Error(response, "Ошибка обновления", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(JSON_function(result))
}

// ФУНКЦИЯ
// DELETE
func HandleDelete(
	response http.ResponseWriter,
	request *http.Request,
	ID uint,
	DELETE_function func(uint) error,
) {
	err := DELETE_function(ID)
	if err != nil {
		http.Error(response, "Ошибка удаления", http.StatusInternalServerError)
		return
	}
	response.WriteHeader(http.StatusNoContent)
}

// ФУНКЦИЯ
// CUSTOM
func HandleCustom[R any](
	response http.ResponseWriter,
	request *http.Request,
	join string,
	where string,
	args any,
	customFunc func(where string, args any, join string) ([]R, error),
	jsonFunc func([]R) any,
) {
	result, err := customFunc(where, args, join)
	if err != nil {
		http.Error(response, "Ошибка запроса", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(jsonFunc(result))
}
