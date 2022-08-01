package models

import (
	"github.com/SangiliG/cms-first/pkg/config"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Claims struct {
	Name string `json:"name"`
	Role string `json:"role"`
	jwt.StandardClaims
}

var db *gorm.DB

type Movie struct {
	gorm.Model
	Name     string `json:"name"`
	Director string `json:"director"`
	Rating   string `json:"rating"`
	Year     string `json:"year"`
	Poster   string `json:"poster"`
}

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Movie{})
}

func (m *Movie) CreateMovie() *Movie {
	db.NewRecord(m)
	db.Create(&m)
	return m
}

func GetAllMovies() []Movie {
	var Movies []Movie
	db.Find(&Movies)
	return Movies
}

func GetMoviesById(Id int64) (*Movie, *gorm.DB) {
	var getMovie Movie
	db := db.Where("ID=?", Id).Find(&getMovie)
	return &getMovie, db
}

func DeleteMovie(ID int64) Movie {
	var movie Movie
	db.Where("ID=?", ID).Delete(movie)
	return movie
}

func GetMoviesBySearch(String string) (*Movie, *gorm.DB) {
	var movieDetails Movie
	// attribute := StringGet.String
	db := db.Where("LIKE=?", String).Find(movieDetails)
	return &movieDetails, db
}

// func GetMovieByPage(LIMIT int64) []Movie {
// 	var getMovies []Movie
// 	db := db.Limit("LIMIT=?", LIMIT).Find(&getMovies)
// 	return &getMovies, db
// }
