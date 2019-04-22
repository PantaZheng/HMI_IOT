package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/database"
	"log"
)

//TODO:导师和学生关系
type User struct {
	gorm.Model
	OpenId   string     `gorm:"unique;" json:"openid"`
	Code     string     `json:"code"`
	Name     string     `json:"name"`
	Level    string     `json:"Level"`
	Missions []*Mission `gorm:"many2many:user_missions"`
}

type MemberInfo struct {
	Id uint `json:"id"`
	Name string `json:"name"`
}

func init() {
	database.DB.DropTable("users")
	log.Printf("删除用户表\n")
	database.DB.Set("gorm:table_options", "DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").AutoMigrate(&User{})
	MakeTestData()
}

func MakeTestData(){
	EnrollUser(&User{OpenId: "test1", Level:"unEnrolled"})
	EnrollUser(&User{OpenId: "test2", Level:"unEnrolled"})
	EnrollUser(&User{OpenId: "test3", Level:"unEnrolled"})
	EnrollUser(&User{OpenId: "student1",Name:"student1", Level:"student"})
	EnrollUser(&User{OpenId: "student2",Name:"student2", Level:"student"})
	EnrollUser(&User{OpenId: "student3",Name:"student3", Level:"student"})
	EnrollUser(&User{OpenId: "teacher1",Name:"戴国骏", Level:"teacher"})
	EnrollUser(&User{OpenId: "teacher2",Name:"张桦", Level:"teacher"})
	EnrollUser(&User{OpenId: "teacher_unknown",Name:"其他导师", Level:"teacher"})
	log.Printf("创建测试用户数据")
}


//根据Role获得成员信息
func GetMembersByRole(role string) ( memberList [] MemberInfo) {
	var users [] User
	database.DB.Find(&users,&User{Level: role}).Select("id","name")
	memberList=make([]MemberInfo,len(users))
	for i,v := range users {
		memberList[i].Id=v.ID
		memberList[i].Name=v.Name
	}
	log.Printf("Get:\t"+role+"s\n")
	return
}

//登记信息
func EnrollUser(user *User){
	recordUser:=User{}
	database.DB.FirstOrCreate(&recordUser,&User{OpenId:user.OpenId})
	database.DB.Model(&recordUser).Updates(user)
	log.Printf("EnrollUser\trole:"+user.Level +"\topenid:"+user.OpenId)
}

