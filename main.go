package main

import (
	"./controller"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
	"github.com/kataras/iris/middleware/logger"
)

func init(){
	controller.Check()
	controller.DefaultMenu()
}

func newApp() (api *iris.Application){
	api = iris.New()
	api.Use(logger.New())

	api.RegisterView(iris.HTML("./view",".html").Delims("[[","]]"))
	api.PartyFunc("/anon",func (anon router.Party){
		anon.PartyFunc("/wechat", func(weChat router.Party) {
			weChat.Any("/", controller.WeChat)
		})
	})
	api.PartyFunc("/teacher",func(teacher router.Party){
		teacher.Post("/enroll",controller.EnrollTeacher)
		teacher.Get("/list",controller.ListTeacher)
	})
	api.PartyFunc("/student",func(student router.Party){
		student.Post("/enroll",controller.EnrollStudent)
		student.Get("/list",controller.ListStudent)
	})
	api.PartyFunc("/project",func (project router.Party){
		project.Get("/",func(ctx iris.Context){
			_ = ctx.View("project/detail.html")
		})
	})

	return
}

func main() {
	app := newApp()
	_ = app.Run(iris.Addr(":80"))
}

