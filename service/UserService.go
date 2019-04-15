package service

import (
	"../models"
	"github.com/chanxuehong/wechat/mp/user"
	"github.com/chanxuehong/wechat/oauth2"
	_ "github.com/chanxuehong/wechat/oauth2"
	"log"
)

func checkOpenId(openid string,code string) (checkOpenId string){
	if openid==""{
		token := &oauth2.Token{}
		if err:=ExchangeToken(token,code);err!=nil{
			log.Printf("ExchangeTokenError: %v",err)
		}
		return token.OpenId
	}else{
		return openid
	}
}

func CheckTableUser(){
	models.DropTableUsers()
	models.MakeTestData()
	//models.CheckTableUser()
}

func GetMembers(role string) (memberList []models.MemberInfo){
	memberList=models.GetMembersByRole(role)
	return
}


//用户初始化
func UserInit(weChatInfo *user.UserInfo) string {
	userInit:=&models.User{}
	userInit.OpenId=weChatInfo.OpenId
	userInit.Role="unEnrolled"
	if err:=models.EnrollUser(userInit);err!=nil{
		panic(err.Error())
	}
	log.Printf("UserInit:\t"+weChatInfo.OpenId)
	return "欢迎关注"
}


func Enroll(userEnroll  *models.User)(openid string){
	userEnroll.OpenId=checkOpenId(userEnroll.OpenId, userEnroll.Code)
	if err:=models.EnrollUser(userEnroll);err!=nil{
		panic(err.Error())
	}
	log.Printf(userEnroll.OpenId+"\tEnrollRole\t"+userEnroll.Role+"\n")
	return userEnroll.OpenId
}
