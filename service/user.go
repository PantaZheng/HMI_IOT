package service

import (
	"errors"
	"github.com/chanxuehong/wechat/mp/user"
	"github.com/chanxuehong/wechat/oauth2"
	"github.com/pantazheng/bci/models"
	"log"
)

const title = "service.user."

var (
	//LevelMap 用户权限管理
	LevelMap = map[string]int{
		//Stranger 未绑定
		"Stranger": 1,
		//Emeritus Professor emeritus 专家教授
		"Emeritus": 2,
		//Student 学生
		"Student": 3,
		//Assistant 助理
		"Assistant": 4,
		//Senior Senior lecturer 高级讲师
		"Senior": 5,
		//Full Full professor 全职教授
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
	u0 := &UserJSON{OpenID: "Stranger1", WechatName: "小蜘蛛", Code: "Spider-Man", Name: "Peter Benjamin Parker", Level: 1, Telephone: "110"}
	u1 := &UserJSON{OpenID: "Emeritus1", WechatName: "万磁王", Code: "002", Name: "Max Eisenhardt", IDCard: "Magneto", Level: 2}
	u2 := &UserJSON{WechatName: "金刚狼", IDCard: "Wolverine", Name: "Logan Howlett", Level: 3}
	u3 := &UserJSON{OpenID: "Assistant1", WechatName: "小辣椒", Name: "Pepper Potts", Level: 4}
	u4 := &UserJSON{WechatName: "钢铁侠", IDCard: "Iron Man", Name: "Tony Stark", Level: 5}
	u5 := &UserJSON{OpenID: "Full1", WechatName: "灭霸", IDCard: "5", Name: "Thanos", Level: 6}
	_ = u0.Create()
	_ = u1.Create()
	_ = u2.Create()
	_ = u3.Create()
	_ = u4.Create()
	_ = u5.Create()
}

//UserJSON2User UserJSON转换到User.
func (userJSON *UserJSON) UserJSON2User() (user models.User) {
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

//User2UserJSON User表单转换到UserJSON.
func User2UserJSON(user *models.User) (userJSON UserJSON) {
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
		err = errors.New(title + "exchangeOpenId:\t" + err.Error())
	}
	return
}

//simplify 简化，仅保留：id、Name、Level.
func (userJSON *UserJSON) simplify() {
	userJSON.OpenID = ""
	userJSON.WechatName = ""
	userJSON.Code = ""
	userJSON.IDCard = ""
	userJSON.Level = 0
	userJSON.Telephone = ""
}

func (userJSON *UserJSON) checkLevel() (err error) {
	err = errors.New(title + "Find:\t" + "权限等级不在列表中")
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
	u := new(UserJSON)
	u.OpenID = weChatInfo.OpenId
	u.WechatName = weChatInfo.Nickname
	u.Level = LevelMap["Stranger"]
	if err := u.Create(); err != nil {
		return title + "UserInitByWechat:\t" + err.Error()
	}
	log.Printf("UserInit:\t" + weChatInfo.OpenId)
	return "欢迎关注"
}

//Create 创建User.
func (userJSON *UserJSON) Create() (err error) {
	/**
	@Author: PantaZheng
	@Description: Service层 User创建
	@Date: 2019/5/9 23:11
	*/
	if err = userJSON.checkLevel(); err == nil {
		u := userJSON.UserJSON2User()
		if err = u.Create(); err == nil {
			//userJSON接收数据库创建记录
			*userJSON = User2UserJSON(&u)
		}
	}
	if err != nil {
		err = errors.New(title + "Create:\t" + err.Error())
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
		} else {
			wechatUser := &models.User{OpenID: userJSON.OpenID}
			//查找微信关联信息
			if err = wechatUser.FindOne(); err != nil {
				err = errors.New("数据库查找关联OpenID用户出错:\t" + err.Error())
			} else if wechatUser.Level > LevelMap["Stranger"] {
				err = errors.New("用户" + wechatUser.Name + "已经绑定过,如有修改需要请联系管理员")
			} else {
				presortedUser := &models.User{IDCard: userJSON.IDCard}
				_ = presortedUser.FindOne()
				//检查是否有预存信息
				if presortedUser.Name != "" {
					//有预存信息，比对姓名
					if presortedUser.Name != userJSON.Name {
						err = errors.New("用户名:" + userJSON.Name + "和身份证号:" + userJSON.IDCard + "不匹配,请检查输入信息")
					} else if err = wechatUser.Delete(); err == nil {
						//预存User添加微信信息
						presortedUser.OpenID = userJSON.OpenID
						presortedUser.WechatName = userJSON.WechatName
						if err = presortedUser.Updates(); err == nil {
							*userJSON = User2UserJSON(presortedUser)
						}
					}
				} else {
					//无预存信息,修改微信初始化的User，添加姓名和身份证号
					wechatUser.Name = userJSON.Name
					wechatUser.IDCard = userJSON.IDCard
					if err = wechatUser.Updates(); err == nil {
						*userJSON = User2UserJSON(wechatUser)
					}
				}
			}
		}
	}
	if err != nil {
		err = errors.New(title + "Bind:\t" + err.Error())
	}
	return
}

//First 单用户查找的原子方法
func (userJSON *UserJSON) First() (err error) {
	u := userJSON.UserJSON2User()
	if err = u.First(); err == nil {
		*userJSON = User2UserJSON(&u)
	} else {
		err = errors.New(title + "First:\t" + err.Error())
	}
	return
}

func (userJSON *UserJSON) FindOne() (err error) {
	u := userJSON.UserJSON2User()
	if err = u.FindOne(); err == nil {
		*userJSON = User2UserJSON(&u)
	} else {
		err = errors.New(title + "FindOne:\t" + err.Error())
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
	u := userJSON.UserJSON2User()
	if users, err := u.Find(); err == nil {
		usersJSON = make([]UserJSON, len(users))
		for i, v := range users {
			usersJSON[i] = User2UserJSON(v)
		}
	} else {
		err = errors.New(title + "Find:\t" + err.Error())
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
		if len(usersJSON) == 0 {
			err = errors.New("当前权限的没有用户存在")
		}
	}
	if err != nil {
		err = errors.New(title + "UsersFindByLevel\t:" + err.Error())
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
		if userJSON.ID == 0 {
			err = errors.New("更新信息必须包含用户ID")
		} else {
			u := userJSON.UserJSON2User()
			if err = u.Updates(); err == nil {
				*userJSON = User2UserJSON(&u)
			}
		}
	}
	if err != nil {
		err = errors.New(title + "Updates:\t" + err.Error())
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
	u := userJSON.UserJSON2User()
	if err = u.Delete(); err == nil {
		*userJSON = User2UserJSON(&u)
	}
	if err != nil {
		err = errors.New(title + "Delete:\t" + err.Error())
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
