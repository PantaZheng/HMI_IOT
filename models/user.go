package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/database"
	"log"
)

const (
	Level_Stranger = iota	// Stranger 未绑定
	Level_Emeritus 			// Professor emeritus 专家教授
	Level_Student			// Student 学生
	Level_Assistant			// Assistant 助理
	Level_Senior			// Senior lecturer 高级讲师
	Level_Full				// Full professor 全职教授
)

type User struct {
	gorm.Model
	OpenId		string		`gorm:"unique;"`
	Code		string
	Name		string
	IDCard		string
	Level		string
	Missions 	[]*Mission	`gorm:"many2many:user_missions"`
}

type UserJson struct {
	ID			uint				`json:"id"`
	OpenId		string     			`json:"openid"`
	Code		string     			`json:"code"`
	Name		string     			`json:"name"`
	IDCard		string				`json:"idCard"`
	Level		string     			`json:"level"`
	Missions	[]*MissionBriefJson	`json:"missions"`
}

type UserBriefJson struct {
	ID		uint	`json:"id"`
	Name	string	`json:"name"`
	Level	string	`json:"level"`
}

func init() {
	database.DB.DropTable("users")
	log.Printf("删除用户表\n")
	database.DB.Set("gorm:table_options", "DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").AutoMigrate(&User{})
	userTest()
}

func userTest(){
	_,_=UserCreate(&User{OpenId: "test1", Level:"unEnrolled"})
	_,_=UserCreate(&User{OpenId: "student1",Name:"student1", Level:"student"})
	_,_=UserCreate(&User{OpenId: "teacher1",Name:"戴国骏", Level:"teacher"})
	_,_=UserCreate(&User{OpenId: "teacher2",Name:"张桦", Level:"teacher"})
	_,_=UserCreate(&User{OpenId: "teacher_unknown",Name:"其他导师", Level:"teacher"})
	log.Printf("user测试")
}

func (user *User) userJson2User(userJson *UserJson){
	user.ID=userJson.ID
	user.OpenId=userJson.OpenId
	user.Name=userJson.Name
	user.IDCard=userJson.IDCard
	user.Level=userJson.Level
	for _,v:=range userJson.Missions{
		recordMission:=new(Mission)
		database.DB.First(&recordMission,v.ID)
	}
}

//登记信息
func UserCreate(user *User)(userBriefJson UserBriefJson,err error){
	recordUser:=User{}
	database.DB.FirstOrCreate(&recordUser,&User{OpenId:user.OpenId})
	database.DB.Model(&recordUser).Updates(user)
	log.Printf("UserCreate\trole:"+user.Level +"\topenid:"+user.OpenId)
	return
}


//根据Role获得成员信息
func GetMembersByRole(role string) ( memberList [] UserBriefJson) {
	var users [] User
	database.DB.Find(&users,&User{Level: role}).Select("id","name")
	memberList=make([]UserBriefJson,len(users))
	for i,v := range users {
		memberList[i].ID=v.ID
		memberList[i].Name=v.Name
	}
	log.Printf("Get:\t"+role+"s\n")
	return
}



