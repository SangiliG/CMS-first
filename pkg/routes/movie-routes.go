package routes

import (
	"github.com/SangiliG/cms-first/pkg/controllers"
	"github.com/SangiliG/cms-first/pkg/middleware"
	"github.com/gorilla/mux"
)

var RegisterMovieRoutes = func(router *mux.Router) {
	router.HandleFunc("/movies/", middleware.VerifyLogin(controllers.CreateMovie)).Methods("POST")
	router.HandleFunc("/movies/", controllers.GetMovie).Methods("GET")
	router.HandleFunc("/movies/{movieId}", controllers.GetMovieById).Methods("GET")
	router.HandleFunc("/movies/{movieId}", middleware.VerifyLogin(controllers.UpdateMovie)).Methods("PUT")
	router.HandleFunc("/movies/{movieId}", middleware.VerifyLogin(controllers.DeleteMovie)).Methods("DELETE")
	// 	router.HandleFunc("/movies/{string}", controllers.GetMoviesBySearch).Methods("GET")
}
