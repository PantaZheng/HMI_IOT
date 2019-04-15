package models

import (
	"../database"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type User struct {
	gorm.Model
	OpenId string `gorm:"primary_key;unique;VARCHAR(191)" json:"openid"`
	Code string `gorm:"not null VARCHAR(255)" json:"code"`
	Name         string `gorm:"not null VARCHAR(255)" json:"name"`
	Sex          string `gorm:"not null VARCHAR" json:"sex"`
	Role         string `gorm:"not null VARCHAR(191)" json:"role"`
	School       string `gorm:"not null VARCHAR(255)" json:"school"`
	Supervisor	string `gorm:"not null VARCHAR(191)" json:"supervisor"`
	//HduId        string `gorm:"VARCHAR(191)"`
	//Level        string `gorm:"VARCHAR(191)"`
	//TagId        int `gorm:"VARCHAR(191)"`
	//WeChatAccount string `gorm:"VARCHAR(191)"`
	//WechatNickname string `gorm:"not null VARCHAR(255)"`
	//QQ string `gorm:"VARCHAR(191)"`
	//Telephone string `gorm:"VARCHAR(191)"`
	//Email string `gorm:"VARCHAR(191)"`
	//EduStartDate string `gorm:"VARCHAR(191)"`
	//Graduate string `gorm:"VARCHAR(191)"`
}

type MemberInfo struct {
	Id uint `json:"id"`
	Name string `json:"name"`
}

func CheckTableUser() {
	//if !database.DB.HasTable(&User{}){
	//	database.DB.Set("gorm:table_options", "DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").CreateTable(&User{})
	//	log.Printf("新建用户表\n")
	//	MakeTestData()
	//}else{
	//	log.Printf("用户表已存在\n")
	//}
	database.DB.Set("gorm:table_options", "DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").AutoMigrate(&User{})
	MakeTestData()
}

func DropTableUsers(){
	database.DB.DropTable("users")
	log.Printf("删除用户表\n")
}

func MakeTestData(){
	_= EnrollUser(&User{OpenId: "test1",Role:"unEnrolled"})
	_= EnrollUser(&User{OpenId: "test2",Role:"unEnrolled"})
	_= EnrollUser(&User{OpenId: "test3",Role:"unEnrolled"})
	_= EnrollUser(&User{OpenId: "student1",Name:"student1",Role:"student",Supervisor:"teacher1"})
	_= EnrollUser(&User{OpenId: "student2",Name:"student2",Role:"student",Supervisor:"teacher1"})
	_= EnrollUser(&User{OpenId: "student3",Name:"student3",Role:"student",Supervisor:"teacher2"})
	_= EnrollUser(&User{OpenId: "teacher1",Name:"戴国骏",Role:"teacher"})
	_= EnrollUser(&User{OpenId: "teacher2",Name:"张桦",Role:"teacher"})
	_= EnrollUser(&User{OpenId: "teacher_unknown",Name:"其他导师",Role:"teacher"})
	log.Printf("创建测试用户数据")
}

//根据WeChatID获取用户
func GetUser(openid string) (user *User,err error) {
	user = new(User)
	user.OpenId =openid
	if err = database.DB.Where(&User{OpenId:openid}).First (user).Error; err != nil {
		fmt.Printf("GetUserByIdErr:%s", err)
	}
	return
}

//根据Role获得成员信息
func GetMembersByRole(role string) ( memberList [] MemberInfo) {
	var users []User
	database.DB.Model(&User{}).Where(&User{Role:role}).Find(&users)
	memberList=make([] MemberInfo,len(users))
	for i,v := range users{
		memberList[i].Id= v.ID
		memberList[i].Name= v.Name
	}
	log.Printf("GetAllMembers,role:\t"+role+"\n")
	return memberList
}

func recordNotFound(openid string) bool{
	if database.DB.Where("open_id=?",openid).Find(&User{}).RecordNotFound(){
		log.Printf("RecordUserNotFound\t"+openid+"\n")
		return true
	}
	log.Printf("RecordUserFound:\t"+openid+"\n")
	return false
}

//数据库创建用户
func dbCreateUser(newUser *User)(){
	database.DB.Model(&User{}).Create(newUser)
	//log.Printf("dbCreateUser:\t"+ newUser.OpenId)
}

//数据库更新用户信息
func dbUpdateUser(newUser *User) (err error){
	oldUser:=&User{}
	if oldUser,err = GetUser(newUser.OpenId);err!=nil{
		return err
	}
	if err = database.DB.Model(&User{}).Where(&User{OpenId:oldUser.OpenId}).Updates(newUser).Error; err != nil {
		return err
	}
	//log.Printf("dbUpdateUser:\t"+oldUser.OpenId)
	return
}

//登记信息
func EnrollUser( user *User) (err error){
	if recordNotFound(user.OpenId){
		dbCreateUser(user)
	}else{
		if err=dbUpdateUser(user);err!=nil{
			return
		}
	}
	log.Printf("EnrollUser\trole:"+user.Role+"\topenid:"+user.OpenId)
	return
}

