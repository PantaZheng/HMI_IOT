package models

import "github.com/jinzhu/gorm"

type Module struct{
	gorm.Model
	ID				uint	`gorm:"primary_key"`
	Name			string
	Creator 		string
	CreateTime  	string
	StartTime   	string
	EndTime     	string
	Content			string
	Tag				string
	Principal    	User	//一对一
	Participants 	[]*User 	`gorm:"many2many:user_modules"`//多对多
	Missions		[]*Mission	//一对多
	ProjectID		uint 		//归属项目
}

type ModuleJson struct{
	Name			string		`json:"name"`
	Creator 		string		`json:"creator"`
	CreateTime  	string	  	`json:"create_time"`//创建时间
	StartTime   	string    	`json:"start_time"`//开始时间
	EndTime     	string    	`json:"end_time"`//结束时间
	Participants 	[]string	`json:"participants"`//参与人员
	Content			string		`json:"content"`
	Tag				string		`json:"tag"`
	Missions		[]Mission	`json:"missions"`
}

