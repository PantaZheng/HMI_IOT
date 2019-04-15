package models

import (
	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model

	Title string 	`gorm:"not null VARCHAR(255)" json:"title"`
	Content string `gorm:"not null VARCHAR(255)" json:"content"`
	StartTime string `gorm:"not null VARCHAR(255)" json:"startTime"`
 	EndTime string 	`gorm:"not null VARCHAR(255)" json:"endTime"`
	Status int `gorm:"not null" json:"endTime"`
}