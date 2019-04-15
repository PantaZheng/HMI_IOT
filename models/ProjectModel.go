package models

import (
	"../database"
	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model
	Title       string    `json:"title"`
	StartTime   string    `json:"startTime"`
 	EndTime     string    `json:"endTime"`
	Status      int       `json:"status"`
	Acceptances  string    `json:"acceptance"`
	Type        string    `json:"type"`
	Leaders     []User    `json:"leaders"`
	Instructors []User    `json:"instructors"`
	Missions    []Mission `json:"missions"`
}

func GetLeaders(id uint)(leaders []User){
	database.DB.Find(&leaders,id).Select("leaders")
	return
}

func GetInstructors(id uint)(instructors []User){
	database.DB.Find(&instructors,id).Select("instructors")
	return
}

//func EnrollProject(project *Project)
