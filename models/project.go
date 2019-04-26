package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/database"
)

type Project struct {
	gorm.Model
	Name       		string
	Type        	string
	Creator 		string
	CreateTime  	string
	StartTime   	string
 	EndTime     	string
	Content      	string
	Target			[]string
	LeaderID	 	uint										//一对多外键
	Leader       	User                                       	//belongs to
	Participants 	[]*UserBriefJson							//参与人员
	TagResult       bool
	TagSet			[]Tag
	Status      	int
	Acceptances  	string
	Instructors 	[]User
	Missions    	[]Mission
}

type BriefProject struct {
	ID uint	`json:"id"`
	Name string	`json:"name"`
	StartTime   string    	`json:"startTime"`
	EndTime     string    	`json:"endTime"`
	PrincipalName    string    `json:"principal"`
	Tag	string `json:"tag"`
	Content string `json:"content"`
}


type Tag struct{
	gorm.Model
	Judge 	User
	Tag		bool
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
	database.DB.FirstOrCreate(&recordProject,&Project{Name:recordProject.Name})

}
