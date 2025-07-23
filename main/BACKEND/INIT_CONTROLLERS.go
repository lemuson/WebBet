package backend

import (
	"Web-Bet/main/BACKEND/controllers"
	"Web-Bet/main/BACKEND/services"

	"gorm.io/gorm"
)

var ControllerMatch *controllers.MatchController
var ControllerTeam *controllers.TeamController
var ControllerSport *controllers.SportController
var ControllerUser *controllers.UserController
var ControllerBet *controllers.BetController
var ControllerResult *controllers.ResultController

func InitControllers(DataBase *gorm.DB) {
	matchService := services.Create_MatchService(DataBase)
	ControllerMatch = controllers.Create_MatchController(matchService)

	teamService := services.Create_TeamService(DataBase)
	ControllerTeam = controllers.Create_TeamController(teamService)

	sportService := services.Create_SportService(DataBase)
	ControllerSport = controllers.Create_SportController(sportService)

	userService := services.Create_UserService(DataBase)
	ControllerUser = controllers.Create_UserController(userService)

	betService := services.Create_BetService(DataBase)
	ControllerBet = controllers.Create_BetController(betService, userService)

	resultService := services.Create_ResultService(DataBase)
	ControllerResult = controllers.Create_ResultController(resultService)
}
