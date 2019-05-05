package service

import (
	"github.com/chanxuehong/wechat/mp/user"
	"github.com/chanxuehong/wechat/oauth2"
	"github.com/pantazheng/bci/models"
	"log"
	"strconv"
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
	newUser :=new(models.UserJson)
	newUser.OpenId=weChatInfo.OpenId
	newUser.Level=models.LevelStranger
	_, _ = models.UserCreate(newUser)
	log.Printf("UserInit:\t"+weChatInfo.OpenId)
	return "欢迎关注"
}

func Enroll(userEnroll  *models.User)(openid string){
	userEnroll.OpenId=checkOpenId(userEnroll.OpenId, userEnroll.Code)
	_, _ = models.UserCreate(userEnroll)
	log.Printf(userEnroll.OpenId+"\tEnrollRole\t"+strconv.Itoa(userEnroll.Level)+"\n")
	return userEnroll.OpenId
}
