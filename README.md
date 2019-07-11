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
	ID        uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	State     uint   `json:"state"`
	OwnerName string `json:"ownerName" gorm:"-"`
}

type Gain struct {
	GainCore
	CreatedAt  time.Time  `json:"-"`
	CreateTime string     `json:"createTime" gorm:"-"`
	UpdatedAt  time.Time  `json:"-"`
	UpdateTime string     `json:"updateTime" gorm:"-"`
	DeletedAt  *time.Time `sql:"index" json:"-"`
	FileName   string     `json:"fileName"`
	Remark     string     `json:"remark"`

	MissionID   uint   `json:"missionID"`
	MissionName string `json:"missionName" gorm:"-"`
	OwnerID     uint   `json:"ownerID"`
	ModuleID    uint   `json:"moduleID"`
	ModuleName  string `json:"moduleName" gorm:"-"`
	LeaderID    uint   `json:"leaderID"`
	LeaderName  string `json:"leaderName "gorm:"-"`
	ProjectID   uint   `json:"projectID"`
	ProjectName string `json:"projectName" gorm:"-"`
	ManagerID   uint   `json:"managerID"`
	ManagerName string `json:"managerName" gorm:"-"`
}
```

名称|method|path|传入body参数|接收body参数
-|-|-|-|-
GainInsert|post|`/`|`Gain`+file|`Gain`
GainFindByID|get|`/id/{id:uint}`|-|`Gain`
GainsFindByMissionID|get|`/mission/{id:uint}`|-|`[]GainCore`
GainsFindByOwnerID|get|`/owner/{id:uint}`|-|`[]GainCore`
GainsFindByModuleID|get|`/module/{id:uint}`|-|`[]GainCore`
GainsFindByLeaderID|get|`/leader/{id:uint}`|-|`[]GainCore`
GainsFindByProjectID|get|`/project/{id:uint}`|-|`[]GainCore`
GainsFindByManagerID|get|`/manager/{id:uint}`|-|`[]GainCore`
GainsFindAll|get|`/all`|-|`[]GainCore`
GainDownFileByID|get|`/file/{id:uint}`|-|file
GainUpdates|put|`/`|`Gain`|`Gain`
GainDeleteByID|delete|`/{id:uint}`|-|`Gain`

### Mission

入口: `/mission`

```go
type MissionCore struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Name      string `json:"name"`
	State     uint   `json:"state"`
	OwnerName string `gorm:"-" json:"ownerName"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type Mission struct {
	MissionCore
	CreatedAt  time.Time  `json:"-"`
	CreateTime string     `json:"createTime" gorm:"-"`
	UpdatedAt  time.Time  `json:"-"`
	UpdateTime string     `gorm:"-" json:"updateTime"`
	DeletedAt  *time.Time `sql:"index" json:"-"`

	Content string `json:"content"`
	Target  string `json:"target"`

	OwnerID     uint   `json:"ownerID"`
	ModuleID    uint   `json:"moduleID"`
	ModuleName  string `gorm:"-" json:"moduleName"`
	LeaderID    uint   `json:"leaderID"`
	LeaderName  string `gorm:"-" json:"leaderName"`
	ProjectID   uint   `json:"projectID"`
	ProjectName string `gorm:"-" json:"projectName"`
	ManagerID   uint   `json:"managerID"`
	ManagerName string `gorm:"-" json:"managerName"`
}
```

名称|method|path|传入body参数|接收body参数
-|-|-|-|-
MissionInsert|post|`/`|`Mission`|`Mission`
MissionFindByID|get|`/id/{id:uint}`|-|`Mission`
MissionsFindByOwnerID|get|`/owner/{id:uint}`|-|`[]MissionCore`
MissionsFindByModuleID|get|`/module/{id:uint}`|-|`[]MissionCore`
MissionsFindByLeaderID|get|`/leader/{id:uint}`|-|`[]MissionCore`
MissionsFindByProjectID|get|`/project/{id:uint}`|-|`[]MissionCore`
MissionsFindByManagerID|get|`/manager/{id:uint}`|-|`[]MissionCore`
MissionsFindALL|get|`/all`|-|`[]MissionCore`
MissionUpdate|put|`/`|`Mission`|`Mission`
MissionDeleteByID|delete|`/id/{id:uint}`|-|`Mission`

### Module

入口: `/module`

```go
type ModuleCore struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	Name       string `json:"name"`
	State      uint   `json:"state"`
	LeaderName string `gorm:"-" json:"leaderName"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
}

type Module struct {
	ModuleCore
	CreatedAt  time.Time  `json:"-"`
	CreateTime string     `gorm:"-" json:"createTime"`
	UpdatedAt  time.Time  `json:"-"`
	UpdateTime string     `gorm:"-" json:"updateTime"`
	DeletedAt  *time.Time `sql:"index" json:"-"`

	Content string `json:"content"`
	Target  string `json:"target"`

	LeaderID    uint   `json:"leaderID"`
	ProjectID   uint   `json:"projectID"`
	ProjectName string `gorm:"-" json:"projectName"`
	ManagerID   uint   `json:"managerID"`
	ManagerName string `gorm:"-" json:"managerName"`
}
```

名称|method|path|传入body参数|接收body参数
-|-|-|-|-
ModuleInsert|post|`/`|`Module`|`Module`
ModuleFindByID|get|`/id/{id:uint}`|-|`Module`
ModulesFindByLeaderID|get|`/leader/{id:uint}`|-|`[]ModuleCore`
ModulesFindByProjectID|get|`/project/{id:uint}`|-|`[]ModuleCore`
ModulesFindByManagerID|get|`/manager/{id:uint}`|-|`[]ModuleCore`
ModulesFindByMemberID|get|`/member/{id:uint}`|-|`[]ModuleCore`
ModulesFindAll|get|`/all`|-|`[]ModuleCore`
ModuleUpdate|put|`/`|`Module`|`Module`
ModuleDeleteByID|delete|`/id/{id:uint}`|-|`Module`

### Project

入口: `/project`

```go
type ProjectCore struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Name        string `json:"name"`
	State       uint   `json:"state"`
	ManagerName string `gorm:"-" json:"managerName"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
}

type Project struct {
	ProjectCore
	CreatedAt  time.Time  `json:"-"`
	CreateTime string     `gorm:"-" json:"createTime"`
	UpdatedAt  time.Time  `json:"-"`
	UpdateTime string     `gorm:"-" json:"updateTime"`
	DeletedAt  *time.Time `sql:"index" json:"-"`

	Content string `json:"content"`
	Target  string `json:"target"`

	ManagerID uint `json:"managerID"`
}

type ModuleFrame struct {
	ModuleCore
	Missions []MissionCore `json:"missions"`
}

type ProjectFrame struct {
	ProjectCore
	Modules []ModuleFrame `json:"modules"`
}
```

名称|method|path|传入body参数|接收body参数
-|-|-|-|-
ProjectCreate|post|`/`|`Project`|`Project`
ProjectFindByID|get|`/id/{id:uint}`|-|`Project`
ProjectFrameByID|get|`/frame/{id:uint}`|-|`ProjectFrame`
ProjectsFindByManagerID|get|`/manager/{id:uint}`|-|`[]ProjectCore`
ProjectsFindByMemberID|get|`/member/{id:uint}`|-|`[]ProjectCore`
ProjectsFindALl|get|`/all`|-|`[]ProjectCore`
ProjectUpdate|put|`/`|`Project`|`Project`
ProjectDeleteByID|delete|`/id/{id:uint}`|-|`Project`
