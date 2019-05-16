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
		//Emeritus Professor emeritus 专家教授，只能查看项目简要信息
		"Emeritus": 2,
		//Student 学生，查看项目的简要信息、参与模块详细信息、参与任务详细信息
		"Student": 3,
		//Senior Senior lecturer 高级讲师，只能查看自己项目下的所有信息
		"Senior": 4,
		//Assistant 助理,全部权限
		"Assistant": 5,
		//Full Full professor 全职教授，全部权限
		"Full": 6,
	}
)

//UserJSON 用户Json原型
type UserJSON struct {
	/**
	@Author: PantaZheng
	@Description:用户JSON
	@Date: 2019/5/9 10:42
	*/
	ID         uint   `json:"id"`
	OpenID     string `json:"openid"`
	WechatName string `json:"wechatName"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	IDCard     string `json:"idCard"`
	Level      int    `json:"level"`
	Telephone  string `json:"telephone"`
}

func userTestData() {
	log.Println("userTestData")
	users := make([]UserJSON, 8)
	users[0] = UserJSON{OpenID: "Stranger1", WechatName: "小蜘蛛", Code: "Spider-Man", Name: "Peter Benjamin Parker", Level: 1, Telephone: "110"}
	users[1] = UserJSON{OpenID: "Emeritus1", WechatName: "万磁王", Code: "002", Name: "Max Eisenhardt", IDCard: "Magneto", Level: 2}
	users[2] = UserJSON{WechatName: "金刚狼", IDCard: "Wolverine", Name: "Logan Howlett", Level: 3}
	users[3] = UserJSON{OpenID: "Assistant1", WechatName: "小辣椒", Name: "Pepper Potts", Level: 4}
	users[4] = UserJSON{WechatName: "钢铁侠", IDCard: "Iron Man", Name: "Tony Stark", Level: 5}
	users[5] = UserJSON{OpenID: "Full1", WechatName: "灭霸", IDCard: "5", Name: "Thanos", Level: 6}
	users[6] = UserJSON{IDCard: "6", Name: "海王", Level: 6}
	users[7] = UserJSON{IDCard: "7", Name: "雷神", Level: 6}
	for _, v := range users {
		if err := v.Create(); err != nil {
			log.Println(err.Error())
		} else {
			log.Println(v)
		}
	}
}

//userJSON2User UserJSON转换到User.
func (userJSON *UserJSON) userJSON2User() (user models.User) {
	/**
	@Author: PantaZheng
	@Description: UserJSON转化为User
	@Date: 2019/5/9 13:14
	*/
	user.ID = userJSON.ID
	user.OpenID = userJSON.OpenID
	user.WechatName = userJSON.WechatName
	user.Name = userJSON.Name
	user.IDCard = userJSON.IDCard
	user.Level = userJSON.Level
	user.Telephone = userJSON.Telephone
	return
}

//user2UserJSON User表单转换到UserJSON.
func user2UserJSON(user models.User) (userJSON UserJSON) {
	/**
	  @Author: PantaZheng
	  @Description:
	  @Date: 2019/5/9 12:04
	*/
	userJSON.ID = user.ID
	userJSON.OpenID = user.OpenID
	userJSON.WechatName = user.WechatName
	userJSON.Name = user.Name
	userJSON.IDCard = user.IDCard
	userJSON.Level = user.Level
	userJSON.Telephone = user.Telephone
	return
}

//user2UserBriefJSON 简化，仅保留：id、Name、Level.
func userJSON2UserBriefJSON(userJSON1 UserJSON) (userJSON2 UserJSON) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 2:57
	*/
	userJSON2 = UserJSON{}
	userJSON2.ID = userJSON1.ID
	userJSON2.Name = userJSON1.Name
	userJSON2.Level = userJSON1.Level
	return
}

func (userJSON *UserJSON) exchangeOpenID() (err error) {
	/**
	@Author: PantaZheng
	@Description: 根据code换取openid
	@Date: 2019/5/9 12:32
	*/
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

func usersJSON2Users(usersJSON []UserJSON) (users []models.User) {
	users = make([]models.User, len(usersJSON))
	for i, v := range usersJSON {
		users[i] = v.userJSON2User()
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
	/**
	@Author: PantaZheng
	@Description: 用户登录微信初始化微信,默认level等级为Stranger
	@Date: 2019/5/9 23:13
	*/
	message := ""
	u := UserJSON{OpenID: weChatInfo.OpenId}
	if err := u.FindOne(); err == nil {
		u.WechatName = weChatInfo.Nickname
		if err := u.Updates(); err == nil {
			message = "欢迎老用户" + u.WechatName + "重新关注"
		} else {
			message = titleUser + "UserInitByWechat:\t" + err.Error()
		}
	} else {
		u.WechatName = weChatInfo.Nickname
		u.Level = LevelMap["Stranger"]
		if err = u.Create(); err == nil {
			message = "欢迎新用户" + u.WechatName + ",请进行绑定操作"
		} else {
			message = titleUser + "UserInitByWechat:\t" + err.Error()
		}
	}
	return message
}

//Create 创建User.
func (userJSON *UserJSON) Create() (err error) {
	/**
	@Author: PantaZheng
	@Description: Service层 User创建
	@Date: 2019/5/9 23:11
	*/
	if err = userJSON.checkLevel(); err == nil {
		u := userJSON.userJSON2User()
		if err = u.Create(); err == nil {
			*userJSON = user2UserJSON(u)
		}
	}
	if err != nil {
		err = errors.New(titleUser + "Create:\t" + err.Error())
	}
	return
}

//Bind 用户绑定.
func (userJSON *UserJSON) Bind() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 2:31
	*/
	//检查openid和code
	if err = userJSON.exchangeOpenID(); err == nil {
		if userJSON.IDCard == "" || userJSON.Name == "" {
			err = errors.New("绑定必须有同时身份证和姓名信息\t")
		} else if userJSON.Level == LevelMap["Stranger"] {
			err = errors.New("绑定权限不能设置为1级\t")
		} else if err = userJSON.checkLevel(); err == nil {
			wechatUser := models.User{OpenID: userJSON.OpenID}
			//查找微信关联信息
			if err = wechatUser.FindOne(); err != nil {
				err = errors.New("数据库查找关联OpenID用户出错:\t" + err.Error())
			} else if wechatUser.Level > LevelMap["Stranger"] {
				err = errors.New("用户" + wechatUser.Name + "已经绑定过,如有修改需要请联系管理员")
			} else {
				presortedUser := models.User{IDCard: userJSON.IDCard}
				_ = presortedUser.FindOne()
				//检查是否有预存信息
				if presortedUser.Name != "" {
					presortedUser.OpenID = wechatUser.OpenID
					presortedUser.WechatName = wechatUser.WechatName
					//有预存信息，比对姓名
					if presortedUser.Name != userJSON.Name {
						err = errors.New("用户名:" + userJSON.Name + "和身份证号:" + userJSON.IDCard + "不匹配,请检查输入信息")
					} else if err = wechatUser.Delete(); err == nil {
						if presortedUser.Level == LevelMap["Stranger"] {
							presortedUser.Level = userJSON.Level
						}
						if err = presortedUser.Updates(); err == nil {
							*userJSON = user2UserJSON(presortedUser)
						}
					}
				} else {
					//无预存信息,修改微信初始化的User，添加姓名和身份证号
					wechatUser.Name = userJSON.Name
					wechatUser.IDCard = userJSON.IDCard
					wechatUser.Level = userJSON.Level
					if err = wechatUser.Updates(); err == nil {
						*userJSON = user2UserJSON(wechatUser)
					}
				}
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
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 1:52
	*/
	userJSON = UserJSON{ID: id}
	err = userJSON.First()
	return
}

//UserFindByOpenID 通过微信OpenID查找单用户.
func UserFindByOpenID(openid string) (userJSON UserJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 1:52
	*/
	userJSON = UserJSON{OpenID: openid}
	err = userJSON.FindOne()
	return
}

//UserFindByIDCard 通过身份证查找单用户.
func UserFindByIDCard(idCard string) (userJSON UserJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 1:52
	*/
	userJSON = UserJSON{IDCard: idCard}
	err = userJSON.FindOne()
	return
}

//Find 多用户查找的原子方法.
func (userJSON *UserJSON) Find() (usersJSON []UserJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 13:49
	*/
	u := userJSON.userJSON2User()
	if users, err := u.Find(); err == nil {
		usersJSON = users2BriefUsersJSON(users)
	} else {
		err = errors.New(titleUser + "Find:\t" + err.Error())
	}
	return
}

//UsersFindByLevel 通过权限等级查找用户表.
func UsersFindByLevel(level int) (usersJSON []UserJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 14:01
	*/
	userJSON := UserJSON{Level: level}
	if err = userJSON.checkLevel(); err == nil {
		usersJSON, err = userJSON.Find()
	}
	if err != nil {
		err = errors.New(titleUser + "UsersFindByLevel\t:" + err.Error())
	}
	return
}

//Updates 更新用户数据，id定位用户记录.
func (userJSON *UserJSON) Updates() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 14:09
	*/
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
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 14:16
	*/
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
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 14:28
	*/
	userJSON = UserJSON{ID: id}
	err = userJSON.Delete()
	return
}

//UserDeleteByIDCard 通过I身份证删除用户.
func UserDeleteByIDCard(idCard string) (userJSON UserJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/10 14:30
	*/
	userJSON = UserJSON{IDCard: idCard}
	err = userJSON.Delete()
	return
}
