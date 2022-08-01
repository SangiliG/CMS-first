package models

import (
	"github.com/SangiliG/cms-first/pkg/config"
	"github.com/jinzhu/gorm"
	
)

var se *gorm.DB

type Season struct {
	gorm.Model
	Name     string `json:"name"`
	Language string `json:"language"`
	Number   string `json:"number"`
	Year     string `json:"year"`
}

func init() {
	config.Connect()
	se = config.GetDB()
	se.AutoMigrate(&Season{})
}

func (s *Season) CreateSeason() *Season {
	se.NewRecord(s)
	se.Create(&s)
	return s
}

func GetAllSeasons() []Season {
    var Seasons []Season
    se.Find(&Seasons)
    return Seasons
}

func GetSeasonsById(Id int64) (*Season, *gorm.DB) {
	var getSeason Season
	se := se.Where("ID=?", Id).Find(&getSeason)
	return &getSeason, se
}

func DeleteSeason(ID int64) Season {
	var season Season
	se.Where("ID=?", ID).Delete(season)
	return season
}
