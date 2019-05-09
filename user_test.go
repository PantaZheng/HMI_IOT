package main

import (
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/database"
	"log"
	"testing"
)

//User 数据库用户表
type User struct {
	gorm.Model
	OpenID string `gorm:"unique"`
	Code   string
	Name   string
	IDCard string
	Level  int
}

//Creator 创建User
func (user *User) Creator() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/9 13:29
	*/
	if err = database.DB.Create(&user).Error; err != nil {
		return
	}
	if err = database.DB.First(&user).Error; err != nil {
		return
	}
	return
}

func TestUser(t *testing.T) {
	database.DB.DropTable("users")
	database.DB.Set("gorm:table_options", "DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;").AutoMigrate(&User{})
	user := &User{Name: "111"}
	log.Println(user)
	err := user.Creator()
	log.Println(err)
	log.Println(user)
}

//func userCreate() (userJson models.UserJSON) {
//	log.Println("userCreate")
//	user := new(models.UserJSON)
//	user.Name = "project_test"
//	if tmp, err := models.UserCreate(user); err != nil {
//		log.Println(err)
//	} else {
//		log.Println(tmp)
//		userJson = tmp
//	}
//	return
//}
//
//func userUpdate(userJson *models.UserJSON) {
//	log.Println("userUpdate")
//	if p, err := models.UserUpdate(userJson); err != nil {
//		log.Println(err)
//	} else {
//		log.Println(p)
//	}
//	return
//}
//
//func userFind() {
//	log.Println("userFind")
//	p := new(models.User)
//	for i := 1; i <= 10; i++ {
//		p.ID = uint(i)
//		if userJson, err := models.UserFind(p); err != nil {
//			log.Println(err)
//		} else {
//			log.Println(userJson)
//		}
//	}
//}
//
//func usersFindByLevel() {
//	log.Println("usersFindByLevel")
//	for i := 1; i <= 10; i++ {
//		if ps, err := models.UsersFindByLevel(i); err != nil {
//			log.Println(err)
//		} else {
//			log.Println(ps)
//		}
//	}
//}
//
//func userDelete() {
//	log.Println("userDelete")
//	u := new(models.User)
//	for i := 1; i <= 5; i++ {
//		u.ID = uint(i)
//		if userJson, err := models.UserDelete(u); err != nil {
//			log.Println(err)
//		} else {
//			log.Println(userJson)
//		}
//	}
//
//}
