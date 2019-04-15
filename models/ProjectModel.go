package models

import (
	"../database"
	"github.com/jinzhu/gorm"
	"log"
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

func init(){
	database.DB.DropTable("projects")
	log.Printf("删除用户表")
	database.DB.Set("gorm:table_options", "DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").AutoMigrate(&Project{})
}

func GetLeaders(id uint)(leaders []User){
	database.DB.Find(&leaders,id).Select("leaders")
	return
}

func GetInstructors(id uint)(instructors []User){
	database.DB.Find(&instructors,id).Select("instructors")
	return
}

func EnrollProject(project *Project){
	recordProject:=Project{}
	database.DB.FirstOrCreate(&recordProject,&Project{Title:recordProject.Title})

}
