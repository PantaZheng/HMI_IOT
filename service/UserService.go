package service

import (
	"../models"
	"fmt"
	"github.com/chanxuehong/wechat/mp/user"
	"github.com/kataras/iris"
	"log"
	"strconv"
)

func CheckTableUser(){
	models.CheckTableUser()
}

func GetStudents(ctx iris.Context) {
	memberList:=models.GetAllMembers("student")
	if _, err :=ctx.JSON(&memberList);err!=nil {
		log.Printf("GetStudents:\t%v",err)
	}
}

func GetTeachers(ctx iris.Context){
	memberList:=models.GetAllMembers("teacher")
	if _, err :=ctx.JSON(&memberList);err!=nil {
		log.Printf("GetTeachers:\t%v",err)
	}
}

//用户初始化
func UserInit(weChatInfo *user.UserInfo) string {
	if models.RecordUserNotFound(weChatInfo) {
		models.CreateUserByWeChatInfo(weChatInfo)
		fmt.Printf(weChatInfo.OpenId+"新用户关注")
		return "欢迎关注,新用户请进行登记"
	}
	log.Printf(weChatInfo.OpenId+"老用户关注")
	return "欢迎关注,感谢再次关注"
}

//教师信息更新
func UpdateTeacher(ctx iris.Context) {
	teacherInfo:=&models.TeacherInfo{}
	if err:=ctx.ReadJSON(teacherInfo);err!=nil{
		panic(err.Error())
	}
	models.EnrollTeacher(teacherInfo,tagTeacher)
	AddRoleTag(teacherInfo.OpenId,tagTeacher)
	log.Printf(teacherInfo.Name+"教师信息更新tag"+strconv.Itoa(tagTeacher)+"\n")
}

//学生信息更新
func UpdateStudent(ctx iris.Context) {
	studentInfo:=&models.StudentInfo{}
	if err:=ctx.ReadJSON(studentInfo);err!=nil{
		panic(err.Error())
	}
	models.EnrollStudent(studentInfo,tagStudent)

	AddRoleTag(studentInfo.OpenId,tagStudent)
	log.Printf(studentInfo.Name+"同学信息更新tag"+strconv.Itoa(tagStudent)+"\n")
}

//去除Tag
func Purify(ctx iris.Context){
	pureInfo:=&models.PureInfo{}
	if err:=ctx.ReadJSON(pureInfo);err!=nil{
		panic(err.Error())
	}
	tagId:=models.PurifyUser(pureInfo.OpenId)
	DelRoleTag(pureInfo.OpenId,tagId)
	log.Printf("去除用户"+pureInfo.OpenId+"的TagId:"+strconv.Itoa(tagId)+"\n")
}
