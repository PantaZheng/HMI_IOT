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
	api.Get("/project/index.html", func(ctx iris.Context) {
		_=ctx.View("/project/index.html")
	})
	api.Get("/mission/index.html", func(ctx iris.Context){
		_=ctx.View("/mission/index.html")
	})
	api.Get("/createUser/", func(ctx iris.Context) {
		_=ctx.View("/createUser/index.html")
	})


	api.PartyFunc("/anon",func (anon router.Party){
		anon.PartyFunc("/wechat", func(weChat router.Party) {
			weChat.Any("/", controller.WeChat)
		})
		anon.Post("/enroll",controller.Enroll)
		anon.Get("/list/{role:string}",controller.List)
	})

	api.PartyFunc("/teacher",func(teacher router.Party){
		//teacher.Get("/list",controller.GetTeachers)
	})
	api.PartyFunc("/student",func(student router.Party){
		//student.Get("/list",controller.GetStudents)
	})


	return
}

func main() {
	app := newApp()
	_ = app.Run(iris.Addr(":80"))
}

