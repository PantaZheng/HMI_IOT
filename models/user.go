package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/database"
	"strconv"
)

//User 数据库用户表.
type User struct {
	gorm.Model
	OpenID    string `gorm:"unique"`
	Code      string
	Name      string
	IDCard    string `gorm:"unique"`
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
	user1=&User{OpenID:"test1",Level:1
	}
	_, _ = UserCreate(&UserJSON{OpenID: "test1", Level: LevelEmeritus})
	_, _ = UserCreate(&UserJSON{OpenID: "student1", Name: "student1", Level: LevelStudent})
	_, _ = UserCreate(&UserJSON{OpenID: "assistant1", Name: "assistant1", Level: LevelAssistant})
	_, _ = UserCreate(&UserJSON{OpenID: "full1", Name: "戴国骏", Level: LevelFull})
	_, _ = UserCreate(&UserJSON{OpenID: "senior2", Name: "张桦", Level: LevelSenior})
	_, _ = UserCreate(&UserJSON{OpenID: "teacher_unknown", Name: "其他导师", Level: LevelSenior})
}

func (user *User) checkUnique() (err error) {
	/**
	@Author: PantaZheng
	@Description:检查UserJSON的唯一性要求是否满足，ID,
	OpenID,IDCard
	@Date: 2019/5/9 10:44
	*/
	if user.OpenID == "" && user.ID == 0 && user.IDCard == "" {
		err = errors.New("checkUnique:\t\n需要OpenID或ID或IDCard来满足用户唯一性")
	}
	return
}

//Creator 创建User.
func (user *User) Create() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/9 13:29
	*/
	if err = user.checkUnique(); err != nil {
		return
	}
	if err = database.DB.Create(&user).Error; err != nil {
		return
	}
	return
}

//First 查找首个用户.
func (user *User) First() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/9 14:02
	*/
	if err = user.checkUnique(); err != nil {
		return
	}
	err = database.DB.First(&user).Error
	return
}

//Finds 查找多个用户.
func (user *User) Find() (users []*User, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/9 14:08
	*/
	err = database.DB.Find(&users, &user).Error
	return
}

//Updates 非覆盖式更新，零值不更新.
func (user *User) Updates() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/9 14:29
	*/
	hitUser := User{}
	hitUser = *user
	if err= hitUser.First();err!=nil{
		return
	}
	err = database.DB.Model(&hitUser).Updates(&user).Error
	return
}

//Save 覆盖式更新，零值更新.
func (user *User) Save() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/9 14:29
	*/
	hitUser := User{}
	hitUser = *user
	if err= hitUser.First();err!=nil{
		return
	}
	err = database.DB.Model(&hitUser).Save(&user).Error
	return
}

//Delete 先将openid和idCard置为id来实现，再软删除.
func (user *User) Delete() (err error){
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/9 14:36
	*/
	if err= user.First();err!=nil{
		return
	}
	user.OpenID=string(user.ID)
	user.IDCard=string(user.ID)
	if err=user.Updates();err!=nil{
		return
	}
	err = database.DB.Delete(&user).Error
	return
}


