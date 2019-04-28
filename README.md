# 505

## 脑机协同微信公众号

## 域名

`http://bci.renjiwulian.com`

---

## 注意事项

1. 各子模块的index是单独定位的，其同目录下的js等文件，请使用`../dir/文件`格式
1. `index.html`文件不单独暴露,使用`index`暴露

## TODOLIST

### models

- [ ] frame
    - [ ] `type Frame`
    - [ ] `type FrameJson`
- [ ] gain
    - [x] `type Gain struct`
    - [x] `type GainJson struct`

        ```go
        type GainJson struct {
    ID            uint	`json:"id"`
    Name        string	`json:"name"`
    Type		string	`json:"type"`
    File		string	`json:"file"`
    UpTime		string	`json:"up_time"`
    Remark		string	`json:"remark"`
    OwnerID		uint	`json:"owner_id"`
    MissionID	uint	`json:"mission_id"`
}

        ```

    - [x] `func gainTestData`
    - [x] `func (gain *Gain) gainJson2Gain(gainJson *GainJson)`
    - [x] `func (gainJson *GainJson) gain2GainJson(gain *Gain)`
    - [x] `func GainCreate(gainJson *GainJson) (recordGainJson GainJson,err error)`
    - [x] `func GainFindByID(gain *Gain)(recordGainJson GainJson,err error)`
    - [x] `func GainsFindByOwner(owner *User)(gainsJson []GainJson,err error)`
    - [x] `func GainsFindByMission(mission *Mission)(gainsJson []GainJson,err error)`
    - [ ] `func GainUpdate`
    - [ ] `func GainDeleteByID`

- [ ] init
    - [ ] 表单删除
    - [ ] 表单迁移
    - [ ] 添加测试数据
- [ ] mission
- [ ] module
- [ ] project
- [ ] user



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

## API

### anon

- 微信接口
    - Any `/anon/wechat`

### user

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
