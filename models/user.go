package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/database"
	"strconv"
)

const titleUser = "models.user."

//User 数据库用户表.
type User struct {
	gorm.Model
	OpenID     string `gorm:"not null,unique"`
	WechatName string
	Name       string
	IDCard     string `gorm:"not null,unique"`
	Level      int
	Telephone  string
	//CProjects  []*Project `gorm:"foreignkey:CreatorID"`
	//LProjects  []*Project `gorm:"foreignkey:LeaderID"`
	//PProjects  []*Project `gorm:"many2many:user_projects"`
	CModules  []Module  `gorm:"foreignkey:CreatorID"`
	LModules  []Module  `gorm:"foreignkey:LeaderID"`
	PModules  []Module  `gorm:"many2many:user_modules"`
	CMissions []Mission `gorm:"foreignkey:CreatorID"`
	PMissions []Mission `gorm:"many2many:user_missions"`
	OGains    []Gain    `gorm:"foreignkey:OwnerID"`
}

//检查是否有OpenID和IDCard，零值设置为ID,并更新字段信息
func (user *User) makeOpenIDIDCARDNotEmpty() (tag bool) {
	if user.OpenID == "" || user.IDCard == "" {
		if user.OpenID == "" {
			user.OpenID = strconv.Itoa(int(user.ID))
		}
		if user.IDCard == "" {
			user.IDCard = strconv.Itoa(int(user.ID))
		}
		tag = true
	}
	return
}

//Create Create()
func (user *User) Create() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/9 13:29
	*/
	user.ID = 0
	if user.OpenID == "" && user.IDCard == "" {
		err = errors.New("需要OpenID或IDCard来满足用户唯一性")
	} else {
		if err = database.DB.Create(&user).Error; err == nil {
			if user.makeOpenIDIDCARDNotEmpty() {
				err = user.Updates()
			}
		}
	}
	if err != nil {
		err = errors.New(titleUser + "Create:\t" + err.Error())
	}
	return
}

//First 根据id查找用户.
func (user *User) First() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/9 14:02
	*/
	if user.ID > 0 {
		u := User{}
		u.ID = user.ID
		if err = database.DB.First(&u).Error; err == nil {
			*user = u
		}
	} else {
		err = errors.New("ID必须为正数")
	}
	if err != nil {
		err = errors.New(titleUser + "First:\t" + err.Error())
	}
	return
}

//FindOne 单个查找非主键.
func (user *User) FindOne() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 22:22
	*/
	var users []User
	if err = database.DB.Find(&users, &user).Error; err == nil {
		if l := len(users); l > 1 {
			err = errors.New("多个匹配，请确保唯一性")
		} else if l == 0 {
			err = errors.New("没有匹配记录")
		} else {
			*user = users[0]
		}
	}
	if err != nil {
		err = errors.New(titleUser + "FindOne:\t" + err.Error())
	}
	return
}

//Find 查找多个用户.
func (user *User) Find() (users []User, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/9 14:08
	*/
	database.DB.Find(&users)
	if err = database.DB.Find(&users, &user).Error; err == nil {
		if len(users) == 0 {
			err = errors.New("record not found")
		}
	}
	if err != nil {
		err = errors.New(titleUser + "Find:\t" + err.Error())
	}
	return
}

//Updates 非覆盖式更新,零值不更新,根据ID定位用户.
func (user *User) Updates() (err error) {
	/**
	@Author: PantaZheng
	@Description: 非覆盖式更新,零值不更新,根据ID定位用户.注意:Model不可缺
	@Date: 2019/5/9 14:29
	*/
	u := new(User)
	u.ID = user.ID
	if err = database.DB.Model(&u).Updates(&user).Error; err != nil {
		err = errors.New(titleUser + "Updates:\t" + err.Error())
	}
	return
}

//Delete 先将openid和idCard置为id，再软删除.
func (user *User) Delete() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/9 14:36
	*/
	if err = user.FindOne(); err == nil {
		user.OpenID = strconv.Itoa(int(user.ID))
		user.IDCard = strconv.Itoa(int(user.ID))
		if err = user.Updates(); err == nil {
			err = database.DB.Delete(&user).Error
		}
		if err != nil {
			err = errors.New(titleUser + "Delete:\t" + err.Error())
		}
	}
	return
}
