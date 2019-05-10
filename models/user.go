package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/database"
)

//User 数据库用户表.
type User struct {
	gorm.Model
	OpenID     string `gorm:"unique"`
	WechatName string
	Name       string
	IDCard     string `gorm:"unique"`
	Level      int
	Telephone  string
	//CProjects  []*Project `gorm:"foreignkey:CreatorID"`
	//LProjects  []*Project `gorm:"foreignkey:LeaderID"`
	//PProjects  []*Project `gorm:"many2many:user_projects"`
	//CModules   []*Module  `gorm:"foreignkey:CreatorID"`
	//LModules   []*Module  `gorm:"foreignkey:LeaderID"`
	//PModules   []*Module  `gorm:"many2many:user_modules"`
	//CMissions  []*Mission `gorm:"foreignkey:CreatorID"`
	//PMissions  []*Mission `gorm:"many2many:user_missions"`
	//OGains     []*Gain    `gorm:"foreignkey:OwnerID"`
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

//Create 创建User.
func (user *User) Create() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/9 13:29
	*/
	if err = user.checkUnique(); err != nil {
		return
	}
	if err = database.DB.Create(user).Error; err != nil {
		return
	}
	return
}

//first 查找首个用户.
func (user *User) First() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/9 14:02
	*/
	if err = user.checkUnique(); err != nil {
		return
	}
	first := database.DB.First(user)
	err = first.Error
	if first.RecordNotFound() {
		err = errors.New("RecordNotFound")
	}
	return
}

//Find 查找多个用户.
func (user *User) Find() (users []*User, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/9 14:08
	*/
	err = database.DB.Find(users, user).Error
	return
}

//Updates 非覆盖式更新，零值不更新.
func (user *User) Updates(newUser *User) (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/9 14:29
	*/
	err = database.DB.Model(user).Updates(newUser).Error
	return
}

//Save 覆盖式更新，零值更新.
func (user *User) Save(newUser *User) (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/9 14:29
	*/
	err = database.DB.Model(user).Updates(newUser).Error
	return
}

//Delete 先将openid和idCard置为id来实现，再软删除.
func (user *User) Delete() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/9 14:36
	*/
	if err = user.Updates(&User{OpenID: string(user.ID), IDCard: string(user.ID)}); err != nil {
		return
	}
	err = database.DB.Delete(user).Error
	return
}
