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

var NewEpisode models.Episode

func GetEpisode(w http.ResponseWriter, r *http.Request) {
    newEpisodes := models.GetAllEpisodes()
    res, _ := json.Marshal(newEpisodes)
    w.Header().Set("Content-Type", "pkglication/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func GetEpisodeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	episodeId := vars["episodeId"]
	ID, err := strconv.ParseInt(episodeId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	episodeDetails, _ := models.GetEpisodesById(ID)
	res, _ := json.Marshal(episodeDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateEpisode(w http.ResponseWriter, r *http.Request) {
	CreateEpisode := &models.Episode{}
	utils.ParseBody(r, CreateEpisode)
	b := CreateEpisode.CreateEpisode()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteEpisode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	episodeId := vars["episodeId"]
	ID, err := strconv.ParseInt(episodeId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	episode := models.DeleteEpisode(ID)
	res, _ := json.Marshal(episode)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateEpisode(w http.ResponseWriter, r *http.Request) {
	var updateEpisode = &models.Episode{}
	utils.ParseBody(r, updateEpisode)
	vars := mux.Vars(r)
	episodeId := vars["episodeId"]
	ID, err := strconv.ParseInt(episodeId, 0, 0)
	if err != nil {
		fmt.Println("error while updating")

	}
	episodeDetails, db := models.GetEpisodesById(ID)
	if updateEpisode.Name != "" {
		episodeDetails.Name = updateEpisode.Name
	}
	if updateEpisode.Language!= "" {
		episodeDetails.Language = updateEpisode.Language
	}
	if updateEpisode.Number != "" {
		episodeDetails.Number = updateEpisode.Number
	}
	if updateEpisode.Year != "" {
		episodeDetails.Year = updateEpisode.Year
	}
	db.Save(&episodeDetails)
	res, _ := json.Marshal(episodeDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

