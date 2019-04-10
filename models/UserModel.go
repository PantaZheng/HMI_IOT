package models

import (
	"../database"
	"fmt"
	"github.com/chanxuehong/wechat/mp/user"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	WeChatOpenID string `gorm:"primary_key;unique;VARCHAR(191)"`
	Name         string `gorm:"not null VARCHAR(255)"`
	Sex          string `gorm:"not null VARCHAR"`
	Role         string `gorm:"not null VARCHAR(191)"`
	School       string `gorm:"not null VARCHAR(255)"`
	Supervisor	string `gorm:"not null VARCHAR(191)"`
	HduId        string `gorm:"VARCHAR(191)"`
	Level        string `gorm:"VARCHAR(191)"`
	TagId        int `gorm:"VARCHAR(191)"`

	WeChatAccount string `gorm:"VARCHAR(191)"`
	//WechatNickname string `gorm:"not null VARCHAR(255)"`

	QQ string `gorm:"VARCHAR(191)"`
	Telephone string `gorm:"VARCHAR(191)"`
	Email string `gorm:"VARCHAR(191)"`

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

type MemberInfo struct {
	Id uint `json:"id"`
	Name string `json:"name"`
}

type PureInfo struct{
	WeChatOpenID string `json:"weChatOpenID"`
}

func CheckTableUser() {
	if !database.DB.HasTable(&User{}){
		database.DB.Set("gorm:table_options", "DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").CreateTable(&User{})
		fmt.Printf("新建用户表\n")

	}else{
		fmt.Printf("用户表已存在\n")
	}
}

func DropTableUsers(){
	database.DB.DropTable("users")
	fmt.Printf("删除表\n")
}

func MakeTestData(){
	fmt.Printf("测试数据")
	CreateUser(&user.UserInfo{OpenId:"test1"})
	CreateUser(&user.UserInfo{OpenId:"test2"})
	CreateUser(&user.UserInfo{OpenId:"test3"})
	database.DB.Model(&User{}).Create(
		&User{WeChatOpenID:"student1",Name:"student1",Role:"student",Supervisor:"teacher1"})
	database.DB.Model(&User{}).Create(
		&User{WeChatOpenID:"student2",Name:"student2",Role:"student",Supervisor:"teacher1"})
	database.DB.Model(&User{}).Create(
		&User{WeChatOpenID:"student3",Name:"student3",Role:"student",Supervisor:"teacher2"})
	database.DB.Model(&User{}).Create(
		&User{WeChatOpenID:"teacher1",Name:"戴国骏",Role:"teacher"})
	database.DB.Model(&User{}).Create(
		&User{WeChatOpenID:"teacher2",Name:"张桦",Role:"teacher"})
	database.DB.Model(&User{}).Create(
		&User{WeChatOpenID:"teacher_unknown",Name:"其他导师",Role:"teacher"})
}

//根据WeChatID获取用户
func getUserByWeChatID(weChatOpenID string) (existedUser *User) {
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
func GetAllMembers(role string) ( memberList [] MemberInfo) {
	var users []User
	database.DB.Model(&User{}).Where(&User{Role:role}).Find(&users)
	memberList=make([] MemberInfo,len(users))
	for i,v := range users{
		memberList[i].Id= v.ID
		memberList[i].Name= v.Name
	}
	return memberList
}

func RecordUserNotFound(weChatInfo *user.UserInfo) bool{
	if database.DB.Where("we_chat_open_id=?",weChatInfo.OpenId).Find(&User{}).RecordNotFound(){
		fmt.Printf(weChatInfo.OpenId+"RecordUserNotFound\n")
		return true
	}
	fmt.Printf(weChatInfo.OpenId+"RecordUserFound\n")
	return false
}

//新关注用户创建
func CreateUser(weChatInfo *user.UserInfo){
	anonUser := new(User)
	anonUser.Role = "unEnrolled"
	anonUser.WeChatOpenID = weChatInfo.OpenId
	//anonUser.WechatNickname = weChatInfo.Nickname
	database.DB.Model(&User{}).Create(anonUser)
}

//数据库更新用户信息
func dbUpdateUser(newUser *User) (oldUser *User){
	oldUser = getUserByWeChatID(newUser.WeChatOpenID)
	//newUser.ID=oldUser.ID
	if err := database.DB.Model(&oldUser).Updates(newUser).Error; err != nil {
		fmt.Printf("CreateUserErr:%s", err)
	}
	return
}

//教师登记
func EnrollTeacher(teacherInfo *TeacherInfo, tagId int) {
	teacher := new(User)
	teacher.WeChatOpenID =teacherInfo.WeChatOpenID
	teacher.Role = "teacher"
	teacher.Name = teacherInfo.Name
	teacher.School=teacherInfo.School
	teacher.Sex=teacherInfo.Sex
	teacher.Telephone= teacherInfo.Telephone
	teacher.TagId=tagId
	dbUpdateUser(teacher)
}

//学生登记
func EnrollStudent(studentInfo *StudentInfo, tagId int) {
	student := new(User)
	student.WeChatOpenID=studentInfo.WeChatOpenID
	student.Role = "student"
	student.Name = studentInfo.Name
	student.School= studentInfo.School
	student.Sex= studentInfo.Sex
	student.Telephone= studentInfo.Telephone
	student.Supervisor=studentInfo.Supervisor
	student.TagId = tagId
	dbUpdateUser(student)
}

func PurifyUser(weChatOpenId string)(tagId int){
	Pure := new(User)
	Pure.WeChatOpenID=weChatOpenId
	Pure.Role = "unEnrolled"
	return dbUpdateUser(Pure).TagId
}
