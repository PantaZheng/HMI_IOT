package service

import (
	"errors"
	"github.com/chanxuehong/wechat/mp/user"
	"github.com/chanxuehong/wechat/oauth2"
	"github.com/pantazheng/bci/models"
	"log"
)

//UserJSON 用户Json原型
type UserJSON struct {
	/**
	@Author: PantaZheng
	@Description:用户JSON
	@Date: 2019/5/9 10:42
	*/
	ID        int    `json:"id,omitempty"`
	OpenID    string `json:"openid,omitempty"`
	Code      string `json:"code,omitempty"`
	Name      string `json:"name"`
	IDCard    string `json:"idCard,omitempty"`
	Level     int    `json:"level"`
	Telephone string `json:"telephone,omitempty"`
}

//User2UserJSON User表单转换到UserJSON.
func User2UserJSON(user *models.User, userJSON *UserJSON) {
	/**
	  @Author: PantaZheng
	  @Description:
	  @Date: 2019/5/9 12:04
	*/
	userJSON.ID = int(user.ID)
	userJSON.OpenID = user.OpenID
	userJSON.Code = user.Code
	userJSON.Name = user.Name
	userJSON.IDCard = user.IDCard
	userJSON.Level = user.Level
	userJSON.Telephone = user.Telephone
}

//UserJSON2User
func UserJSON2User(userJSON *UserJSON, user *models.User) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/9 13:14
	*/
	user.ID = uint(userJSON.ID)
	user.OpenID = userJSON.OpenID
	user.Name = userJSON.Name
	user.Code = userJSON.Code
	user.IDCard = userJSON.IDCard
	user.Level = userJSON.Level
	user.Telephone = userJSON.Telephone
}

func (userJSON UserJSON) checkUniqueConstraint(err error) {
	/**
	@Author: PantaZheng
	@Description:检查UserJSON的唯一性要求是否满足，ID,
	OpenID,IDCard
	@Date: 2019/5/9 10:44
	*/
	if userJSON.OpenID == "" && userJSON.ID == 0 {
		err = errors.New("checkUniqueConstraint:\t\n需要OpenID或ID来满足用户唯一性")
	}
}

func exchangeOpenId(code string) (openid string, err error) {
	/**
	@Author: PantaZheng
	@Description: 根据code换取openid
	@Date: 2019/5/9 12:32
	*/
	token := &oauth2.Token{}
	if err := ExchangeToken(token, code); err != nil {
		err = errors.New("exchangeOpenId:\t\n" + err.Error())
	}
	openid = token.OpenId
	return
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
	//TODO:判断id_card是否存在，存在就捆绑；没有就创建用户
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
