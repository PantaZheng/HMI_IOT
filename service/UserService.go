package service

import (
	"../models"
	"github.com/chanxuehong/wechat/mp/user"
)

//用户关注事件
func SubscribeInit(weChatinfo *user.UserInfo) string {
	models.CheckTableUser()
	if models.CheckUserByWeChatID(weChatinfo.OpenId)==false {
		return "欢迎老用户重新关注"
	}else{
		models.CreateUser(weChatinfo)
		return "欢迎新用户关注，请先注册个人信息，由于微信刷新机制，请先取消关注再重新关注以获得专属菜单"
	}
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