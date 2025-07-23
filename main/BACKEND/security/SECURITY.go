package security

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/brianvoe/sjwt"
)

var JWT_keys = map[string][]byte{
	"Администратор": []byte("Zaitsev-SECRET-KEY-FOR-ADMIN"),
	"Пользователь":  []byte("Zaitsev-SECRET-KEY-FOR-USER"),
}

// ПОЛУЧЕНИЕ ТОКЕНА ИЗ COOKIES
func getToken(request *http.Request) string {
	cookie, err := request.Cookie("token")
	if err == nil {
		return cookie.Value
	}

	return "NO-JWT"
}

// ПОЛУЧЕНИЕ userID ИЗ JWT (ТОЛЬКО ДЛЯ ПОЛЬЗОВАТЕЛЕЙ)
func GetUserID_JWT(request *http.Request) (uint, bool) {
	tokenFromRequest := getToken(request)
	if tokenFromRequest == "NO-JWT" {
		return 0, false
	}

	token := strings.TrimPrefix(tokenFromRequest, "Bearer ")
	if !sjwt.Verify(token, JWT_keys["Пользователь"]) {
		return 0, false
	}

	claims, _ := sjwt.Parse(token)
	if err := claims.Validate(); err != nil {
		return 0, false
	}

	userIDStr, err := claims.GetStr("userId")
	if err != nil {
		return 0, false
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return 0, false
	}
	return uint(userID), true
}

// ПРОВЕРКА НА ПОЛЬЗОВАТЕЛЯ
func IsUSER(request *http.Request) bool {
	token := getToken(request)
	return checkJWT(token, JWT_keys["Пользователь"])
}

// ПРОВЕРКА НА АДМИНА
func IsADMIN(request *http.Request) bool {
	token := getToken(request)
	return checkJWT(token, JWT_keys["Администратор"])
}

// ПРОВЕРКА JWT
func checkJWT(token string, JWT_key []byte) bool {
	if token == "NO-JWT" {
		return false
	}

	token = strings.TrimPrefix(token, "Bearer ")
	if !sjwt.Verify(token, JWT_key) {
		return false
	}

	claims, _ := sjwt.Parse(token)
	if err := claims.Validate(); err != nil {
		return false
	}
	return true
}
