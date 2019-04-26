package models

import "github.com/jinzhu/gorm"

type Module struct{
	gorm.Model
	Name         string
	Creator      string
	CreateTime   string
	StartTime    string
	EndTime      string
	Content      string
	Tag          bool
	LeaderID	 uint										//一对多外键
	Leader       User                                       //belongs to
	Participants []*User 	`gorm:"many2many:user_modules"` //多对多
	ProjectID    uint                                       //归属项目
	Missions     []*Mission                          		//一对多
}

type ModuleJson struct{
	ID				uint				`json:"id"`
	Name			string				`json:"name"`
	Creator 		string				`json:"creator"`
	CreateTime  	string	  			`json:"create_time"`//创建时间
	StartTime   	string    			`json:"start_time"`//开始时间
	EndTime     	string    			`json:"end_time"`//结束时间
	Content			string				`json:"content"`
	Tag				string				`json:"tag"`
	Leader			UserBriefJson		`json:"leader"`
	Participants 	[]*UserBriefJson	`json:"participants"`//参与人员
	ProjectID		uint				`json:"project_id"`
	Missions		[]Mission			`json:"missions"`
}

type ModuleBriefJson struct{
	ID uint `json:"id"`
	Name			string				`json:"name"`
	CreateTime  	string	  			`json:"create_time"`//创建时间
	Content			string				`json:"content"`
	Tag				string				`json:"tag"`
}
