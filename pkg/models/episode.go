package models

import (
	"github.com/SangiliG/cms-first/pkg/config"
	"github.com/jinzhu/gorm"
)

var ep *gorm.DB

type Episode struct {
	gorm.Model
	Name     string `json:"name"`
	Language string `json:"language"`
	Number   string `json:"number"`
	Year     string `json:"year"`
}

func init() {
	config.Connect()
	ep = config.GetDB()
	ep.AutoMigrate(&Episode{})
}

func (e *Episode) CreateEpisode() *Episode {
	ep.NewRecord(e)
	ep.Create(&e)
	return e
}

func GetAllEpisodes() []Episode {
	var Episodes []Episode
	ep.Find(&Episodes)
	return Episodes
}

func GetEpisodesById(Id int64) (*Episode, *gorm.DB) {
	var getEpisode Episode
	ep := db.Where("ID=?", Id).Find(&getEpisode)
	return &getEpisode, ep
}

func DeleteEpisode(ID int64) Episode {
	var episode Episode
	ep.Where("ID=?", ID).Delete(episode)
	return episode
}
