package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/database"
	"strconv"
)

const (
	LevelStranger         = iota // Stranger 未绑定
	LevelEmeritus                // Professor emeritus 专家教授
	LevelStudent                 // Student 学生
	LevelAssistant              // Assistant 助理
	LevelSenior                 // Senior lecturer 高级讲师
	LevelFull                   // Full professor 全职教授
)

type User struct {
	gorm.Model
	OpenId		string		`gorm:"unique"`
	Code		string
	Name		string
	IDCard		string
	Level		int
	Telephone	string
	LProjects	[]*Project	`gorm:"foreignkey:LeaderID"`
	PProjects	[]*Project	`gorm:"many2many:user_projects"`
	LModules	[]*Module	`gorm:"foreignkey:LeaderID"`
	PModules	[]*Module	`gorm:"many2many:user_modules"`
	PMissions	[]*Mission	`gorm:"many2many:user_missions"`
	PAchieves	[]*Gain		`gorm:"foreignkey:OwnerID"`
}

type UserJson struct {
	ID			uint		`json:"id"`
	OpenId		string		`json:"openid"`
	Code		string		`json:"code"`
	Name		string		`json:"name"`
	IDCard		string		`json:"id_card"`
	Level		int			`json:"level"`
	Telephone	string		`json:"telephone"`
}

type UserBriefJson struct {
	ID		uint	`json:"id"`
	Name	string	`json:"name"`
	Level	int		`json:"level"`
}

func userTestData(){
	_,_=UserCreate(&UserJson{OpenId: "test1", Level: LevelEmeritus})
	_,_=UserCreate(&UserJson{OpenId: "student1",Name:"student1", Level: LevelStudent})
	_,_=UserCreate(&UserJson{OpenId: "assistant1",Name:"assistant1",Level:LevelAssistant})
	_,_=UserCreate(&UserJson{OpenId: "full1",Name:"戴国骏", Level: LevelFull})
	_,_=UserCreate(&UserJson{OpenId: "senior2",Name:"张桦", Level:LevelSenior})
	_,_=UserCreate(&UserJson{OpenId: "teacher_unknown",Name:"其他导师", Level:LevelSenior})
}

func (user *User) userJson2User(userJson *UserJson){
	user.ID=userJson.ID
	user.OpenId=userJson.OpenId
	user.Name=userJson.Name
	user.IDCard=userJson.IDCard
	user.Level=userJson.Level
	user.Telephone=userJson.Telephone
}

func (userJson *UserJson) user2UserJson(user *User){
	userJson.ID=user.ID
	userJson.Name=user.Name
	userJson.OpenId=user.OpenId
	userJson.IDCard=user.IDCard
	userJson.Level=user.Level
	userJson.Telephone=user.Telephone
}

func (userBriefJson *UserBriefJson) user2UserBriefJson(user *User){
	userBriefJson.ID=user.ID
	userBriefJson.Name=user.Name
	userBriefJson.Level=user.Level
}

//创建用户
func UserCreate(userJson *UserJson)(recordUserJson UserJson,err error){
	newUser:=new(User)
	newUser.userJson2User(userJson)
	if newUser.OpenId==""{
		tempUser:=&User{}
		database.DB.Last(tempUser)
		newUser.OpenId=strconv.Itoa(int(tempUser.ID)+1)
	}
	if err=database.DB.Create(&newUser).Error;err!=nil{
		return
	}
	if err=database.DB.Model(&newUser).First(&newUser).Error;err==nil{
		recordUserJson.user2UserJson(newUser)
	}
	return
}

func UserFind(user *User)(recordUserJson UserJson,err error){
	recordUser:=new(User)
	if err=database.DB.First(&recordUser,&user).Error;err==nil{
		recordUserJson.user2UserJson(recordUser)
	}
	return
}

func UsersFindByLevel(level int)(usersBriefJson []UserBriefJson,err error){
	users:=make([]User,1)
	if database.DB.Find(&users,&User{Level:level}).RecordNotFound(){
		err=errors.New("ProjectsFindByLeader No Project Record")
	}else{
		for _,v:=range users{
			tempJson:=&UserBriefJson{}
			tempJson.user2UserBriefJson(&v)
			usersBriefJson=append(usersBriefJson,*tempJson)
		}
	}
	return
}

func UserUpdate(userJson *UserJson)(recordUserJson UserJson,err error){
	updateUser:=new(User)
	updateUser.userJson2User(userJson)
	recordUser:=new(User)
	recordUser.ID=updateUser.ID
	if database.DB.First(&recordUser,&recordUser).RecordNotFound(){
		err = errors.New("UserUpdate No User Record")
	}else{
		database.DB.Model(&recordUser).Updates(updateUser)
		recordUserJson.user2UserJson(recordUser)
	}
	return
}

func UserDelete(user *User)(recordUserJson UserJson, err error){
	recordUser:=new(User)
	if database.DB.First(&recordUser,&user).RecordNotFound(){
		err=errors.New("UserDelete No User Record")
	}else{
		recordUserJson.user2UserJson(recordUser)
		err=database.DB.Unscoped().Delete(&recordUser).Error
	}
	return
}




