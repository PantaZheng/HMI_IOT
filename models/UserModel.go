package models

import (
	"../database"
	"fmt"
	"github.com/chanxuehong/wechat/mp/user"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	WechatOpenID string `gorm:"unique;VARCHAR(191)"`
	Name  string `gorm:"not null VARCHAR(191)"`
	Sex string `gorm:"not null VARCHAR"`
	Role 	string `gorm:"not null VARCHAR(191)"`
	School string `gorm:"not null VARCHAR(191)"`
	HduId	string `gorm:"unique;VARCHAR(191)"`
	Level string `gorm:"VARCHAR(191)"`

	WechatAccount string `gorm:"unique;VARCHAR(191)"`
	WechatNickname string `gorm:"not null VARCHAR(191)"`

	QQ string `gorm:"unique;VARCHAR(191)"`
	Telephone string `gorm:"unique;VARCHAR(191)"`
	Email string `gorm:"unique;VARCHAR(191)"`

	EduStartDate string `gorm:"VARCHAR(191)"`
	Graduate string `gorm:"VARCHAR(191)"`
}

type TeacherInfo struct{
	Name string `json:"name"`
	Sex string `json:"sex"`
	School string `json:"school"`
	Telephone string `json:"tel"`
}

type StudentInfo struct{
	Name string `json:"name"`
	Sex string `json:"sex"`
	Teacher string `json:"teacher"`
	School string `json:"school"`
	Telephone string `json:"tel"`
}

func CheckTableUser() {
	if errCreate := database.DB.Create(User{});errCreate!=nil{
			fmt.Printf("createTable:%v", errCreate)
	}
}

//检查是否已经存在用户
func CheckUserByWeChatID(weChatOpenID string) bool {
	usr := new(User)
	usr.WechatOpenID=weChatOpenID
	return database.DB.First(usr).RecordNotFound()
}

//根据WeChatID获取用户
func GetUserByWechatID(weChatOpenID string) (existedUser *User) {
	existedUser = new(User)
	existedUser.WechatOpenID=weChatOpenID
	if err := database.DB.First(existedUser).Error; err != nil {
		fmt.Printf("GetUserByIdErr:%s", err)
	}
	return
}

//根据WeChatID获取用户身份
func GetUserRoleByWechatID(weChatOpenID string) string{
	existedUser := new(User)
	existedUser.WechatOpenID=weChatOpenID
	if err := database.DB.First(existedUser).Error; err != nil {
		fmt.Printf("GetUserByIdErr:%s", err)
	}
	return existedUser.Role
}

//新关注用户创建
func CreateUser(weChatInfo *user.UserInfo){
	anonUser := new(User)
	anonUser.Role = "unEnrolled"
	anonUser.WechatOpenID = weChatInfo.OpenId
	anonUser.WechatNickname = weChatInfo.Nickname

	if err := database.DB.Create(anonUser).Error; err != nil {
		fmt.Printf("CreateUserErr:%s", err)
	}
}

//教师登记
func EnrollTeacher(weChatOpenID string, teacherInfo *TeacherInfo) {
	teacher := new(User)
	teacher.Role = "teacher"
	teacher.Name = teacherInfo.Name
	teacher.School=teacherInfo.School
	teacher.Sex=teacherInfo.Sex
	teacher.Telephone= teacherInfo.Telephone

	oldUser := GetUserByWechatID(weChatOpenID)

	if err := database.DB.Model(oldUser).Updates(teacher).Error; err != nil {
		fmt.Printf("CreateUserErr:%s", err)
	}
}

//学生登记
func EnrollStudent(weChatOpenID string, studentInfo *StudentInfo) {
	student := new(User)
	student.Role = "student"
	student.Name = studentInfo.Name
	student.School= studentInfo.School
	student.Sex= studentInfo.Sex
	student.Telephone= studentInfo.Telephone

	oldUser := GetUserByWechatID(weChatOpenID)

	if err := database.DB.Model(oldUser).Updates(student).Error; err != nil {
		fmt.Printf("CreateUserErr:%s", err)
	}
}

