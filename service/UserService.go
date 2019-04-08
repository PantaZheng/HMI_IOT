package service

import (
	"../models"
	"fmt"
	"github.com/chanxuehong/wechat/mp/user"
	"github.com/kataras/iris"
)

func CheckTableUser(){
	models.CheckTableUser()
	models.MakeTestData()
}

func GetStudents(ctx iris.Context) {
	memberList:=models.GetAllMembers("student")
	if _, err :=ctx.JSON(&memberList);err!=nil {
		fmt.Printf("GetStudents:%v",err)
	}
}

func GetTeachers(ctx iris.Context){
	memberList:=models.GetAllMembers("teacher")
	if _, err :=ctx.JSON(&memberList);err!=nil {
		fmt.Printf("GetTeachers:%v",err)
	}
}

//用户初始化
func UserInit(weChatInfo *user.UserInfo) string {
		models.CreateUser(weChatInfo)
	fmt.Printf(weChatInfo.OpenId+"用户关注")
	return "欢迎关注,新用户请进行登记"
}

//教师信息更新
func UpdateTeacher(ctx iris.Context) {
	teacherInfo:=&models.TeacherInfo{}
	if err:=ctx.ReadJSON(teacherInfo);err!=nil{
		panic(err.Error())
	}
	models.EnrollTeacher(teacherInfo)
	fmt.Printf(teacherInfo.Name+"教师信息更新")
}

//学生信息更新
func UpdateStudent(ctx iris.Context) {
	studentInfo:=&models.StudentInfo{}
	if err:=ctx.ReadJSON(studentInfo);err!=nil{
		panic(err.Error())
	}
	models.EnrollStudent(studentInfo)
	fmt.Printf(studentInfo.Name+"同学信息更新")
}

