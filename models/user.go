package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/database"
	"strconv"
)

//TODO:使用MAP实现
const (
	// LevelStranger 未绑定
	LevelStranger = iota
	// LevelEmeritus Professor emeritus 专家教授
	LevelEmeritus
	// LevelStudent 学生
	LevelStudent
	// LevelAssistant 助理
	LevelAssistant
	// LevelSenior Senior lecturer 高级讲师
	LevelSenior
	// LevelFull Full professor 全职教授
	LevelFull
)

//User 数据库用户表
type User struct {
	gorm.Model
	OpenID    string `gorm:"unique"`
	Code      string
	Name      string
	IDCard    string
	Level     int
	Telephone string
	CProjects []*Project `gorm:"foreignkey:CreatorID"`
	LProjects []*Project `gorm:"foreignkey:LeaderID"`
	PProjects []*Project `gorm:"many2many:user_projects"`
	CModules  []*Module  `gorm:"foreignkey:CreatorID"`
	LModules  []*Module  `gorm:"foreignkey:LeaderID"`
	PModules  []*Module  `gorm:"many2many:user_modules"`
	CMissions []*Mission `gorm:"foreignkey:CreatorID"`
	PMissions []*Mission `gorm:"many2many:user_missions"`
	OGains    []*Gain    `gorm:"foreignkey:OwnerID"`
}

func userTestData() {
	_, _ = UserCreate(&UserJSON{OpenID: "test1", Level: LevelEmeritus})
	_, _ = UserCreate(&UserJSON{OpenID: "student1", Name: "student1", Level: LevelStudent})
	_, _ = UserCreate(&UserJSON{OpenID: "assistant1", Name: "assistant1", Level: LevelAssistant})
	_, _ = UserCreate(&UserJSON{OpenID: "full1", Name: "戴国骏", Level: LevelFull})
	_, _ = UserCreate(&UserJSON{OpenID: "senior2", Name: "张桦", Level: LevelSenior})
	_, _ = UserCreate(&UserJSON{OpenID: "teacher_unknown", Name: "其他导师", Level: LevelSenior})
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

//UserFind 模型层寻找单一用户
func UserFind(user *User) (err error) {
	recordUser := new(User)
	if err = database.DB.First(&user).Error; err == nil {
		recordUserJSON.user2UserJSON(recordUser)
	}
	return
}

//UsersFindByLevel 根据level寻找多用户
func UsersFindByLevel(level int) (usersBriefJSON []UserBriefJSON, err error) {
	users := make([]User, 1)
	if database.DB.Find(&users, &User{Level: level}).RecordNotFound() {
		err = errors.New("ProjectsFindByLeader No Project Record")
	} else {
		for _, v := range users {
			tempJSON := &UserBriefJSON{}
			tempJSON.User2UserBriefJSON(&v)
			usersBriefJSON = append(usersBriefJSON, *tempJSON)
		}
	}
	return
}

//UserUpdate 更新用户信息
func UserUpdate(userJSON *UserJSON) (recordUserJSON UserJSON, err error) {
	updateUser := new(User)
	updateUser.userJSON2User(userJSON)
	recordUser := new(User)
	recordUser.ID = updateUser.ID
	if database.DB.First(&recordUser, &recordUser).RecordNotFound() {
		err = errors.New("UserUpdate No User Record")
	} else {
		database.DB.Model(&recordUser).Updates(updateUser)
		recordUserJSON.user2UserJSON(recordUser)
	}
	return
}

//UserDelete 模型层删除用户信息
func UserDelete(user *User) (recordUserJSON UserJSON, err error) {
	recordUser := new(User)
	if database.DB.First(&recordUser, &user).RecordNotFound() {
		err = errors.New("UserDelete No User Record")
	} else {
		recordUserJSON.user2UserJSON(recordUser)
		err = database.DB.Unscoped().Delete(&recordUser).Error
	}
	return
}
