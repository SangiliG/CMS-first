package routes

import (
	"github.com/SangiliG/cms-first/pkg/controllers"
	"github.com/SangiliG/cms-first/pkg/middleware"
	"github.com/gorilla/mux"
)

var RegisterEpisodeRoutes = func(router *mux.Router) {
	router.HandleFunc("/episodes/", middleware.VerifyLogin(controllers.CreateEpisode)).Methods("POST")
	router.HandleFunc("/episodes/", controllers.GetEpisode).Methods("GET")
	router.HandleFunc("/episodes/{episodeId}", controllers.GetEpisodeById).Methods("GET")
	router.HandleFunc("/episodes/{episodeId}", middleware.VerifyLogin(controllers.UpdateEpisode)).Methods("PUT")
	router.HandleFunc("/episodes/{episodeId}", middleware.VerifyLogin(controllers.DeleteEpisode)).Methods("DELETE")
}
