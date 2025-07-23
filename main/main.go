package main

import (
	backend "Web-Bet/main/BACKEND"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("main/frontend"))))
	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("main/frontend/uploads"))))

	DataBase := backend.InitDB()
	backend.InitControllers(DataBase)
	backend.InitRoutes(router, DataBase)

	log.Fatal(http.ListenAndServe(":8080", router))
}
