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

const (
	// LevelStranger 未绑定
	LevelStranger  = iota
	// LevelEmeritus Professor emeritus 专家教授
	LevelEmeritus
	// LevelStudent 学生
	LevelStudent
	// LevelAssistant 助理
	LevelAssistant
	// LevelSenior Senior lecturer 高级讲师
	LevelSenior
	// LevelFull Full professor 全职教授
	LevelFull
)

//UserJSON 用户Json原型
type UserJSON struct {
	ID        uint   `json:"id"`
	OpenID    string `json:"openid"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	IDCard    string `json:"idCard"`
	Level     int    `json:"level"`
	Telephone string `json:"telephone"`
}

//UserBriefJSON 简洁版的用户Json信息
type UserBriefJSON struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}
```

名称|method|path|传入body参数|接收body参数
-|-|-|-|-
UserCreate|post|`/`|`UserJson`|`UserJson`
UserFindByID|get|`/id/{id:uint}`|-|`UserJson`
UserFindByIDCard|get|`/id_card/{id_card:string}`|-|`UserJson`
UserFindByOpenID|get|`/openid/{openid:string}`|-|`UserJson`
UsersFindByLevel|get|`/level/{level:int}`|-|`[]UserBriefJson`
UserUpdate|put|`/update`|`UserJson`|`UserJson`
UserBind|put|`/bind`|`UserJson`|`UserJson`
UserDeleteByID|delete|`/id/{id:uint}`|-|`UserJson`
UserDeleteByOpenID|delete|`/id/{openid:string}`|-|`UserJson`

### Gain

入口: `/gain`

```go
type GainJson struct {
	ID        uint          `json:"id"`
	Name      string        `json:"name"`
	Type      string        `json:"type"`
	File      string        `json:"file"`
	UpTime    string        `json:"upTime"`
	Remark    string        `json:"remark"`
	Owner     UserBriefJSON `json:"owner"`
	MissionID uint          `json:"missionID"`
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

- [ ] 单表JSON数据格式合并，在service层做置空处理
- [ ] 微信在service层面的对后台消息的反馈
- [ ] 操作函数`er`命名

2019-5-8

- [ ] PACE进度实现
- [ ] FRAME框架实现
- [ ] 之前的实例化方式都是指针创建，改用对象
- [ ] level改用`map`实现
- [ ] createTime不用加time,只用data
- [ ] models、service采用`method`重写`function`，----
- [ ] 使用`interface`写不同等级用户，允许拥有不同的权限
- [ ] mysql 使用`utf8mb4`，用户表加入微信用户名
  - `mb4`: most bytes 4
- [ ] 操作系统将文件重定向到文件
- [ ] JSON数据做忽略零值处理`omitempty`
- [x] JSON命名风格改为骆驼风格
- [ ] 时间风格，在utils中添加转换时间的工具
  - `layout: 2006-01-02`

  ```bash
  > .log 2>&1
  ```

### models

- [x] gain
  - [x] `type Gain struct`
  - [x] `type GainJson struct`
  - [x] `gainTestData`
  - [x] `gainJson2Gain`
  - [x] `gain2GainJson`
  - [x] `GainCreate`
  - [x] `GainFind`
  - [x] `GainsFindByOwner`
  - [x] `GainsFindByMission`
  - [x] `GainUpdate`
    - 必须携带ID
    - 目前只允许通过ID删除单条记录
    - UpTime更新为当前时间
  - [x] `GainDelete`
    - 必须携带ID
    - 目前由于只允许通过ID进行删除单条记录
- [x] mission
  - [x] `type Mission struct`
  - [x] `type MissionJson struct`
  - [x] `type MissionBriefJson struct`
  - [x] `missionTestData`
  - [x] `missionJson2Mission`
  - [x] `mission2MissionJSON`
  - [x] `mission2MissionBriefJSON`
  - [x] `MissionCreate`
  - [x] `MissionFind`
  - [x] `MissionsFindByModule`
  - [x] `MissionUpdate`
  - [x] `MissionDelete`
- [x] module
  - [x] `type Module struct`
  - [x] `type ModuleJson struct`
  - [x] `type ModuleBriefJson struct`
  - [x] `moduleTestData`
  - [x] `moduleJson2Module`
  - [x] `module2ModuleJson`
  - [x] `module2ModuleBriefJson`
  - [x] `ModuleCreate`
  - [x] `ModuleFind`
  - [x] `ModulesFindByLeader`
  - [x] `ModulesFindByProject`
  - [x] `ModuleUpdate`
  - [x] `ModuleDelete`
- [x] project
  - [x] `type Project struct`
  - [x] `type ProjectJson struct`
  - [x] `type BriefProject struct`
  - [x] `type TagJson struct`
  - [x] `projectTestData`
  - [x] `target2TargetsJson`
  - [x] `targetsJson2Target`
  - [x] `tagSet2TagsJson`
  - [x] `tagsJson2TagSet`
  - [x] `projectJson2Project`
  - [x] `project2ProjectJson`
  - [x] `project2ProjectBriefJson`
  - [x] `ProjectCreate`
  - [x] `ProjectFind`
  - [x] `ProjectsFindByLeader`
  - [x] `ProjectsFindByParticipant`
  - [x] `ProjectUpdate`
  - [x] `ProjectDelete`
  - [x] `../project_test`
- [x] user
  - [x] `type User struct`
  - [x] `type UserJson struct`
  - [x] `type UserBriefJson struct`
  - [x] `userTestData`
  - [x] `userJson2User`
  - [x] `user2UserJson`
  - [x] `user2UserBriefJson`
  - [x] `UserCreate`
  - [x] `UserFind`
  - [x] `UsersFindByLevel`
  - [x] `UserUpdate`
  - [x] `UserDelete`
- [x] init
  - [x] 表单删除
  - [x] 表单迁移
  - [x] 添加测试数据
- [ ] frame

### servcie

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
  - [x] `UserInitByWechat`
  - [x] `UserCreate`
  - [x] `UserUpdate`
  - [x] `UserFindByID`
  - [x] `UserFindByIDCard`
  - [x] `UserFindByOpenID`
  - [x] `UserDeleteByID`
  - [x] `UserDeleteByOpenID`
  - [x] `UserBind`

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
  - [x] `UserFindByID`
  - [x] `UserFindByIDCard`
  - [x] `UserFindByOpenID`
  - [x] `UserDeleteByID`
  - [x] `UserDeleteByOpenID`
  - [x] `UserUpdate`
  - [x] `UserBind`
