package main

import (
	"./controller"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
	"github.com/kataras/iris/middleware/logger"
)

func init(){
	controller.Check()
	controller.Menu()
}

func newApp() (api *iris.Application){
	api = iris.New()
	api.Use(logger.New())

	api.OnErrorCode(404,func(ctx iris.Context){
		if _,err:=ctx.Writef("404 not found");err!=nil{
			fmt.Printf("%v",err)
		}
	})
	
	api.StaticWeb("/","./view")
	api.RegisterView(iris.HTML("./view", ".html").Delims("[[","]]"))
	api.Get("/project", func(ctx iris.Context) {
		_=ctx.View("/project/index.html")
	})
	api.Get("/mission", func(ctx iris.Context){
		_=ctx.View("/mission/index.html")
	})
	api.Get("/test", func(ctx iris.Context) {
		_=ctx.View("/test/test.html")
	})

	api.PartyFunc("/anon",func (anon router.Party){
		anon.PartyFunc("/wechat", func(weChat router.Party) {
			weChat.Any("/", controller.WeChat)
		})
	})

	api.PartyFunc("/teacher",func(teacher router.Party){
		teacher.Post("/enroll",controller.EnrollTeacher)
		teacher.Post("/purify",controller.Purify)
		teacher.Get("/list",controller.ListTeacher)
	})
	api.PartyFunc("/student",func(student router.Party){
		student.Post("/enroll",controller.EnrollStudent)
		student.Post("/purify",controller.Purify)
		student.Get("/list",controller.ListStudent)
	})


	return
}

func main() {
	app := newApp()
	_ = app.Run(iris.Addr(":80"))
}

