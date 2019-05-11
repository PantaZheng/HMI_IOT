package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/database"
	"log"
	"strconv"
)

const title = "models.user."

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

func (user *User) makeOpenIDIDCARDNotEmpty() (tag bool) {
	//检查是否有OpenID和IDCard，零值设置为ID,并更新字段信息
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
	log.Println(user)
	if user.makeOpenIDIDCARDNotEmpty() {
		log.Println(user)
		if err = database.DB.Updates(user).Error; err != nil {
			return
		}
	}
	return
}

//first 根据id查找用户.
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
	return
}

func (user *User) FindOne() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 22:22
	*/
	var users []User
	find := database.DB.Find(&users, &user)
	err = find.Error
	if err == nil {
		if l := len(users); l > 1 {
			err = errors.New("多个匹配，请确保唯一性")
		} else if l == 0 {
			err = errors.New("没有匹配记录")
		} else {
			*user = users[0]
		}
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
	err = database.DB.Find(&users, &user).Error
	return
}

//Updates 非覆盖式更新，零值不更新.根据ID更新
func (user *User) Updates() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/9 14:29
	*/
	u := new(User)
	u.ID = user.ID
	err = database.DB.Model(u).Updates(user).Error
	return
}

//Delete 先将openid和idCard置为id，再软删除.
func (user *User) Delete() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/9 14:36
	*/
	user.OpenID = strconv.Itoa(int(user.ID))
	user.IDCard = strconv.Itoa(int(user.ID))
	if err = user.Updates(); err != nil {
		return
	}
	err = database.DB.Delete(user).Error
	return
}
