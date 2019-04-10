package service

import (
	"../models"
	"fmt"
	"github.com/chanxuehong/wechat/mp/user"
	"github.com/kataras/iris"
	"strconv"
)

func CheckTableUser(){
	models.DropTableUsers()
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
	if models.RecordUserNotFound(weChatInfo) {
		models.CreateUser(weChatInfo)
		fmt.Printf(weChatInfo.OpenId+"新用户关注")
		return "欢迎关注,新用户请进行登记"
	}
	fmt.Printf(weChatInfo.OpenId+"老用户关注")
	return "欢迎关注,感谢再次关注"
}

//教师信息更新
func UpdateTeacher(ctx iris.Context) {
	teacherInfo:=&models.TeacherInfo{}
	if err:=ctx.ReadJSON(teacherInfo);err!=nil{
		panic(err.Error())
	}
	models.EnrollTeacher(teacherInfo,tagTeacher)
	AddRoleTag([]string{teacherInfo.WeChatOpenID},tagTeacher)
	fmt.Printf(teacherInfo.Name+"教师信息更新tag"+strconv.Itoa(tagTeacher)+"\n")
}

//学生信息更新
func UpdateStudent(ctx iris.Context) {
	studentInfo:=&models.StudentInfo{}
	if err:=ctx.ReadJSON(studentInfo);err!=nil{
		panic(err.Error())
	}
	models.EnrollStudent(studentInfo,tagStudent)
	AddRoleTag([]string{studentInfo.WeChatOpenID},tagStudent)
	fmt.Printf(studentInfo.Name+"同学信息更新tag"+strconv.Itoa(tagStudent)+"\n")
}

//去除Tag
func Purify(ctx iris.Context){
	pureInfo:=&models.PureInfo{}
	if err:=ctx.ReadJSON(pureInfo);err!=nil{
		panic(err.Error())
	}
	tagId:=models.PurifyUser(pureInfo.WeChatOpenID)
	DelRoleTag([]string{pureInfo.WeChatOpenID},tagId)
	fmt.Printf("去除用户"+pureInfo.WeChatOpenID+"的TagId:"+strconv.Itoa(tagId)+"\n")
}
