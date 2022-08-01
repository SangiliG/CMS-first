package models

import (
	"github.com/SangiliG/cms-first/pkg/config"
	"github.com/jinzhu/gorm"
	
)

var sh *gorm.DB

type Show struct {
	gorm.Model
	Name     string `json:"name"`
	Language string `json:"language"`
	Genre    string `json:"genre"`
	Year     string `json:"year"`
}

func init() {
	config.Connect()
	sh = config.GetDB()
	sh.AutoMigrate(&Show{})
}

func (s *Show) CreateShow() *Show {
	sh.NewRecord(s)
	sh.Create(&s)
	return s
}

func GetAllShows() []Show {
    var Shows []Show
    sh.Find(&Shows)
    return Shows
}

func GetShowsById(Id int64) (*Show, *gorm.DB) {
	var getShow Show
	sh := se.Where("ID=?", Id).Find(&getShow)
	return &getShow, sh
}

func DeleteShow(ID int64) Show {
	var show Show
	sh.Where("ID=?", ID).Delete(show)
	return show
}
