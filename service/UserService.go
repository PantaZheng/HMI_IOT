package service

import (
	"../models"
	"github.com/chanxuehong/wechat/mp/user"
)

func CheckTableUser(){
	models.CheckTableUser()
}

//用户关注事件
func SubscribeInit(weChatinfo *user.UserInfo) string {
	if models.CheckUserByWeChatID(weChatinfo.OpenId) {
		if models.GetUserRoleByWechatID(weChatinfo.OpenId)!="unEnrolled"{
			return "欢迎已登记用户"
		}
	}else{
		models.CreateUser(weChatinfo)
	}
	return "欢迎关注，请先登记个人信息"
}

//教师信息注册
func UpdateTeacher(weChatOpenID string, teacherInfo *models.TeacherInfo) (msg string){
	models.EnrollTeacher(weChatOpenID,teacherInfo)
	return "欢迎教师"+teacherInfo.Name
}

//学生信息注册
func UpdateStudent(weChatOpenID string, studentInfo *models.StudentInfo) (msg string){
	models.EnrollStudent(weChatOpenID,studentInfo)
	return "欢迎同学"+studentInfo.Name
}