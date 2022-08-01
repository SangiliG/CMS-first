package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SangiliG/cms-first/pkg/models"
	"github.com/SangiliG/cms-first/pkg/utils"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var NewMovie models.Movie

func GetMovie(w http.ResponseWriter, r *http.Request) {
	newMovies := models.GetAllMovies()
	res, _ := json.Marshal(newMovies)
	w.Header().Set("Content-Typ", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	movieDetails, _ := models.GetMoviesById(ID)
	res, _ := json.Marshal(movieDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	CreateMovie := &models.Movie{}
	utils.ParseBody(r, CreateMovie)
	b := CreateMovie.CreateMovie()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	movie := models.DeleteMovie(ID)
	res, _ := json.Marshal(movie)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {

	var updateMovie = &models.Movie{}

	utils.ParseBody(r, updateMovie)

	vars := mux.Vars(r)

	movieId := vars["movieId"]

	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("error while updating")
	}
	movieDetails, db := models.GetMoviesById(ID)
	if updateMovie.Name != "" {
		movieDetails.Name = updateMovie.Name
	}
	if updateMovie.Director != "" {
		movieDetails.Director = updateMovie.Director
	}
	if updateMovie.Rating != "" {
		movieDetails.Rating = updateMovie.Rating
	}
	if updateMovie.Year != "" {
		movieDetails.Year = updateMovie.Year
	}
	if updateMovie.Poster != "" {
		movieDetails.Poster = updateMovie.Poster
	}
	db.Save(&movieDetails)
	res, _ := json.Marshal(movieDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetMoviesBySearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	String := vars["String"]
	movieDetails, _ := models.GetMoviesBySearch(String)
	res, err := json.Marshal(movieDetails)
	if err != nil {
		fmt.Println("error while searching")
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// func GetAllMovies(c *gin.Context) {
// 	pagination := utils.GeneratePaginationFromRequest(c)
// 	var movie models.Movie
// 	movieLists, err := repo.GetAllMovies(&movie, &pagination)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": movieLists,
// 	})
// }
