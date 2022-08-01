package routes

import (
	"github.com/SangiliG/cms-first/pkg/controllers"
	"github.com/SangiliG/cms-first/pkg/middleware"
	"github.com/gorilla/mux"
)

var RegisterShowRoutes = func(router *mux.Router) {
	router.HandleFunc("/shows/", middleware.VerifyLogin(controllers.CreateShow)).Methods("POST")
	router.HandleFunc("/shows/", controllers.GetShow).Methods("GET")
	router.HandleFunc("/shows/{showId}", controllers.GetShowById).Methods("GET")
	router.HandleFunc("/shows/{showId}", middleware.VerifyLogin(controllers.UpdateShow)).Methods("PUT")
	router.HandleFunc("/shows/{showId}", middleware.VerifyLogin(controllers.DeleteShow)).Methods("DELETE")
}
