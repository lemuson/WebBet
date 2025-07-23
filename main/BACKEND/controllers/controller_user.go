package controllers

import (
	"Web-Bet/main/BACKEND/models"
	"Web-Bet/main/BACKEND/security"
	"Web-Bet/main/BACKEND/services"
	"encoding/json"
	"net/http"
	"time"

	"github.com/brianvoe/sjwt"
)

type UserController struct {
	service *services.UserService
}

// Создание контроллера
func Create_UserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}

// GET /Web-Bet/api/users
func (controller *UserController) GetAll(response http.ResponseWriter, request *http.Request) {
	HandleGetAll(
		response,
		request,
		controller.service.GetAll,
		func(user []models.User) any {
			return users_JSON(user)
		})
}

// GET /Web-Bet/api/users/{id}
func (controller *UserController) GetByID(response http.ResponseWriter, request *http.Request) {
	ID, ok := security.GetUserID_JWT(request)
	if !ok {
		http.Error(response, "Unauthorized", http.StatusUnauthorized)
		return
	}
	//ЕСЛИ АДМИН, ТО

	HandleGetByID(
		response,
		request,
		ID,
		controller.service.GetByID,
		func(user *models.User) any {
			return user_JSON(*user)
		})
}

// GET /Web-Bet/api/users/me
func (controller *UserController) GetMe(response http.ResponseWriter, request *http.Request) {
	ID, ok := security.GetUserID_JWT(request)
	if !ok {
		http.Error(response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	HandleGetByID(
		response,
		request,
		ID,
		controller.service.GetByID,
		func(user *models.User) any {
			return user_JSON(*user)
		})
}

// POST /Web-Bet/api/users
func (controller *UserController) Create(response http.ResponseWriter, request *http.Request) {
	HandleCreate(
		response,
		request,
		controller.service.Create,
		func(user *models.User) any {
			return user_JSON(*user)
		})
}

// PUT /Web-Bet/api/users
func (controller *UserController) Update(response http.ResponseWriter, request *http.Request) {

}

// PUT /Web-Bet/api/users/me
func (controller *UserController) UpdateMe(response http.ResponseWriter, request *http.Request) {
	ID, ok := security.GetUserID_JWT(request)
	if !ok {
		http.Error(response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var input models.User
	if err := json.NewDecoder(request.Body).Decode(&input); err != nil {
		http.Error(response, "Неизвестный JSON", http.StatusBadRequest)
		return
	}

	updatedUser, err := controller.service.Update(uint(ID), &input)
	if err != nil {
		http.Error(response, "Ошибка при обновлении пользователя", http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(user_JSON(*updatedUser))
}

// Вход в аккаунт
func (controller *UserController) Login(response http.ResponseWriter, request *http.Request) {
	var input models.LoginInput
	if err := json.NewDecoder(request.Body).Decode(&input); err != nil {
		http.Error(response, "Неверный формат запроса", http.StatusBadRequest)
		return
	}

	user, err := controller.service.GetByLoginAndPassword(input.Логин, input.Пароль)
	if err != nil {
		http.Error(response, "Неверный логин или пароль", http.StatusUnauthorized)
		return
	}

	token := sjwt.New()
	token.Set("userId", user.ID)
	token.Set("exp", time.Now().Add(time.Hour*3).Unix())
	tokenStr := token.Generate(security.JWT_keys[user.Роль.Название])

	http.SetCookie(response, &http.Cookie{
		Name:     "token",
		Value:    tokenStr,
		Expires:  time.Now().Add(3 * time.Hour),
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})
	response.WriteHeader(http.StatusOK)
}

// Выход из аккаунта
func (controller *UserController) Logout(response http.ResponseWriter, request *http.Request) {
	http.SetCookie(response, &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})
	http.Redirect(response, request, "/Web-Bet/auth", http.StatusSeeOther)
}

// Замена JSON структуры пользователя
func user_JSON(user models.User) models.UserResponse {
	response := models.UserResponse{
		ID:    user.ID,
		Логин: user.Логин,
	}

	if user.ДанныеПользователя != nil {
		response.Данные = models.UserDataResponse{
			Имя:     user.ДанныеПользователя.Имя,
			Телефон: user.ДанныеПользователя.Телефон,
			Баланс:  user.ДанныеПользователя.Баланс,
		}
	}

	bets := make([]models.BetResponse, len(user.Ставки))
	for j, bet := range user.Ставки {
		bets[j] = models.BetResponse{
			IDМатч:      bet.Прогноз.IDМатч,
			Матч:        bet.Прогноз.Матч.Команда1.Название + "-" + bet.Прогноз.Матч.Команда2.Название,
			Прогноз:     bet.Прогноз.Результат.Название,
			Размер:      bet.Размер,
			Коэффициент: bet.Коэффициент,
			Статус:      bet.СтатусСтавки.Название,
		}
	}
	response.Ставки = bets

	return response
}

// Замена JSON структуры пользователей
func users_JSON(users []models.User) []models.UserResponse {
	responses := make([]models.UserResponse, len(users))
	for i, user := range users {
		responses[i] = user_JSON(user)
	}
	return responses
}
