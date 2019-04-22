package models

import "github.com/jinzhu/gorm"

type Mission struct{
	gorm.Model
	Name       string		`json:"name"`//名称
	Creator 	string	  	`json:"creator"`//创建者
	CreateTime  string	  	`json:"createTime"`//创建时间
	StartTime   string    	`json:"startTime"`//开始时间
	EndTime     string    	`json:"endTime"`//结束时间
	Participants []string	`json:"participants"`//任务参与人员

	Content		string		`json:"content"`//任务详细内容


}