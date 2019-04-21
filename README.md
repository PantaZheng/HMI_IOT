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

[菜单预览](https://pantazheng.github.io/HMI_IOT/design/index.html)

### 菜单接口

- 人员
    1. 绑定 click
    1. 架构 view    `/frame/`
- 内容
    1. 新建 view    `/new/`
    1. 项目 view    `/project/`
    1. 任务 view    `/mission/`
- 进度      view    `/pace/`

### 权限

五级权限

1. 教师助理
    : **level1** (拥有所有权限)
    1. 查看所有项目及项目下的模块、任务
    1. 新增/删除/修改 项目/模块/任务
    1. 查看任务下提交上来的所有文件: 论文、专利、著作、成果、期刊
1. PI
    : **level2** 戴国骏...
    1. 查看所有项目及项目下的模块、任务
    1. 查看任务下提交上来的所有文件: 论文、专利、著作、成果、期刊
1. PR
    : **level3** 曾虹，周文晖，张桦，韩
    1. 只能查看到自己参与的项目
        1. 项目下自己管理的项目可以看到其下的所有任务及成果
        1. <mark>项目下非自己管理的模块能看到任务及完成进度</mark>，当不能看到任务下的成果(论文、专利、著作、成果、期刊)
    1. 新增/删除/修改 项目/模块/任务
    1. <mark>进度，只能看到自己参与项目的进度</mark>
1. 学生
    :  **level4**
    1. 只能看到自己参与的项目，只有查看功能，不能增删改
    1. 进度，只能看到自己参与的项目的进度
1. 专家教授
    : **level5**

    1. 需绑定，绑定为查看模式
    1. 能看到所有项目及详情，只有查看功能
    1. 可以看到项目，任务，模块等大概的进度，不能看到上传的文件成果

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



