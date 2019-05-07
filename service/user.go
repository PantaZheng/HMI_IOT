package service

import (
	"github.com/chanxuehong/wechat/mp/user"
	"github.com/chanxuehong/wechat/oauth2"
	"github.com/pantazheng/bci/models"
	"log"
)

func checkOpenId(user *models.UserJSON) (checkUser *models.UserJSON) {
	if user.OpenID == "" && user.Code != "" {
		token := &oauth2.Token{}
		if err := ExchangeToken(token, user.Code); err != nil {
			log.Printf("ExchangeTokenError: %v", err)
		}
		user.OpenID = token.OpenId
	}
	return user
}

//用户微信初始化
func UserInitByWechat(weChatInfo *user.UserInfo) string {
	newUser := new(models.UserJSON)
	newUser.OpenID = weChatInfo.OpenId
	newUser.Level = models.LevelStranger
	_, _ = models.UserCreate(newUser)
	log.Printf("UserInit:\t" + weChatInfo.OpenId)
	return "欢迎关注"
}

func UserCreate(user *models.UserJSON) (userJson models.UserJSON, err error) {
	checkUser := new(models.UserJSON)
	checkUser = checkOpenId(user)
	return models.UserCreate(checkUser)
}

func UserUpdate(user *models.UserJSON) (userJson models.UserJSON, err error) {
	checkUser := new(models.UserJSON)
	checkUser = checkOpenId(user)
	return models.UserUpdate(checkUser)
}

func UserFindByID(id uint) (recordUserJson models.UserJSON, err error) {
	recordUser := new(models.User)
	recordUser.ID = id
	return models.UserFind(recordUser)
}

func UserFindByIDCard(idCard string) (recordUserJson models.UserJSON, err error) {
	recordUser := new(models.User)
	recordUser.IDCard = idCard
	return models.UserFind(recordUser)
}

func UserFindByOpenID(openid string) (recordUserJson models.UserJSON, err error) {
	recordUser := new(models.User)
	recordUser.OpenID = openid
	return models.UserFind(recordUser)
}

func UsersFindByLevel(level int) (usersBriefJson []models.UserBriefJSON, err error) {
	return models.UsersFindByLevel(level)
}

func UserDeleteByID(id uint) (recordUserJson models.UserJSON, err error) {
	recordUser := new(models.User)
	recordUser.ID = id
	return models.UserDelete(recordUser)
}

func UserDeleteByOpenID(openid string) (recordUserJson models.UserJSON, err error) {
	recordUser := new(models.User)
	recordUser.OpenID = openid
	return models.UserDelete(recordUser)
}

func UserBind(user *models.UserJSON) (userJson models.UserJSON, err error) {
	checkUser := new(models.UserJSON)
	checkUser = checkOpenId(user)
	//检查是否存在微信初始创建用户，有就删除
	if recordUser1, err := UserFindByOpenID(checkUser.OpenID); err == nil {
		if recordUser1.IDCard == "" {
			_, _ = UserDeleteByOpenID(recordUser1.OpenID)
		}
	}
	//获取更新对象的ID
	if recordUser2, err := UserFindByIDCard(checkUser.IDCard); err == nil {
		checkUser.ID = recordUser2.ID
		return models.UserUpdate(checkUser)
	}
	return
}
