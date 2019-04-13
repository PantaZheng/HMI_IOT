package models

import (
	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model

	Title string
	Content string
	StartTime string
	EndTime string
	Status int

}