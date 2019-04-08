package models

import (
	"../database"
	"fmt"
	"github.com/chanxuehong/wechat/mp/user"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	WeChatOpenID string `gorm:"unique;VARCHAR(191)"`
	Name         string `gorm:"not null VARCHAR(191)"`
	Sex          string `gorm:"not null VARCHAR"`
	Role         string `gorm:"not null VARCHAR(191)"`
	School       string `gorm:"not null VARCHAR(191)"`
	Supervisor	string `gorm:"not null VARCHAR(191)"`
	HduId        string `gorm:"unique;VARCHAR(191)"`
	Level        string `gorm:"VARCHAR(191)"`
	TagID        string `gorm:"VARCHAR(191)"`

	WeChatAccount string `gorm:"unique;VARCHAR(191)"`
	//WechatNickname string `gorm:"not null VARCHAR(255)"`

	QQ string `gorm:"unique;VARCHAR(191)"`
	Telephone string `gorm:"unique;VARCHAR(191)"`
	Email string `gorm:"unique;VARCHAR(191)"`

	EduStartDate string `gorm:"VARCHAR(191)"`
	Graduate string `gorm:"VARCHAR(191)"`
}

type TeacherInfo struct{
	WeChatOpenID string `json:"weChatOpenID"`
	Name string `json:"name"`
	Sex string `json:"sex"`
	School string `json:"school"`
	Telephone string `json:"telephone"`
}

type StudentInfo struct{
	WeChatOpenID string `json:"weChatOpenID"`
	Name string `json:"name"`
	Sex string `json:"sex"`
	Supervisor string `json:"supervisor"`
	School string `json:"school"`
	Telephone string `json:"telephone"`
}

func CheckTableUser() {
	if !database.DB.HasTable(User{}){
		database.DB.CreateTable(User{})
		fmt.Printf("新建用户表")
	}else{
		fmt.Printf("用户表已存在")
	}
}

func MakeTestData(){
	database.DB.Model(&User{}).FirstOrCreate(
		&User{WeChatOpenID:"student1",Name:"student1",Role:"student",Supervisor:"teacher1"})
	database.DB.Model(&User{}).FirstOrCreate(
		&User{WeChatOpenID:"student2",Name:"student2",Role:"student",Supervisor:"teacher1"})
	database.DB.Model(&User{}).FirstOrCreate(
		&User{WeChatOpenID:"student3",Name:"student3",Role:"student",Supervisor:"teacher2"})
	database.DB.Model(&User{}).FirstOrCreate(
		&User{WeChatOpenID:"teacher1",Name:"teacher1",Role:"teacher"})
	database.DB.Model(&User{}).FirstOrCreate(
		&User{WeChatOpenID:"teacher2",Name:"teacher2",Role:"teacher"})
}

//根据WeChatID获取用户
func GetUserByWeChatID(weChatOpenID string) (existedUser *User) {
	existedUser = new(User)
	existedUser.WeChatOpenID =weChatOpenID
	if err := database.DB.Model((&User{})).First (existedUser).Error; err != nil {
		fmt.Printf("GetUserByIdErr:%s", err)
	}
	return
}


//根据WeChatID获取用户身份
func GetUserRoleByWeChatID(weChatOpenID string) string{
	existedUser := new(User)
	existedUser.WeChatOpenID =weChatOpenID
	if err := database.DB.Model((&User{})).First(existedUser).Error; err != nil {
		fmt.Printf("GetUserByIdErr:%s", err)
	}
	return existedUser.Role
}

//根据Role获得成员信息
func GetAllMembers(role string) ( memberList [] string) {
	var users []User
	database.DB.Model(&User{}).Where(&User{Role:role}).Find(&users)
	memberList=make([]string,len(users))
	for i,v := range users{
		memberList[i]= v.Name
	}
	return memberList
}

//新关注用户创建
func CreateUser(weChatInfo *user.UserInfo){
	anonUser := new(User)
	anonUser.Role = "unEnrolled"
	anonUser.WeChatOpenID = weChatInfo.OpenId
	//anonUser.WechatNickname = weChatInfo.Nickname

	database.DB.FirstOrCreate(anonUser)
}

//教师登记
func EnrollTeacher(teacherInfo *TeacherInfo) {
	teacher := new(User)
	teacher.WeChatOpenID =teacherInfo.WeChatOpenID
	teacher.Role = "teacher"
	teacher.Name = teacherInfo.Name
	teacher.School=teacherInfo.School
	teacher.Sex=teacherInfo.Sex
	teacher.Telephone= teacherInfo.Telephone

	oldUser := GetUserByWeChatID(teacher.WeChatOpenID)

	if err := database.DB.Model(oldUser).Updates(teacher).Error; err != nil {
		fmt.Printf("CreateUserErr:%s", err)
	}
}

//学生登记
func EnrollStudent(studentInfo *StudentInfo) {
	student := new(User)
	student.WeChatOpenID=studentInfo.WeChatOpenID
	student.Role = "student"
	student.Name = studentInfo.Name
	student.School= studentInfo.School
	student.Sex= studentInfo.Sex
	student.Telephone= studentInfo.Telephone
	student.Supervisor=studentInfo.Supervisor

	oldUser := GetUserByWeChatID(student.WeChatOpenID)

	if err := database.DB.Model(oldUser).Updates(student).Error; err != nil {
		fmt.Printf("CreateUserErr:%s", err)
	}
}

