package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SangiliG/cms-first/pkg/models"
	"github.com/SangiliG/cms-first/pkg/utils"
	"github.com/gorilla/mux"
)

var NewSeason models.Season

func GetSeason(w http.ResponseWriter, r *http.Request) {
    newSeasons := models.GetAllSeasons()
    res, _ := json.Marshal(newSeasons)
    w.Header().Set("Content-Type", "pkglication/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func GetSeasonById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	seasonId := vars["seasonId"]
	ID, err := strconv.ParseInt(seasonId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	seasonDetails, _ := models.GetSeasonsById(ID)
	res, _ := json.Marshal(seasonDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateSeason(w http.ResponseWriter, r *http.Request) {
	CreateSeason := &models.Season{}
	utils.ParseBody(r, CreateSeason)
	b := CreateSeason.CreateSeason()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteSeason(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	seasonId := vars["seasonId"]
	ID, err := strconv.ParseInt(seasonId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	season := models.DeleteSeason(ID)
	res, _ := json.Marshal(season)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateSeason(w http.ResponseWriter, r *http.Request) {
	var updateSeason = &models.Season{}
	utils.ParseBody(r, updateSeason)
	vars := mux.Vars(r)
	seasonId := vars["seasonId"]
	ID, err := strconv.ParseInt(seasonId, 0, 0)
	if err != nil {
		fmt.Println("error while updating")

	}
	seasonDetails, db := models.GetSeasonsById(ID)
	if updateSeason.Name != "" {
		seasonDetails.Name = updateSeason.Name
	}
	if updateSeason.Language != "" {
		seasonDetails.Language = updateSeason.Language
	}
	if updateSeason.Number != "" {
		seasonDetails.Number = updateSeason.Number
	}
	if updateSeason.Year != "" {
		seasonDetails.Year =updateSeason.Year
	}
	db.Save(&seasonDetails)
	res, _ := json.Marshal(seasonDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

