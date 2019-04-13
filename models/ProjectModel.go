package models

import (
	"../database"
	"fmt"
	"github.com/chanxuehong/wechat/mp/user"
	"github.com/jinzhu/gorm"
	"log"
)

type Project struct {
	gorm.Model

	Title string
	Content string
	StartTime string
	EndTime string
	Status int

}