package frontend

import (
	"Web-Bet/main/BACKEND/security"
	"net/http"
	"text/template"
)

// Страница авторизации
// /Web-Bet/auth
func AuthPage(response http.ResponseWriter, request *http.Request) {
	renderTemplate(response, "main/FRONTEND/HTML/auth.html")
}

// Страница профиля
// /Web-Bet/profile
func ProfilePage(response http.ResponseWriter, request *http.Request) {
	if security.IsUSER(request) {
		renderTemplate(response, "main/FRONTEND/HTML/user/user_profile.html")
		return
	}

	if security.IsADMIN(request) {
		renderTemplate(response, "main/FRONTEND/HTML/admin/admin_profile.html")
		return
	}

	http.Redirect(response, request, "/Web-Bet/auth", http.StatusFound)
}

// Страница матчей
// /Web-Bet/matches
func MatchesPage(response http.ResponseWriter, request *http.Request) {
	if security.IsADMIN(request) {
		renderTemplate(response, "main/FRONTEND/HTML/admin/admin_matches.html")
		return
	}
	renderTemplate(response, "main/FRONTEND/HTML/user/user_matches.html")
}

// Страница ставки
// /Web-Bet/match/{id}
func MatchInfoPage(response http.ResponseWriter, request *http.Request) {
	if security.IsUSER(request) {
		renderTemplate(response, "main/FRONTEND/HTML/user/user_match_info.html")
		return
	}

	if security.IsADMIN(request) {
		renderTemplate(response, "main/FRONTEND/HTML/admin/admin_match_info.html")
		return
	}

	http.Redirect(response, request, "/Web-Bet/auth", http.StatusFound)
}

// ЗАГРУЗКА HTML
func renderTemplate(response http.ResponseWriter, HTML_path string) {
	HTML_template, err := template.ParseFiles(HTML_path)
	if err != nil {
		http.Error(response, "Ошибка загрузки страницы: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := HTML_template.Execute(response, nil); err != nil {
		http.Error(response, "Ошибка отображения страницы: "+err.Error(), http.StatusInternalServerError)
	}
}

// ТЕСТОВАЯ СТРАНИЦА
func TestPage(response http.ResponseWriter, request *http.Request) {
	renderTemplate(response, "main/FRONTEND/HTML/test_user.html")
}
