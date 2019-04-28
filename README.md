# 505

## 脑机协同微信公众号

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

Json:

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

- 微信接口
    - Any `/anon/project`
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
	ID				uint             `json:"id"`
	Name			string           `json:"name"`
	Creator			string           `json:"creator"`
	CreateTime		string           `json:"create_time"`
	StartTime		string           `json:"start_time"`
	EndTime			string           `json:"end_time"`
	Content			string           `json:"content"`
	File			string           `json:"file"`
	Tag				bool             `json:"tag"`
	Participants	[]*UserBriefJson `json:"participants"`
	ModuleID        uint             `json:"module"`
}

type MissionBriefJson struct{
	ID         uint		`json:"id"`
	Name       string	`json:"name"`
	CreateTime string	`json:"create_time"`
	Content    string	`json:"content"`
	Tag		   string	`json:"tag"`
}
```

---

## TODOLIST

### models

- [ ] gain
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
- [ ] mission
    - [x]
- [ ] module
- [ ] project
- [ ] user
- [ ] init
    - [ ] 表单删除
    - [ ] 表单迁移
    - [ ] 添加测试数据
- [ ] frame

### servcie

- [ ] gain
    - [x] `GainCreate`
    - [x] `GainFindByID`
    - [x] `GainsFindByOwnerID`
    - [x] `GainsFindByMissionID`
    - [x] `GainUpdate`
        - 必须携带ID
    - [x] `GainDeleteByID`
        - 必须携带ID

### controller

- [ ] gain
    - [x] `GainCreate`
    - [x] `GainFindByID`
    - [x] `GainsFindByOwnerID`
    - [x] `GainsFindByMissionID`
    - [x] `GainUpdate`
        - 必须携带ID
    - [x] `GainDeleteByID`
        - 必须携带ID
