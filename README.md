# BCI脑机协同微信公众号

## 域名

`http://bci.renjiwulian.com`

---

## 注意事项

1. 各子模块的index是单独定位的，其同目录下的js等文件，请使用`../dir/文件`格式
1. `index.html`文件不单独暴露,使用`index`暴露

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
type UserJson struct {
    ID       uint               `json:"id"`
    OpenId   string             `json:"openid"`
    Code     string             `json:"code"`
    Name     string             `json:"name"`
    Level    string             `json:"level"`
    Missions []MissionBriefJson `json:"missions"`
}

type UserBriefJson struct {
    ID uint `json:"id"`
    Name string `json:"name"`
}
```

- 登记
    - Post `/anon/enroll`
    - send
        - code不为空,其他都可为空
    - resp

        ```golang
        OpenId       string `gorm:"primary_key;unique;VARCHAR(191)" json:"openid"`
        ```

- 拉取名单
    - Get `/anon/list/{role:string}`

### Gain

入口: `/gain`

```go
type GainJson struct {
    ID          uint	`json:"id"`
    Name		string	`json:"name"`
    Type		string	`json:"type"`
    File		string	`json:"file"`
    UpTime		string	`json:"up_time"`
    Remark		string	`json:"remark"`
    OwnerID		uint	`json:"owner_id"`
    MissionID	uint	`json:"mission_id"`
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
type MissionJson struct{
	ID				uint				`json:"id"`
	Name			string				`json:"name"`
	CreatorID		uint				`json:"creator"` //不做关联，查看创建者信息，需调用user的接口
	CreateTime		string				`json:"create_time"`
	StartTime		string				`json:"start_time"`
	EndTime			string				`json:"end_time"`
	Content			string				`json:"content"`
	File			string				`json:"file"`
	Tag				bool				`json:"tag"`
	Gains			[]GainJson			`json:"gains"`
	Participants	[]UserBriefJson		`json:"participants"`
	ModuleID		uint				`json:"module"`
}

type MissionBriefJson struct{
	ID			uint	`json:"id"`
	Name		string	`json:"name"`
	CreateTime	string	`json:"create_time"`
	Content		string	`json:"content"`
	Tag			bool	`json:"tag"`
	ModuleID	uint	`json:"module"`
}
```

名称|method|path|传入body参数|接收body参数
-|-|-|-|-
MissionCreate|post|`/`|`MissionJson`|`MissionJson`
MissionFindByID|get|`/id/{id:uint}`|-|`MissionJson`
MissionFindByName|get|`/name/{name:string}`|-|`MissionJson`
MissionsFindByModuleID|get|`/module/{id:uint}`|-|`[]MissionBriefJson`
MissionUpdate|put|`/`|`MissionJson`|`MissionJson`
MissionDeleteByID|delete|`/id/{id:uint}`|-|`MissionJson`
MissionDeleteByName|delete|`/name/{name:string}`|-|`MissionJson`

### Module

入口: `/module`

```go
type ModuleJson struct{
	ID				uint				`json:"id"`
	Name			string				`json:"name"`
	CreatorID 		uint				`json:"creator"`
	CreateTime  	string	  			`json:"create_time"`//创建时间
	StartTime   	string    			`json:"start_time"`//开始时间
	EndTime     	string    			`json:"end_time"`//结束时间
	Content			string				`json:"content"`
	Tag				bool				`json:"tag"`
	ProjectID		uint				`json:"project"`
	LeaderID		uint				`json:"leader"`
	Participants 	[]UserBriefJson		`json:"participants"`//参与人员
	Missions		[]MissionBriefJson	`json:"missions"`
}

type ModuleBriefJson struct{
	ID				uint				`json:"id"`
	Name			string				`json:"name"`
	CreateTime  	string	  			`json:"create_time"`//创建时间
	Content			string				`json:"content"`
	Tag				string				`json:"tag"`
	ProjectID		uint				`json:"project"`
}
```

---

## TODO

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
    - [x] `moduleJson2Module`
    - [x] `module2ModuleJson`
    - [x] `module2ModuleBriefJson`
    - [x] `ModuleCreate`
    - [x] `ModuleFind`
    - [x] `ModulesFindByLeader`
    - [x] `ModulesFindByProject`
    - [x] `ModuleUpdate`
    - [x] `ModuleDelete`
    - [x] module_test测试通过 2019-4-30 16:30
- [ ] project
- [ ] user
- [ ] init
    - [ ] 表单删除
    - [ ] 表单迁移
    - [ ] 添加测试数据
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
    - iris测试