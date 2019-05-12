# BCI脑机协同微信公众号

## 域名

`http://bci.renjiwulian.com`

---

### 微信菜单

[菜单预览](https://pantazheng.github.io/HMI_IOT/design/index.html)

- 人员Person
  1. 绑定 view    `/user/index`
  1. 架构 view    `/frame/index`
- 内容
  1. 新建 view    `/new/index`
  1. 项目 view    `/project/index`
  1. 任务 view    `/mission/index`
- 进度      view    `/pace/index`

---

## PATH

### WeChat

- 入口: `/anon/wechat`

### User

入口: `/user`

```go
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
//正常的User用户的CODE和IDCard不得为ID的数字
```

名称|method|path|传入body参数|接收body参数|
-|-|-|:-|:-
UserCreate|post|`/`|`UserJson`<br>IDCard必须存在|`UserJson`<br>id,openid,id_card三者至少存在一个，其他项均可缺省
UserFindByID|get|`/id/{id:uint}`|-|`UserJson`|
UserFindByIDCard|`/id_card/{id_card:string}`|-|`UserJson`
UserFindByOpenID|`/openid/{openid:string}`|-|`UserJson`
UsersFindByLevel|get|`/level/{level:int}`|-|`[]UserJson`|仅包含id,name
UserUpdate|put|`/update`|`UserJson`|`UserJson`
UserBind|put|`/bind`|`UserJson`<br>openid,code仅且存在一个|`UserJson`
UserDeleteByID|delete|`/id/{id:uint}`|-|`UserJson`
UserDeleteByOpenID|delete|`/id/{openid:string}`|-|`UserJson`

### Gain

入口: `/gain`

```go
type GainJson struct {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 1:17
	*/
	ID        uint     `json:"id"`
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	File      string   `json:"file"`
	UpTime    string   `json:"upTime"`
	Remark    string   `json:"remark"`
	Owner     UserJSON `json:"owner"`
	MissionID uint     `json:"missionID"`
}
```

名称|method|path|传入body参数|接收body参数
-|-|-|-|-
GainCreate|post|`/`|`GainJson`|`GainJson`
GainFindByID|get|`/id/{id:uint}`|-|`GainJson`
GainsFindByOwnerID|get|`/owner/{id:uint}`|-|`[]GainJson`
GainsFindByMissionID|get|`/mission/{id:uint}`|-|`[]GainJson`
GainUpdate|put|`/`|`GainJson`|`GainJson`
GainDeleteByID|delete|`/{id:uint}`|-|`GainJson`

### Mission

入口: `/mission`

```go
type MissionJson struct {
	ID           uint            `json:"id"`
	Name         string          `json:"name"`
	Creator      UserBriefJSON   `json:"creator"`
	CreateTime   string          `json:"createTime"`
	StartTime    string          `json:"startTime"`
	EndTime      string          `json:"endTime"`
	Content      string          `json:"content"`
	File         string          `json:"file"`
	Tag          bool            `json:"tag"` //tag由module负责人决定
	Gains        []GainJson      `json:"gains"`
	Participants []UserBriefJSON `json:"participants"`
	ModuleID     uint            `json:"module"`
}

type MissionBriefJson struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	CreateTime string `json:"createTime"`
	Content    string `json:"content"`
	Tag        bool   `json:"tag"`
	ModuleID   uint   `json:"module"`
}

```

名称|method|path|传入body参数|接收body参数
-|-|-|-|-
MissionCreate|post|`/`|`MissionJson`|`MissionJson`
MissionFindByID|get|`/id/{id:uint}`|-|`MissionJson`
MissionsFindByModuleID|get|`/module/{id:uint}`|-|`[]MissionBriefJson`
MissionUpdate|put|`/`|`MissionJson`|`MissionJson`
MissionDeleteByID|delete|`/id/{id:uint}`|-|`MissionJson`

### Module

入口: `/module`

```go
type ModuleJson struct {
	ID           uint               `json:"id"`
	Name         string             `json:"name"`
	Creator      UserBriefJSON      `json:"creator"`
	CreateTime   string             `json:"create_time"` //创建时间
	StartTime    string             `json:"start_time"`  //开始时间
	EndTime      string             `json:"end_time"`    //结束时间
	Content      string             `json:"content"`
	Tag          bool               `json:"tag"`
	ProjectID    uint               `json:"project_id"`
	Leader       UserBriefJSON      `json:"leader"`
	Participants []UserBriefJSON    `json:"participants"` //参与人员
	Missions     []MissionBriefJson `json:"missions"`     //创建或更新不会修改该字段，仅拉取使用
}

type ModuleBriefJson struct {
	ID         uint          `json:"id"`
	Name       string        `json:"name"`
	CreateTime string        `json:"create_time"` //创建时间
	Content    string        `json:"content"`
	Tag        bool          `json:"tag"`
	Leader     UserBriefJSON `json:"leader"`
	ProjectID  uint          `json:"project"`
}
```

名称|method|path|传入body参数|接收body参数
-|-|-|-|-
ModuleCreate|post|`/`|`ModuleJson`|`ModuleJson`
ModuleFindByID|get|`/id/{id:uint}`|-|`ModuleJson`
ModulesFindByLeaderID|get|`/leader/{id:uint}`|-|`[]ModuleBriefJson`
ModulesFindByProjectID|get|`/project/{id:uint}`|-|`[]ModuleBriefJson`
ModuleUpdate|put|`/`|`ModuleJson`|`ModuleJson`
ModuleDeleteByID|delete|`/id/{id:uint}`|-|`ModuleJson`

### Project

```go
type ProjectJson struct {
	ID           uint              `json:"id"`
	Name         string            `json:"name"`
	Type         string            `json:"type"`
	Creator      UserBriefJSON     `json:"creator"`
	CreateTime   string            `json:"create_time"`
	StartTime    string            `json:"start_time"`
	EndTime      string            `json:"end_time"`
	Content      string            `json:"content"`
	Targets      []string          `json:"targets"`
	Leader       UserBriefJSON     `json:"leader"`
	Participants []UserBriefJSON   `json:"participants"`
	Tag          bool              `json:"tag"` //create、update
	TagSet       []TagJson         `json:"tags"`
	Modules      []ModuleBriefJson `json:"modules"` //仅拉取更新
}

type ProjectBriefJson struct {
	ID        uint          `json:"id"`
	Name      string        `json:"name"`
	StartTime string        `json:"startTime"`
	EndTime   string        `json:"endTime"`
	Leader    UserBriefJSON `json:"leader"`
	Tag       bool          `json:"tag"`
	Content   string        `json:"content"`
}

type TagJson struct {
	ID  uint `json:"id"`
	Tag bool `json:"tag"`
}
```

名称|method|path|传入body参数|接收body参数
-|-|-|-|-
ProjectCreate|post|`/`|`ProjectJson`|`ProjectJson`
ProjectFindByID|get|`/id/{id:uint}`|-|`ProjectJson`
ProjectsFindByLeaderID|get|`/leader/{id:uint}`|-|`[]ProjectBriefJson`
ProjectssFindByParticipantID|get|`/participant/{id:uint}`|-|`[]ProjectBriefJson`
ProjectUpdate|put|`/`|`projectJson`|`ProjectJson`
ProjectDeleteByID|delete|`/id/{id:uint}`|-|`ProjectJson`

---

## TODO

### top deign

2019-5-9

- [ ] JSON 放到service层处理 ---
- [ ] 微信在service层面的对后台消息的反馈
- [ ] 判断数据库表是否为空

2019-5-8

- [ ] PACE进度实现
- [ ] FRAME框架实现
- [ ] 之前的实例化方式都是指针创建，改用对象---
- [x] level改用`map`实现
- [ ] createTime不用加time,只用data
- [ ] 时间风格，在utils中添加转换时间的工具
  - `layout: 2006-01-02`
- [ ] models、service采用`method`重写`function`，----
- [ ] 使用`interface`写不同等级用户，允许拥有不同的权限
- [x] mysql 使用`utf8mb4`，用户表加入微信用户名
  - `mb4`: most bytes 4
- [ ] 操作系统将文件重定向到文件

  ```bash
  > .log 2>&1
  ```

- [ ] JSON数据做忽略零值处理`omitempty`---
- [x] JSON命名风格改为骆驼风格

### models

- [ ] gain
  - [x] `type Gain struct`
  - [x] `Create`
    1. uptime
    1. `db.create`
  - [x] `First`
    1. checkid
    1. `db.first`
  - [x] `FindByOwner`
    1. `owner.findone`
    1. `db.Model(&owner).Related(&gains, "OwnerID")`
  - [x] `GainsFindByOwner`
  - [ ] `GainsFindByMission`
  - [ ] `GainUpdate`
    1. checkid
    1. update uptime
    1. `database.DB.Model(&g).Updates(&gain)`
  - [ ] `GainDelete`
    1. checkid
    1. `db.delete`
- [ ] mission
  - [ ] `type Mission struct`
  - [ ] `type MissionJson struct`
  - [ ] `type MissionBriefJson struct`
  - [ ] `missionTestData`
  - [ ] `missionJson2Mission`
  - [ ] `mission2MissionJSON`
  - [ ] `mission2MissionBriefJSON`
  - [ ] `MissionCreate`
  - [ ] `MissionFind`
  - [ ] `MissionsFindByModule`
  - [ ] `MissionUpdate`
  - [ ] `MissionDelete`
- [ ] module
  - [ ] `type Module struct`
  - [ ] `type ModuleJson struct`
  - [ ] `type ModuleBriefJson struct`
  - [ ] `moduleTestData`
  - [ ] `moduleJson2Module`
  - [ ] `module2ModuleJson`
  - [ ] `module2ModuleBriefJson`
  - [ ] `ModuleCreate`
  - [ ] `ModuleFind`
  - [ ] `ModulesFindByLeader`
  - [ ] `ModulesFindByProject`
  - [ ] `ModuleUpdate`
  - [ ] `ModuleDelete`
- [ ] project
  - [ ] `type Project struct`
  - [ ] `type ProjectJson struct`
  - [ ] `type BriefProject struct`
  - [ ] `type TagJson struct`
  - [ ] `projectTestData`
  - [ ] `target2TargetsJson`
  - [ ] `targetsJson2Target`
  - [ ] `tagSet2TagsJson`
  - [ ] `tagsJson2TagSet`
  - [ ] `projectJson2Project`
  - [ ] `project2ProjectJson`
  - [ ] `project2ProjectBriefJson`
  - [ ] `ProjectCreate`
  - [ ] `ProjectFind`
  - [ ] `ProjectsFindByLeader`
  - [ ] `ProjectsFindByParticipant`
  - [ ] `ProjectUpdate`
  - [ ] `ProjectDelete`
  - [ ] `../project_test`
- [x] user
  - [ ] `type User struct` User 数据库用户表.
    - 与其他表的关联没有达成
  - [x] `checkUnique` 检查UserJSON的唯一性要求是否满足，ID,OpenID,IDCard
  - [x] `makeOpenIDIDCARDNotEmpty` 检查是否有OpenID和IDCard，零值设置为ID,并更新字段信息
  - [x] `Create` User Create
    1. `checkUnique`
    1. `db.Create`
    1. `makeOpenIDIDCARDNotEmpty`
        1. `user.Updates`
  - [x] `First` 根据id查找用户.
    1. `checkUnique`
    1. `db.First`
  - [x] `FindOne` 单个查找非主键.
    1. `db.find`
    1. `check len(users)`
  - [x] `Find` 查找多个用户
    1. `db.find`
  - [x] `Updates` Updates 非覆盖式更新,零值不更新,根据ID定位用户.
    1. `db.Updates`
  - [x] `Delete` Delete 先将openid和idCard置为id，再软删除.
    1. ID->OpenID,IDCard
    1. `user.Updates`
    1. `db.Delete`
- [x] init
  - [x] 表单删除
  - [x] 表单迁移
  - [x] 添加测试数据
- [ ] frame

### servcie

- [x] gain
  - [x] `type GainJson struct`
  - [x] `gainTestData`
  - [x] `gain2GainJSON`
  - [x] `gainJSON2GainBriefJSON`
  - [x] `gainJSON2Gain`
  - [x] `Create`
    1. TODO: 相关模块人员查验，当前是`owner.First()`仅校验owner存在
    1. `g.Create()`
  - [x] `First`
    - `g.First()`
  - [x] `GainsFindByOwnerID`
    1. 服务层对任务做精简 Gains2SimpleGains
  - [x] `GainsFindByMissionID`
  - [x] `GainUpdate`
    - 必须携带ID
  - [x] `GainDeleteByID`
    - 必须携带ID
- [x] mission
  - [x] `MissionCreate`
  - [x] `MissionFindByID`
    - 通过mission id去查成果
  - [x] `MissionFindByName`
  - [x] `MissionsFindByModuleID`
  - [x] `MissionUpdate`
    - 必须携带ID
  - [x] `MissionDeleteByID`
  - [x] `MissionDeleteByName`
- [x] module
  - [x] `ModuleCreate`
  - [x] `ModuleFindByID`
  - [x] `ModulesFindByLeaderID`
  - [x] `ModulesFindByProjectID`
  - [x] `ModuleUpdate`
    - id
  - [x] `ModuleDeleteByID`
- [x] project
  - [x] `ProjectCreate`
  - [x] `ProjectFindByID`
  - [x] `ProjectsFindByLeaderID`
  - [x] `ProjectsFindByParticipantID`
  - [x] `ProjectUpdate`
  - [x] `ProjectDeleteByID`
- [x] user
  - [x] `userTestData`
  - [x] `UserJSON2User`
  - [x] `User2UserJSON`
  - [x] `exchangeOpenId`
    1. userJSON.OpenID == ""
    1. userJSON.Code != ""
    1. ExchangeToken
    1. userJSON.OpenID = token.OpenId
  - [x] `simplify` simplify 简化，仅保留：id、Name、Level.
  - [x] `checkLevel`
  - [x] `UserInitByWechat`
    1. `u.FindOne()`
        1. `u.Updates()` old
        1. `u.Create()` new
  - [x] `Create`
    1. `userJSON.checkLevel()`
    1. `u.Create()`
  - [x] `Bind`
    1. `exchangeOpenID`
    1. check IDCard和Name是否有空
    1. `level !=LevelMap["Stranger"]`
    1. `checklevel illegal`
    1. `wechatUser.FindOne()`
    1. `wechatUser.Level > LevelMap["Stranger"]`
    1. `presortedUser.FindOne()`
        1. `presortedUser.Name != ""`
            1. Openid,WechatName->presortedUser
            1. `presortedUser.Name != userJSON.Name`
            1. `if presortedUser.Level == LevelMap["Stranger"]`
            1. `wechatUser.Delete()`
            1. `presortedUser.Updates()`
        1. `wechatUser.Updates()`
  - [x] `First`
    1. `u.First()`
  - [x] `FindOne`
    1. `u.FindOne()`
  - [x] `UserFindByID`
  - [x] `UserFindByOpenID`
  - [x] `UserFindByIDCard`
  - [x] `Find`
    1. `u.Find()`
    1. `simplify()`
  - [x] `UsersFindByLevel`
    1. `checkLevel()`
    1. `userJSON.Find()`
  - [x] `Updates`
    1. `checkLevel`
    1. `u.Updates()`
  - [x] `Delete`
    1. `u.Delete()`
  - [x] `UserDeleteByID`
  - [x] `UserDeleteByOpenID`

### controller

- [x] gain
  - [x] `GainCreate`
  - [x] `GainFindByID`
  - [x] `GainsFindByOwnerID`
  - [x] `GainsFindByMissionID`
  - [x] `GainUpdate`
    - 必须携带ID
  - [x] `GainDeleteByID`
    - 必须携带ID
- [x] mission
  - [x] `MissionCreate`
  - [x] `MissionFindByID`
    - 通过mission id去查成果
  - [x] `MissionFindByName`
    - Name必须具有唯一性，才可使用该接口，多个匹配默认返回最后匹配
  - [x] `MissionsFindByModuleID`
  - [x] `MissionUpdate`
    - 必须携带ID
  - [x] `MissionDeleteByID`
  - [x] `MissionDeleteByName`
    - Name必须具有唯一性，才可使用该接口,否则会删除最后匹配
- [x] module
  - [x] `ModuleCreate`
  - [x] `ModuleFindByID`
  - [x] `ModulesFindByLeaderID`
  - [x] `ModulesFindByProjectID`
  - [x] `ModuleUpdate`
    - id
  - [x] `ModuleDeleteByID`
- [x] project
  - [x] `ProjectCreate`
  - [x] `ProjectFindByID`
  - [x] `ProjectsFindByLeaderID`
  - [x] `ProjectsFindByParticipantID`
  - [x] `ProjectUpdate`
    - id
  - [x] `ProjectDeleteByID`
- [x] user
  - [x] `UserCreate`
  - [x] `UserBind`
  - [x] `UserFindByID`
  - [x] `UserFindByIDCard`
  - [x] `UserFindByOpenID`
  - [x] `UserDeleteByID`
  - [x] `UserDeleteByIDCarD`
  - [x] `UserUpdates`
  - [x] `UserBind`
