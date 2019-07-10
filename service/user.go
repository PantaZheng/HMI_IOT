package service

import (
	"errors"
	"github.com/chanxuehong/wechat/mp/user"
	"github.com/chanxuehong/wechat/oauth2"
	"github.com/pantazheng/bci/models"
	"log"
)

const titleUser = "service.user."

var (
	//LevelMap 用户权限管理
	LevelMap = map[string]int{
		//Stranger 未绑定
		"Stranger": 1,
		//Student 学生，查看项目的简要信息、参与模块详细信息、参与任务详细信息
		"Student": 2,
		//Senior Senior lecturer 高级讲师，只能查看自己项目下的所有信息
		"Senior": 3,
		//Assistant 助理,全部权限
		"Assistant": 4,
		//Full Full professor 全职教授，全部权限
		"Full": 5,
	}
)

//UserJSON 用户Json原型
type UserJSON struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	OpenID     string `json:"openid"`
	WechatName string `json:"wechatName"`
	Code       string `json:"code"`
	Level      int    `json:"level"`
	Telephone  string `json:"telephone"`
}

//userJSON2User UserJSON转换到User.
func (userJSON *UserJSON) userJSON2User() (user models.User) {
	user.ID = userJSON.ID
	user.OpenID = userJSON.OpenID
	user.WechatName = userJSON.WechatName
	user.Name = userJSON.Name
	user.Level = userJSON.Level
	user.Telephone = userJSON.Telephone
	return
}

//user2UserJSON User表单转换到UserJSON.
func user2UserJSON(user models.User) (userJSON UserJSON) {
	userJSON.ID = user.ID
	userJSON.OpenID = user.OpenID
	userJSON.WechatName = user.WechatName
	userJSON.Name = user.Name
	userJSON.Level = user.Level
	userJSON.Telephone = user.Telephone
	return
}

//user2UserBriefJSON 简化，仅保留：id、Name、Level.
func userJSON2UserBriefJSON(userJSON1 UserJSON) (userJSON2 UserJSON) {
	userJSON2 = UserJSON{}
	userJSON2.ID = userJSON1.ID
	userJSON2.Name = userJSON1.Name
	userJSON2.Level = userJSON1.Level
	return
}

func (userJSON *UserJSON) exchangeOpenID() (err error) {
	if userJSON.OpenID == "" {
		if userJSON.Code != "" {
			token := &oauth2.Token{}
			if err = ExchangeToken(token, userJSON.Code); err == nil {
				if token.OpenId == "" {
					err = errors.New("ExchangeToken,未换取到OpenID,请检查用户Code")
				} else {
					userJSON.OpenID = token.OpenId
				}
			}
		} else {
			err = errors.New("openid和code不可同时为空")
		}
	} else if userJSON.Code != "" {
		err = errors.New("有openid就别传code了，我不会去换的")
	}
	if err != nil {
		err = errors.New(titleUser + "exchangeOpenId:\t" + err.Error())
	}
	return
}

func users2BriefUsersJSON(users []models.User) (usersJSON []UserJSON) {
	usersJSON = make([]UserJSON, len(users))
	for i, v := range users {
		u := user2UserJSON(v)
		usersJSON[i] = userJSON2UserBriefJSON(u)
	}
	return
}

func (userJSON *UserJSON) checkLevel() (err error) {
	err = errors.New(titleUser + "checkLevel:\t" + "权限等级不在列表中")
	for _, v := range LevelMap {
		if v == userJSON.Level {
			err = nil
			return
		}
	}
	return
}

//UserInitByWechat 用户登录微信初始化微信.
func UserInitByWechat(weChatInfo *user.UserInfo) string {
	message := ""
	u := UserJSON{OpenID: weChatInfo.OpenId}
	if err := u.FindOne(); err == nil {
		u.WechatName = weChatInfo.Nickname
		if err := u.Updates(); err == nil {
			message = "欢迎用户" + u.WechatName + "关注"
		} else {
			message = titleUser + "UserInitByWechat:\t" + err.Error()
		}
	}
	return message
}

//Insert 创建User.
func (userJSON *UserJSON) Create() (err error) {
	if err = userJSON.checkLevel(); err == nil {
		u := userJSON.userJSON2User()
		if err = u.Insert(); err == nil {
			*userJSON = user2UserJSON(u)
		}
	}
	if err != nil {
		err = errors.New(titleUser + "Insert:\t" + err.Error())
	}
	return
}

//Bind 用户绑定.
func (userJSON *UserJSON) Bind() (err error) {
	//检查openid和code
	if err = userJSON.exchangeOpenID(); err == nil {
		if userJSON.Telephone == "" || userJSON.Name == "" {
			err = errors.New("绑定必须有同时电话和姓名信息\t")
		} else {
			wechatUser := models.User{OpenID: userJSON.OpenID}
			wechatInfo := &user.UserInfo{}
			if wechatInfo, err = GetUserInfo(userJSON.OpenID); err == nil {
				wechatUser.WechatName = wechatInfo.Nickname
				presortedUser := models.User{Telephone: userJSON.Telephone}
				_ = presortedUser.FindOne()
				//检查是否有预存信息
				if presortedUser.Name != "" {
					presortedUser.OpenID = wechatUser.OpenID
					presortedUser.WechatName = wechatUser.WechatName
					//有预存信息，比对姓名
					if presortedUser.Name != userJSON.Name {
						err = errors.New("用户名:" + userJSON.Name + "和电话:" + userJSON.Telephone + "不匹配,请检查输入信息")
					} else if err = presortedUser.Updates(); err == nil {
						*userJSON = user2UserJSON(presortedUser)
					}
				} else {
					//无预存信息,修改微信初始化的User，添加姓名和身份证号
					wechatUser.Name = userJSON.Name
					wechatUser.Telephone = userJSON.Telephone
					wechatUser.Level = LevelMap["Stranger"]
					if err = wechatUser.Insert(); err == nil {
						*userJSON = user2UserJSON(wechatUser)
					}
				}
			} else {
				err = errors.New("换取用户信息出错")
			}
		}
	}
	if err != nil {
		err = errors.New(titleUser + "Bind:\t" + err.Error())
	}
	return
}

//First 单用户查找的原子方法
func (userJSON *UserJSON) First() (err error) {
	u := userJSON.userJSON2User()
	if err = u.First(); err == nil {
		*userJSON = user2UserJSON(u)
	} else {
		err = errors.New(titleUser + "First:\t" + err.Error())
	}
	return
}

//FindOne 非ID方式查找用户
func (userJSON *UserJSON) FindOne() (err error) {
	u := userJSON.userJSON2User()
	if err = u.FindOne(); err == nil {
		*userJSON = user2UserJSON(u)
	} else {
		err = errors.New(titleUser + "FindOne:\t" + err.Error())
	}
	return
}

//UserFindByID 通过数据库ID查找单用户.
func UserFindByID(id uint) (userJSON UserJSON, err error) {
	userJSON = UserJSON{}
	userJSON.ID = id
	err = userJSON.First()
	return
}

//UserFindByOpenID 通过微信OpenID查找单用户.
func UserFindByOpenID(openid string) (userJSON UserJSON, err error) {
	userJSON = UserJSON{OpenID: openid}
	err = userJSON.FindOne()
	return
}

//UserFindByTelephone 通过电话查找单用户.
func UserFindByTelephone(telephone string) (userJSON UserJSON, err error) {
	userJSON = UserJSON{Telephone: telephone}
	err = userJSON.FindOne()
	return
}

//FindByField 多用户查找的原子方法.
func (userJSON *UserJSON) Find() (usersJSON []UserJSON, err error) {
	u := userJSON.userJSON2User()
	if users, err := u.Find(); err == nil {
		usersJSON = users2BriefUsersJSON(users)
	} else {
		err = errors.New(titleUser + "FindByField:\t" + err.Error())
	}
	return
}

//UsersFindByLevel 通过权限等级查找用户表.
func UsersFindByLevel(level int) (usersJSON []UserJSON, err error) {
	userJSON := UserJSON{Level: level}
	if err = userJSON.checkLevel(); err == nil {
		usersJSON, err = userJSON.Find()
	}
	if err != nil {
		err = errors.New(titleUser + "UsersFindByLevel\t:" + err.Error())
	}
	return
}

//UsersList 返回全部用户列表
func UsersList() (usersJSON []UserJSON, err error) {
	userJSON := UserJSON{}
	usersJSON, err = userJSON.Find()
	if err != nil {
		err = errors.New(titleUser + "UsersList\t:" + err.Error())
	}
	return
}

//Updates 更新用户数据，id定位用户记录.
func (userJSON *UserJSON) Updates() (err error) {
	if err = userJSON.checkLevel(); err == nil {
		u := userJSON.userJSON2User()
		if err = u.Updates(); err == nil {
			*userJSON = user2UserJSON(u)
		}
	}
	if err != nil {
		err = errors.New(titleUser + "Updates:\t" + err.Error())
	}
	return
}

//Delete 用户删除的原子方法.
func (userJSON *UserJSON) Delete() (err error) {
	u := userJSON.userJSON2User()
	if err = u.Delete(); err == nil {
		*userJSON = user2UserJSON(u)
	}
	if err != nil {
		err = errors.New(titleUser + "Delete:\t" + err.Error())
	}
	return
}

//UserDeleteByID 通过数据库ID删除用户.
func UserDeleteByID(id uint) (userJSON UserJSON, err error) {
	userJSON = UserJSON{}
	userJSON.ID = id
	err = userJSON.Delete()
	return
}

//UserDeleteByTelephone 通过电话删除用户.
func UserDeleteByTelephone(telephone string) (userJSON UserJSON, err error) {
	userJSON = UserJSON{Telephone: telephone}
	err = userJSON.Delete()
	return
}

func UserDeleteByOpenID(openid string) {
	userJSON := UserJSON{OpenID: openid}
	if err := userJSON.Delete(); err != nil {
		log.Println("UserDeleteByOpenID:\t" + err.Error())
	}
}
