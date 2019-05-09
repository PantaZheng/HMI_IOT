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

//UserJSON 用户Json原型
type UserJSON struct {
	ID        uint   `json:"id"`
	OpenID    string `json:"openid"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	IDCard    string `json:"idCard"`
	Level     int    `json:"level"`
	Telephone string `json:"telephone"`
}

//UserBriefJSON 简洁版的用户Json信息
type UserBriefJSON struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}

func userTestData() {
	_, _ = UserCreate(&UserJSON{OpenID: "test1", Level: LevelEmeritus})
	_, _ = UserCreate(&UserJSON{OpenID: "student1", Name: "student1", Level: LevelStudent})
	_, _ = UserCreate(&UserJSON{OpenID: "assistant1", Name: "assistant1", Level: LevelAssistant})
	_, _ = UserCreate(&UserJSON{OpenID: "full1", Name: "戴国骏", Level: LevelFull})
	_, _ = UserCreate(&UserJSON{OpenID: "senior2", Name: "张桦", Level: LevelSenior})
	_, _ = UserCreate(&UserJSON{OpenID: "teacher_unknown", Name: "其他导师", Level: LevelSenior})
}

func (user *User) userJSON2User(userJSON *UserJSON) {
	user.ID = userJSON.ID
	user.OpenID = userJSON.OpenID
	user.Name = userJSON.Name
	user.IDCard = userJSON.IDCard
	user.Level = userJSON.Level
	user.Telephone = userJSON.Telephone
}

func (userJson *UserJSON) user2UserJSON(user *User) {
	userJson.ID = user.ID
	userJson.Name = user.Name
	userJson.OpenID = user.OpenID
	userJson.IDCard = user.IDCard
	userJson.Level = user.Level
	userJson.Telephone = user.Telephone
}

//User2UserBriefJSON DB的User转化为BriefJSON
func (userBriefJson *UserBriefJSON) User2UserBriefJSON(user *User) {
	/**
	@Author: PantaZheng
	@Description: 该函数可能会被其他文件引用
	@Date: 2019/5/6 23:15
	*/
	userBriefJson.ID = user.ID
	userBriefJson.Name = user.Name
	userBriefJson.Level = user.Level
	return
}

//UserCreate 模型层面创建用户
func UserCreate(userJSON *UserJSON) (recordUserJSON UserJSON, err error) {
	newUser := new(User)
	newUser.userJSON2User(userJSON)
	if newUser.OpenID == "" {
		tempUser := &User{}
		database.DB.Last(tempUser)
		newUser.OpenID = strconv.Itoa(int(tempUser.ID) + 1)
	}
	if err = database.DB.Create(&newUser).Error; err != nil {
		return
	}
	if err = database.DB.Model(&newUser).First(&newUser).Error; err == nil {
		recordUserJSON.user2UserJSON(newUser)
	}
	return
}

//UserFind 模型层寻找单一用户
func UserFind(user *User) (recordUserJSON UserJSON, err error) {
	recordUser := new(User)
	if err = database.DB.First(&recordUser, &user).Error; err == nil {
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
