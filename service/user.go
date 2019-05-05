package service

import (
	"github.com/chanxuehong/wechat/mp/user"
	"github.com/chanxuehong/wechat/oauth2"
	"github.com/pantazheng/bci/models"
	"log"
)

func  checkOpenId (user *models.UserJson) (checkUser *models.UserJson ) {
	if user.OpenId=="" && user.Code!=""{
			token := &oauth2.Token{}
			if err:=ExchangeToken(token,user.Code);err!=nil{
				log.Printf("ExchangeTokenError: %v",err)
			}
			user.OpenId=token.OpenId
	}
	return user
}

//用户微信初始化
func UserInitByWechat(weChatInfo *user.UserInfo) string {
	newUser :=new(models.UserJson)
	newUser.OpenId=weChatInfo.OpenId
	newUser.Level=models.LevelStranger
	_, _ = models.UserCreate(newUser)
	log.Printf("UserInit:\t"+weChatInfo.OpenId)
	return "欢迎关注"
}

func UserCreate(user *models.UserJson)(userJson models.UserJson,err error){
	checkUser := new(models.UserJson)
	checkUser =checkOpenId(user)
	return models.UserCreate(checkUser)
}

func UserUpdate(user *models.UserJson)(userJson models.UserJson,err error){
	checkUser := new(models.UserJson)
	checkUser =checkOpenId(user)
	return models.UserUpdate(checkUser)
}

func UserFindByID(id uint)(recordUserJson models.UserJson,err error){
	recordUser :=new(models.User)
	recordUser.ID=id
	return models.UserFind(recordUser)
}

func UserFindByIDCard(idCard string)(recordUserJson models.UserJson,err error){
	recordUser :=new(models.User)
	recordUser.IDCard=idCard
	return models.UserFind(recordUser)
}

func UserFindByOpenID(openid string)(recordUserJson models.UserJson,err error){
	recordUser :=new(models.User)
	recordUser.OpenId=openid
	return models.UserFind(recordUser)
}

func UsersFindByLevel(level int)(usersBriefJson []models.UserBriefJson,err error){
	return models.UsersFindByLevel(level)
}

func UserDeleteByID(id uint)(recordUserJson models.UserJson,err error){
	recordUser:=new(models.User)
	recordUser.ID=id
	return models.UserDelete(recordUser)
}


func UserDeleteByOpenID(openid string)(recordUserJson models.UserJson,err error){
	recordUser:=new(models.User)
	recordUser.OpenId=openid
	return models.UserDelete(recordUser)
}

func UserBind(user *models.UserJson)(userJson models.UserJson,err error){
	checkUser := new(models.UserJson)
	checkUser =checkOpenId(user)
	//检查是否存在微信初始创建用户，有就删除
	if recordUser1,err:=UserFindByOpenID(checkUser.OpenId);err==nil{
		if recordUser1.IDCard!=""{
			u,_:=UserDeleteByOpenID(checkUser.OpenId)
			log.Println(u)
		}
	}else{
		log.Println(recordUser1)
		log.Println(err)
	}

	//获取更新对象的ID
	if recordUser2,err:=UserFindByIDCard(checkUser.IDCard);err==nil{
		checkUser.ID=recordUser2.ID
		return models.UserUpdate(checkUser)
	}
	return
}
