# 505

## 脑机协同微信公众号

## 域名

`http://bci.renjiwulian.com`

---

## 注意事项

1. 各子模块的index是单独定位的，其同目录下的js等文件，请使用`../dir/文件`格式
1. `index.html`文件不单独暴露

---

## TODO

## 思路

1. 简化绑定功能，重点在查看管理
1. 该公众号主要面向管理层，主要是实现各模块的查看功能，清晰简洁，弱化交互功能
1. 权限控制
    1. 前端根据用户的level字段进行操作，对不同权限展示不同的界面
1. 后台如何实现权限控制
    1. 如何输出一个学生参与的所有项目列表
        - 建表和查找问题： 一对多关系建立
    1. 新建项目时，被选定为参与人员的学生，后台所需操作
1. 项目ID，模块ID，任务ID的设置
    - 如果模块ID中包含项目的信息，则可简化查询
    - 查询任务时就不需要借助PId和moId
    例子：

    pId|moId|miId
    -|-|-
    A|A1|A11
    B|B2|B21
1. 输出某个学生参与的所有任务，是依次去项目列表里面对比学生Id，还是应该新建一个表
    1. 为每个学生和用户新建一张表~~用户表中带项目列表~~
1. 模块的结束~~what~~
    1. 时间到了自动置为结束状态
    1. 模块下各任务均为结束状态则为结束

### 权限

五级权限

1. 教师助理: **level1** (拥有所有权限)
    1. 查看所有项目及项目下的模块、任务
    1. 新增/删除/修改 项目/模块/任务
    1. 查看任务下提交上来的所有文件: 论文、专利、著作、成果、期刊
1. PI: **level2** 戴国骏...
    1. 查看所有项目及项目下的模块、任务
    1. 查看任务下提交上来的所有文件: 论文、专利、著作、成果、期刊
1. PR: **level3** 曾虹，周文晖，张桦，韩
    1. 只能查看到自己参与的项目
        1. 项目下自己管理的项目可以看到其下的所有任务及成果
        1. ~~项目下非自己管理的模块能看到任务及完成进度~~，当不能看到任务下的成果(论文、专利、著作、成果、期刊)
    1. 新增/删除/修改 项目/模块/任务
    1. ~~进度，只能看到自己参与项目的进度~~
1. 学生:  **level4**
    1. 只能看到自己参与的项目，只有查看功能，不能增删改
    1. 进度，只能看到自己参与的项目的进度
1. 专家教授: **level5**
    1. 需绑定，绑定为查看模式
    1. 能看到所有项目及详情，只有查看功能
    1. 可以看到项目，任务，模块等大概的进度，不能看到上传的文件成果

### 菜单

[菜单预览](https://pantazheng.github.io/HMI_IOT/design/index.html)

- 人员
    1. 绑定 click
    1. 架构 view    `/frame/`
- 内容
    1. 新建 view    `/new/`
    1. 项目 view    `/project/`
    1. 任务 view    `/mission/`
- 进度      view    `/pace/`

#### 绑定

~~绑定按钮做什么？？？~~

绑定功能
: 用户直接发送姓名进行绑定

1. 数据库中预存参与到项目中的用户信息 `member.json`
1. 用户关注时，发送消息“欢迎关注脑机协同之项目管理，请正确输入您的姓名进行绑定”~~姓名输入的实现与电话号码~~
1. 用户点击菜单时，若是未绑定状态，则发送提示“请正确输入您的姓名进行绑定”
1. 该功能由后台直接实现，用户由键盘输入姓名，实现以下功能：
    1. 输入的姓名存在于数据库中，则获取用户openid，并存入openid字段，完成绑定，并查看该姓名的level字段，根据level字段进行操作
        - level1: “欢迎使用脑机协同之项目管理公众号，您已绑定为助理模式，请输入您的电话号码”
        - level2, level3: “欢迎使用脑机协同之项目管理公众号，您已绑定为管理员模式，请输入您的电话号码”
        - level4：“欢迎使用脑机协同之项目管理公众号，您已绑定为学生模式，请输入您的电话号码”
    1. 输入的姓名不存在于数据库中，则默认为level4，回复“欢迎使用脑机协同之项目管理公众号，您已绑定为查看模式”
    1. 输入的姓名已被绑定，提示“该姓名以被绑定，请确认以后再次输入。如有疑问，请联系管理员帮您解决，联系方式‘123456’”
        - 以防出现，数据库中名字存储错误的情况或名字被别人绑定的情况
        - ~~重名问题~~
    1. 若已绑定，再次绑定，视为更新

绑定设计四种模式~~在哪实现？~~

1. 助理模式：level1的助理
1. 管理员模式：level2，level3的老师
1. 学生模式：level4的学生
1. 查看模式：level5的专家教授，他们并不参与到项目中来，但需要查看项目进度

```go
//example member.json
[
    {
        "name":"老师助理",
        "tel":"",   //要求加上电话号码字段
        "level":"1",    //1级权限。能查看所有，无增删改
        "openid":"null"//openid为null，则表示该用户未绑定，绑定之后为用户的openid
    },{
        "name":"戴国骏",
        "tel":"",
        "level":"2",//2级权限能查看自己参与的项目以及其下所有模块任务，可新增模块，任务
        "openid":"null"//openid为null，则表示该用户未绑定
    },{
        "name":"曾虹",
        "tel":"",
        "level":"3", //3级权限可查看所有项目，可增删改项目，模块，任务
        "openid":"null"//openid为null，则表示该用户未绑定
    },{
        "name":"学生",
        "tel":"",
        "level":"4",  //4级权限能查看自己参与的项目，无增删改查
        "openid":"null"//openid为null，则表示该用户未绑定
    }
]
```

```go
//professors.json
["name","name","name"]
//level5的专家教授,输入姓名即可,无需电话
```

#### 架构

1. 架构点进去是项目列表（权限能看到的所有项目），选择项目查看该项目的人员架构
1. 人员架构根据项目创建时选择的总负责人，参与的老师，老师带领的学生 建表

API:

- 获取项目列表 get   `/projectList` projectList.json
- 获取人员架构 get   `/frame?pId={pId}`  frame.json


---

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

        ```go
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



