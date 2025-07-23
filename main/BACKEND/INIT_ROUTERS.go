package backend

import (
	frontend "Web-Bet/main/FRONTEND"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func InitRoutes(router *mux.Router, database *gorm.DB) {
	//------------------------------------------------------------------------------------------------------------------------
	// Web-Bet/*
	page_router := router.PathPrefix("/Web-Bet").Subrouter()
	page_router.HandleFunc("/upload-image", upload)
	page_router.HandleFunc("/auth", frontend.AuthPage).Methods("GET")                     // - Web-Bet/auth
	page_router.HandleFunc("/matches", frontend.MatchesPage).Methods("GET")               //- Web-Bet/matches
	page_router.HandleFunc("/matches/sport/{sport}", frontend.MatchesPage).Methods("GET") // - Web-Bet/matches/{sport} (УБРАТЬ?)                                                                            // - Web-Bet/matches
	page_router.HandleFunc("/matches/{id}", frontend.MatchInfoPage).Methods("GET")        // - Web-Bet/matches/{id}
	page_router.HandleFunc("/profile", frontend.ProfilePage).Methods("GET")               // - Web-Bet/profile
	page_router.HandleFunc("/logout", ControllerUser.Logout).Methods("GET")               // - /Web-Bet/logout
	//--------------------------------------------------------------------------------------------------------------------------

	// /Web-Bet/api/*
	API_router := router.PathPrefix("/Web-Bet/api").Subrouter()

	// /Web-Bet/api/matches
	matches_router := API_router.PathPrefix("/matches").Subrouter()
	matches_router.HandleFunc("", ControllerMatch.GetAll).Methods("GET")                   // GET /Web-Bet/api/matches
	matches_router.HandleFunc("/{id}", ControllerMatch.GetByID).Methods("GET")             // GET /Web-Bet/api/matches/{id}
	matches_router.HandleFunc("/sport/{sport}", ControllerMatch.GetBySport).Methods("GET") // GET /Web-Bet/api/matches/sport{sport}
	matches_router.HandleFunc("", ControllerMatch.Create).Methods("POST")                  // POST /Web-Bet/api/matches
	matches_router.HandleFunc("/start/{id}", ControllerMatch.Start).Methods("POST")        // POST /Web-Bet/api/matches/start/{id}

	// /Web-Bet/api/teams
	teams_router := API_router.PathPrefix("/teams").Subrouter()
	teams_router.HandleFunc("", ControllerTeam.GetAll).Methods("GET")         // GET /Web-Bet/api/teams
	teams_router.HandleFunc("/{id}", ControllerTeam.GetByID).Methods("GET")   // GET /Web-Bet/api/teams/{id}
	teams_router.HandleFunc("", ControllerTeam.Create).Methods("POST")        // POST /Web-Bet/api/teams
	teams_router.HandleFunc("/{id}", ControllerTeam.Update).Methods("PUT")    // PUT /Web-Bet/api/teams
	teams_router.HandleFunc("/{id}", ControllerTeam.Delete).Methods("DELETE") // DELETE /Web-Bet/api/teams

	// /Web-Bet/api/sport
	sport_router := API_router.PathPrefix("/sports").Subrouter()
	sport_router.HandleFunc("", ControllerSport.GetAll).Methods("GET")         // GET /Web-Bet/api/sports
	sport_router.HandleFunc("/{id}", ControllerSport.GetByID).Methods("GET")   // GET /Web-Bet/api/sports/{id}
	sport_router.HandleFunc("", ControllerSport.Create).Methods("POST")        // POST /Web-Bet/api/sports
	sport_router.HandleFunc("/{id}", ControllerSport.Update).Methods("PUT")    // UPDATE /Web-Bet/api/sports/{id}
	sport_router.HandleFunc("/{id}", ControllerSport.Delete).Methods("DELETE") // DELETE /Web-Bet/api/sports/{id}

	// /Web-Bet/api/users
	user_router := API_router.PathPrefix("/users").Subrouter()
	user_router.HandleFunc("", ControllerUser.GetAll).Methods("GET")           // GET /Web-Bet/api/users
	user_router.HandleFunc("/register", ControllerUser.Create).Methods("POST") // POST /Web-Bet/api/users
	user_router.HandleFunc("/login", ControllerUser.Login).Methods("POST")     // POST /Web-Bet/api/users/login
	user_router.HandleFunc("/me", ControllerUser.GetMe).Methods("GET")         // GET /Web-Bet/api/users/me
	user_router.HandleFunc("/{id}", ControllerUser.GetByID).Methods("GET")     // GET /Web-Bet/api/users/{id}
	user_router.HandleFunc("/{id}", ControllerUser.Update).Methods("PUT")      // PUT /Web-Bet/api/users/{id}
	user_router.HandleFunc("/{id}", ControllerUser.UpdateMe).Methods("PUT")    // PUT /Web-Bet/api/users/me

	// /Web-Bet/api/bets
	bets_router := API_router.PathPrefix("/bets").Subrouter()
	bets_router.HandleFunc("", ControllerBet.GetAll).Methods("GET")       // GET /Web-Bet/api/bets
	bets_router.HandleFunc("/{id}", ControllerBet.GetByID).Methods("GET") // GET /Web-Bet/api/bets/{id}
	bets_router.HandleFunc("", ControllerBet.Create).Methods("POST")      // POST /Web-Bet/api/bets

	// /Web-Bet/api/results
	result_router := API_router.PathPrefix("/results").Subrouter()
	result_router.HandleFunc("", ControllerResult.GetAll).Methods("GET") // GET /Web-Bet/api/results
}
