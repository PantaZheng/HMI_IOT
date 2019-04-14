# 505

## 脑机协同微信公众号

## 域名

`http://bci.renjiwulian.com`

## 后台逻辑

1. 用户关注公众号，默认菜单
1. 登记信息
1. 退出一下刷新菜单
1. 再进入，个性化菜单

---

## 前端注意事项

1. 各子模块的index是单独定位的，其同目录下的js等文件，请使用`../dir/文件`格式

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

### teacher

- 拉取教师名单
    - Get `/teacher/list`

        ```json
        {
            {
                id: uint
                name: string
            },
            {
                ..
                ..
            },
            ...
        }
        ```

- 教师信息登记
    - Post `/teacher/enroll`

        ```json
        {
            code: string
            name: string
            sex: string
            school: string
            telephone: string
        }
        ```

- 取消教师身份
    - Post `/teacher/purify`

        ```json
        {
            openid: string
        }
        ```

### student

- 拉取学生名单
    - Get `/student/list`

        ```json
        {
            {
                id: uint
                name: string
            },
            {
                ..
                ..
            },
            ...
        }
        ```

- 学生信息登记
    - post `/student/enroll`

        ```json
        {
            code: string
            name: string
            sex: string
            telephone: string
            school: string
            supervisor: string
        }
        ```

- 取消学生身份
    - Post `/teacher/purify`

        ```json
        {
            openid: string
        }
        ```

### 项目/任务

- 获取名下所有项目
    - get `project/all`
- 获取项目详情:
    - get `project/details{projectID}`

        ```json
        {
            time://创建时间
            title://项目题目
            content://项目详细内容
            []members://参与人员
            leader://负责人员
            startTime://开始时间
            endTime：//结束时间
        }
        ```

- 新建项目
    - 教师
    - `/project/new`
    - post:

        ```json
        {
            teacher:
            time:
            title:
            content:
            members:
            leader:
            startTime:
            endTime:
        }
        ```