package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/pantazheng/bci/database"
)

const titleUser = "models.user."

type UserCore struct {
	ID   uint   `gorm:"primary_key",json:"id"`
	Name string `json:"name"`
}

//User 数据库用户表.
type User struct {
	UserCore
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	OpenID     string `gorm:"unique"`
	WechatName string

	Level     int
	Telephone string `gorm:"unique"`
}

//func userTestData() {
//	log.Println("userTestData")
//	users := make([]User, 11)
//	users[0] = User{OpenID: "Stranger1", WechatName: "小蜘蛛", Level: 1, Telephone: "110"}
//	users[0].Name = "Peter Benjamin Parker"
//	users[1] = User{OpenID: "Emeritus1", WechatName: "万磁王", Level: 2}
//	users[1].Name = "Max Eisenhardt"
//	users[2] = User{WechatName: "金刚狼", Level: 2, Telephone: "111"}
//	users[2].Name = "Logan Howlett"
//	users[3] = User{OpenID: "Assistant1", WechatName: "小辣椒", Level: 3}
//	users[3].Name = "Pepper Potts"
//	users[4] = User{WechatName: "钢铁侠", Level: 4, Telephone: "112"}
//	users[4].Name = "Tony Stark"
//	users[5] = User{OpenID: "Full1", WechatName: "灭霸", Level: 5}
//	users[5].Name = "Thanos"
//	users[6] = User{Level: 3, Telephone: "18955537316"}
//	users[6].Name = "韩新亚"
//	users[7] = User{Level: 3, Telephone: "13867188664"}
//	users[7].Name = "曾虹"
//	users[8] = User{Level: 5, Telephone: "13906524548"}
//	users[8].Name = "戴国骏"
//	users[9] = User{Level: 3, Telephone: "13336096310"}
//	users[9].Name = "周文晖"
//	users[10] = User{Level: 3, Telephone: "13777840698"}
//	users[10].Name = "张桦"
//	for _, v := range users {
//		if err := v.Insert(); err != nil {
//			log.Println(err.Error())
//		} else {
//			log.Println(v)
//		}
//	}
//}

//检查是否有OpenID和IDCard，零值设置为ID,并更新字段信息
func (user *User) makeOpenIDTelephoneNotEmpty() (tag bool) {
	if user.OpenID == "" || user.Telephone == "" {
		if user.OpenID == "" {
			user.OpenID = strconv.Itoa(int(user.ID))
		}
		if user.Telephone == "" {
			user.Telephone = strconv.Itoa(int(user.ID))
		}
		tag = true
	}
	return
}

//Insert Insert()
func (user *User) Insert() (err error) {
	user.ID = 0
	if user.OpenID == "" && user.Telephone == "" {
		err = errors.New("需要OpenID或Telephone来满足用户唯一性")
		return
	}
	if err = database.DB.Create(&user).Error; err != nil {
		return
	}
	if user.makeOpenIDTelephoneNotEmpty() {
		err = user.Updates()
	}
	return
}

//First 根据id查找用户.
func (user *User) First() (err error) {
	err = database.DB.Where("id=?", user.ID).First(&user).Error
	return
}

func (userCore *UserCore) First() (err error) {
	u := User{}
	u.ID = userCore.ID
	if err = u.First(); err != nil {
		return
	} else {
		*userCore = u.UserCore
	}
	return
}

//FindOne 单个查找非主键.
func (user *User) FindOne() (err error) {
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

//FindByField 查找多个用户.
func (user *User) Find() (users []User, err error) {
	users = make([]User, 0)
	if err = database.DB.Find(&users, &user).Error; err == nil {
		if len(users) == 0 {
			err = errors.New("record not found")
		}
	}
	if err != nil {
		err = errors.New(titleUser + "FindByField:\t" + err.Error())
	}
	return
}

//Updates 非覆盖式更新,零值不更新,根据ID定位用户.
func (user *User) Updates() (err error) {
	if err = database.DB.Model(User{}).Where("id=?", user.ID).Updates(&user).Error; err != nil {
		return
	}
	err = user.First()
	return
}

//Delete 先将openid和idCard置为id，再软删除.
func (user *User) Delete() (err error) {
	if err = user.FindOne(); err == nil {
		if user.Level > 1 {
			user.OpenID = strconv.Itoa(int(user.ID))
			err = user.Updates()
		} else {
			err = database.DB.Unscoped().Delete(&user).Error
		}
	}
	if err != nil {
		err = errors.New(titleUser + "Delete:\t" + err.Error())
	}
	return
}
