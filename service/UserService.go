package service

import (
	"github.com/chanxuehong/wechat/mp/user"
	"github.com/chanxuehong/wechat/oauth2"
	"github.com/pantazheng/HMI-IoT/models"
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


//用户初始化
func UserInit(weChatInfo *user.UserInfo) string {
	userInit:=&models.User{}
	userInit.OpenId=weChatInfo.OpenId
	userInit.Role="unEnrolled"
	models.EnrollUser(userInit)
	log.Printf("UserInit:\t"+weChatInfo.OpenId)
	return "欢迎关注"
}

func GetMembers(role string) (memberList []models.MemberInfo){
	memberList=models.GetMembersByRole(role)
	return
}

func Enroll(userEnroll  *models.User)(openid string){
	userEnroll.OpenId=checkOpenId(userEnroll.OpenId, userEnroll.Code)
	models.EnrollUser(userEnroll)
	log.Printf(userEnroll.OpenId+"\tEnrollRole\t"+userEnroll.Role+"\n")
	return userEnroll.OpenId
}
