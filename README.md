# BCI脑机协同微信公众号

## 域名

`http://bci.renjiwulian.com`

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
	Level      int    `json:"level"`
	Telephone  string `json:"telephone"`
}
```

名称|method|path|传入body参数|接收body参数|
-|-|-|:-|:-
UserCreate|post|`/`|`UserJson`<br>IDCard,Level|`UserJson`<br>id,openid,id_card三者至少存在一个，其他项均可缺省
UserFindByID|get|`/id/{id:uint}`|-|`UserJson`|
UserFindByTelephone|get|`/telephone/{telephone:string}`|-|`UserJson`
UserFindByOpenID|get|`/openid/{openid:string}`|-|`UserJson`
UsersFindByLevel|get|`/level/{level:int}`|-|`[]UserJson`<br>仅包含id,name
UsersList|get|`/list`|-|`[]UserJson`|
UserUpdate|put|`/update`|`UserJson`|`UserJson`
UserBind|put|`/bind`|`UserJson`<br>openid,code仅且存在一个|`UserJson`
UserDeleteByID|delete|`/id/{id:uint}`|-|`UserJson`
UserDeleteByTelephone|delete|`/telephone/{telephone:string}`|-|`UserJson`

### Gain

入口: `/gain`

```go
type GainCore struct {
	ID        uint   `gorm:"primary_key",json:"id"`
	Name      string `json:"name"`
	State     uint   `json:"state"`
	OwnerName string `gorm:"-",json:"ownerName"`
}

type Gain struct {
	GainCore
	CreatedAt  time.Time  `json:"-"`
	CreateTime string     `gorm:"-",json:"createTime"`
	UpdatedAt  time.Time  `json:"-"`
	UpdateTime string     `gorm:"-",json:"updateTime"`
	DeletedAt  *time.Time `sql:"index",json:"-"`
	File       string     `json:"file"`
	Remark     string     `json:"remark"`

	MissionID   uint   `json:"missionID"`
	MissionName string `gorm:"-",json:"missionName"`
	OwnerID     uint   `json:"ownerID"`
	ModuleID    uint   `json:"moduleID"`
	ModuleName  string `gorm:"-",json:"moduleName"`
	LeaderID    uint   `json:"leaderID"`
	LeaderName  string `gorm:"-",json:"leaderName"`
	ProjectID   uint   `json:"projectID"`
	ProjectName string `gorm:"-",json:"projectName"`
	ManagerID   uint   `json:"managerID"`
	ManagerName string `gorm:"-",json:"managerName"`
}
```

名称|method|path|传入body参数|接收body参数
-|-|-|-|-
GainCreate|post|`/`|`GainJSON`|`GainJSON`
GainUpFileByID|post|`/file/{id:uint}`|file|-
GainFindByID|get|`/id/{id:uint}`|-|`GainJSON`
GainsFindByLeaderID|get|`/leader/{id:uint}`|-|`[]GainCore`
GainsFindByOwnerID|get|`/owner/{id:uint}`|-|`[]GainCore`
GainsFindByMissionID|get|`/mission/{id:uint}`|-|`[]GainCore`
GainsFindAll|get|`all`|-|`[]GainCore`
GainDownFileByID|get|`/file/{id:uint}`|-|file
GainUpdates|put|`/`|`GainJSON`|`GainJSON`
GainDeleteByID|delete|`/{id:uint}`|-|`GainJSON`

### Mission

入口: `/mission`

```go
type MissionCore struct {
	ID        uint   `gorm:"primary_key",json:"id"`
	Name      string `json:"name"`
	State     uint   `json:"state"`
	OwnerName string `gorm:"-",json:"ownerName"`
}

type Mission struct {
	MissionCore
	CreatedAt  time.Time  `json:"-"`
	CreateTime string     `gorm:"-",json:"createTime"`
	UpdatedAt  time.Time  `json:"-"`
	UpdateTime string     `gorm:"-",json:"updateTime"`
	DeletedAt  *time.Time `sql:"index",json:"-"`
	StartTime  string     `json:"startTime"`
	EndTime    string     `json:"endTime"`
	Content    string     `json:"content"`
	Target     string     `json:"target"`

	OwnerID     uint   `json:"ownerID"`
	ModuleID    uint   `json:"moduleID"`
	ModuleName  string `gorm:"-",json:"moduleName"`
	LeaderID    uint   `json:"leaderID"`
	LeaderName  string `gorm:"-",json:"leaderName"`
	ProjectID   uint   `json:"projectID"`
	ProjectName string `gorm:"-",json:"projectName"`
	ManagerID   uint   `json:"managerID"`
	ManagerName string `gorm:"-",json:"managerName"`
}
```

名称|method|path|传入body参数|接收body参数
-|-|-|-|-
MissionCreate|post|`/`|`MissionJSON`|`MissionJSON`
MissionUpFileByID|post|`/file/{id:uint}`|file|-
MissionFindByID|get|`/id/{id:uint}`|-|`MissionJSON`
MissionsFindByLeaderID|get|`/leader/{id:uint}`|-|`[]MissionCore`
MissionsFindByOwnerID|get|`/owner/{id:uint}`|-|`[]MissionCore`
MissionsFindByModuleID|get|`/module/{id:uint}`|-|`[]MissionCore`
MissionsFindALL|get|`/all`|-|`[]MissionCore`
MissionDownFileByID|get|`/file/{id:uint}`|-|file
MissionUpdate|put|`/`|`MissionJSON`|`MissionJSON`
MissionDeleteByID|delete|`/id/{id:uint}`|-|`MissionJSON`

### Module

入口: `/module`

```go
type ModuleCore struct {
	ID         uint   `gorm:"primary_key"`
	Name       string `json:"name"`
	State      uint   `json:"state"`
	LeaderName string `gorm:"-",json:"ownerName"`
}

type Module struct {
	ModuleCore
	CreatedAt  time.Time  `json:"-"`
	CreateTime string     `gorm:"-",json:"createTime"`
	UpdatedAt  time.Time  `json:"-"`
	UpdateTime string     `gorm:"-",json:"updateTime"`
	DeletedAt  *time.Time `sql:"index",json:"-"`
	StartTime  string     `json:"startTime"`
	EndTime    string     `json:"endTime"`
	Content    string     `json:"content"`
	Target     string     `json:"target"`

	LeaderID    uint   `json:"leaderID"`
	ProjectID   uint   `json:"projectID"`
	ProjectName string `gorm:"-",json:"projectName"`
	ManagerID   uint   `json:"managerID"`
	ManagerName string `gorm:"-",json:"managerName"`
}
```

名称|method|path|传入body参数|接收body参数
-|-|-|-|-
ModuleCreate|post|`/`|`ModuleJson`|`ModuleJson`
ModuleFindByID|get|`/id/{id:uint}`|-|`ModuleJson`
ModulesFindByCreatorID|get|`/id/{id:uint}`|-|`[]ModuleJson`
ModulesFindByLeaderID|get|`/leader/{id:uint}`|-|`[]ModuleJson`
ModulesFindByParticipantID|get|`/participant/{id:uint}`|-|`[]ModuleJson`
ModulesFindByProjectID|get|`/project/{id:uint}`|-|`[]ModuleJson`
ModuleUpdate|put|`/`|`ModuleJson`|`ModuleJson`
ModuleDeleteByID|delete|`/id/{id:uint}`|-|`ModuleJson`

### Project

入口: `/project`

```go
type ProjectJSON struct {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/16 15:07
	*/
	ID           uint         `json:"id"`
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	CreatorID    uint         `json:"creatorID"`
	Creator      UserJSON     `json:"creator"`
	CreateTime   string       `json:"createTime"`
	StartTime    string       `json:"startTime"`
	EndTime      string       `json:"endTime"`
	Content      string       `json:"content"`
	Target       string       `json:"target"`
	LeaderID     uint         `json:"leaderID"`
	Leader       UserJSON     `json:"leader"`
	Participants []UserJSON   `json:"participants"`
	Tag          bool         `json:"tag"` //create、update
	TagSet       []TagJson    `json:"tagSet"`
	Modules      []ModuleJSON `json:"modules"` //仅拉取更新
}

type TagJson struct {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/16 15:53
	*/
	ID  uint `json:"id"`
	Tag bool `json:"tag"`
}

type FramePaceJSON struct {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/24 0:32
	*/
	ID        uint              `json:"id"`
	Name      string            `json:"name"`
	StartTime string            `json:"startTime"`
	EndTime   string            `json:"endTime"`
	Leader    UserJSON          `json:"leader"`
	Modules   []ModuleBriefJSON `json:"modules"` //仅拉取更新

}
```

名称|method|path|传入body参数|接收body参数
-|-|-|-|-
ProjectCreate|post|`/`|`ProjectJson`|`ProjectJson`
ProjectFindByID|get|`/id/{id:uint}`|-|`ProjectJson`
ProjectFramPaceByID|get|`frame/{id:uint}`|-|`FramePaceJSON`
ProjectsFindALl|get|`/all`|-|`[]ProjectJSON`
ProjectsFindByCreatorID|get|`/creator/{id:uint}`|-|`[]ProjectBriefJson`
ProjectsFindByLeaderID|get|`/leader/{id:uint}`|-|`[]ProjectBriefJson`
ProjectssFindByParticipantID|get|`/participant/{id:uint}`|-|`[]ProjectBriefJson`
ProjectUpdate|put|`/`|`projectJson`|`ProjectJson`
ProjectDeleteByID|delete|`/id/{id:uint}`|-|`ProjectJson`
---
