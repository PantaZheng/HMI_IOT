package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/database"
	"log"
)

type Project struct {
	gorm.Model
	Name       string		`json:"name"`//名称
	Type        string		`json:"type"`//类型
	Creator 	string	  	`json:"creator"`//创建者
	CreateTime  string	  	`json:"createTime"`//创建时间
	StartTime   string    	`json:"startTime"`//开始时间
 	EndTime     string    	`json:"endTime"`//结束时间

	PrincipalID	string	   	`json:"principal_id"`
	PrincipalName    string    `json:"principal_name"`//负责人姓名
	Principal    User    //负责人

	Status      int       `json:"status"`
	Acceptances  string    `json:"acceptance"`
	Instructors []User    `json:"instructors"`
	Missions    []Mission `json:"missions"`
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


func init(){
	database.DB.DropTable("projects")
	log.Printf("删除用户表")
	database.DB.Set("gorm:table_options", "DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").AutoMigrate(&Project{}).AddForeignKey("principal_id,principal_name","users(id,name)","no action","no action")
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
