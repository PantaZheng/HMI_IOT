# 505

## 脑机协同微信公众号

## 域名

`http://bci.renjiwulian.com`

## 后台逻辑

1. 用户关注公众号
1. 登记信息
1. 退出一下刷新菜单

---

## 前端注意事项

1. 各子模块的index是单独定位的，其同目录下的js等文件，请使用`../dir/文件`格式

---

[ ] 1.user的关系定义
[ ] 1.登记信息是单导师负责，登记项目，多导师，多学生

## 默认测试数据

- 用户测试数据
    1. `OpenId:"test1",Role:"unEnrolled"`
    1. `OpenId:"test2",Role:"unEnrolled"`
    1. `OpenId:"test3",Role:"unEnrolled"`
    1. `OpenId:"student1",Name:"student1",Role:"student",Supervisor:"teacher1"`
    1. `OpenId:"student2",Name:"student2",Role:"student",Supervisor:"teacher1"`
    1. `OpenId:"student3",Name:"student3",Role:"student",Supervisor:"teacher2"`
    1. `OpenId:"teacher1",Name:"戴国骏",Role:"teacher"`
    1. `OpenId:"teacher2",Name:"张桦",Role:"teacher"`
    1. `OpenId:"teacher_unknown",Name:"其他导师",Role:"teacher"`

## API

### anon

- 微信接口
    - Any `/anon/project`
- 登记
    - Post `/anon/enroll`
    - send

        ```golang
        OpenId       string `gorm:"primary_key;unique;VARCHAR(191)" json:"openid"`
        Code         string `gorm:"not null VARCHAR(255)" json:"code"`
        Name         string `gorm:"not null VARCHAR(255)" json:"name"`
        Sex          string `gorm:"not null VARCHAR" json:"sex"`
        Role         string `gorm:"not null VARCHAR(191)" json:"role"`
        School       string `gorm:"not null VARCHAR(255)" json:"school"`
        Supervisor   string `gorm:"not null VARCHAR(191)" json:"supervisor"`
        ```

        - code不为空,其他都可为空
    - resp

        ```golang
        OpenId       string `gorm:"primary_key;unique;VARCHAR(191)" json:"openid"`
        ```

- 拉取名单
    - Get `/anon/list/{role:string}`

## TODO

### 菜单接口

[菜单详情可见](/design/index.html)

1. 绑定 click
1. 架构 view	host/frame/index.html
1. 新建 view    host/new/index.html
1. 项目 view	host/project/index.html
1. 任务 view	host/mission/index.html
1. 进度 view	host/pace/index.html

