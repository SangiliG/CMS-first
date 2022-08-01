package routes

import (
	"github.com/SangiliG/cms-first/pkg/controllers"
	"github.com/SangiliG/cms-first/pkg/middleware"
	"github.com/gorilla/mux"
)

var RegisterSeasonRoutes = func(router *mux.Router) {
	router.HandleFunc("/seasons/", middleware.VerifyLogin(controllers.CreateSeason)).Methods("POST")
	router.HandleFunc("/seasons/", controllers.GetSeason).Methods("GET")
	router.HandleFunc("/seasons/{seasonId}", controllers.GetSeasonById).Methods("GET")
	router.HandleFunc("/seasons/{seasonId}", middleware.VerifyLogin(controllers.UpdateSeason)).Methods("PUT")
	router.HandleFunc("/seasons/{seasonId}", middleware.VerifyLogin(controllers.DeleteSeason)).Methods("DELETE")
}
