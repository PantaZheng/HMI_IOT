package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/database"
	"log"
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
	OpenId    	string		`gorm:"unique;"`
	Code      	string
	Name      	string
	IDCard    	string
	Level     	int
	LProjects	[]*Project `gorm:"foreignkey:LeaderID"`
	PProjects	[]*Project `gorm:"many2many:user_projects"`
	LModules	[]*Module   `gorm:"foreignkey:LeaderID"`
	PModules 	[]*Module   `gorm:"many2many:user_modules"`
	PMissions 	[]*Mission `gorm:"many2many:user_missions"`
	PAchieves	[]*Gain    `gorm:"foreignkey:OwnerID"`
}

type UserJson struct {
	ID			uint				`json:"id"`
	OpenId		string     			`json:"openid"`
	Code		string     			`json:"code"`
	Name		string     			`json:"name"`
	IDCard		string				`json:"id_card"`
	Level		int     			`json:"level"`
	//Missions	[]*MissionBriefJson	`json:"missions"`
}

type UserBriefJson struct {
	ID		uint	`json:"id"`
	Name	string	`json:"name"`
	Level	int	`json:"level"`
}

func userTestData(){
	_,_=UserCreate(&User{OpenId: "test1", Level: LevelEmeritus})
	_,_=UserCreate(&User{OpenId: "student1",Name:"student1", Level: LevelStudent})
	_,_=UserCreate(&User{OpenId: "teacher1",Name:"戴国骏", Level: LevelFull})
	_,_=UserCreate(&User{OpenId: "teacher2",Name:"张桦", Level:LevelSenior})
	_,_=UserCreate(&User{OpenId: "teacher_unknown",Name:"其他导师", Level:LevelSenior})
	log.Printf("user测试")
}

func (user *User) userJson2User(userJson *UserJson){
	user.ID=userJson.ID
	user.OpenId=userJson.OpenId
	user.Name=userJson.Name
	user.IDCard=userJson.IDCard
	user.Level=userJson.Level
	//for _,v:=range userJson.Missions{
	//	recordMission:=new(Mission)
	//	database.DB.First(&recordMission,v.ID)
	//}
}

func (userBriefJson *UserBriefJson) user2UserBriefJson(user *User){
	userBriefJson.ID=user.ID
	userBriefJson.Name=user.Name
	userBriefJson.Level=user.Level
}

//登记信息
func UserCreate(user *User)(userBriefJson UserBriefJson,err error){
	recordUser:=User{}
	database.DB.FirstOrCreate(&recordUser,&User{OpenId:user.OpenId})
	database.DB.Model(&recordUser).Updates(user)
	log.Printf("UserCreate\trole:"+strconv.Itoa(user.Level) +"\topenid:"+user.OpenId)
	return
}


//根据Role获得成员信息
func GetMembersByRole(level int) ( memberList [] UserBriefJson) {
	var users [] User
	database.DB.Find(&users,&User{Level: level}).Select("id","name")
	memberList=make([]UserBriefJson,len(users))
	for i,v := range users {
		memberList[i].ID=v.ID
		memberList[i].Name=v.Name
	}
	log.Printf("Get:\t"+strconv.Itoa(level)+"s\n")
	return
}



