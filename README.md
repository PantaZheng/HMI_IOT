# 505

## 脑机协同微信公众号

### API

#### anon匿名

- url: `/anon/project`

## 需求

### 拉取名单

- 教师名单
    - get `/teacher/list`

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

### 信息登记

- 教师信息登记
    - post `/teacher/enroll`

        ```json
        {
            weChatOpenID: string//微信的识别ID
            name: string
            sex: string 男/女
            school: string
            telephone: string
        }
        ```

- 学生信息登记
    - post `/student/enroll`

        ```json
        {
            weChatOpenID: string//微信的识别ID
            name: string
            sex: string 男/女
            telephone: string
            school: string
            supervisor: string 导师
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