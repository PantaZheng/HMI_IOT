package models

import (
	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model
	Title string 	`gorm:"not null" json:"title"`
	StartTime string `gorm:"not null" json:"startTime"`
 	EndTime string 	`gorm:"not null" json:"endTime"`
	Status int `gorm:"not null" json:"endTime"`
	//Leader []User `gorm:`
}