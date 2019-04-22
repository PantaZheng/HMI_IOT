package models

import "github.com/jinzhu/gorm"

type Module struct{
	gorm.Model
	Name		string		`json:"name"`
	Creator 	string		`json:"creator"`
	CreateTime  string	  	`json:"createTime"`//创建时间
	StartTime   string    	`json:"startTime"`//开始时间
	EndTime     string    	`json:"endTime"`//结束时间

}