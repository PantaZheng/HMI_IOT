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

	api.Favicon("./view/favicon.ico")
	api.OnErrorCode(404,func(ctx iris.Context){
		if _,err:=ctx.Writef("404 not found");err!=nil{
			fmt.Printf("%v",err)
		}
	})


	api.StaticWeb("/css","./view/css")
	api.StaticWeb("/excel","./view/css")
	api.StaticWeb("/fonts","./view/fonts")
	api.StaticWeb("/image","./view/image")
	api.StaticWeb("/images","./view/images")
	api.StaticWeb("/js","./view/js")
	api.StaticWeb("/mission","./view/mission")
	api.StaticWeb("/project","./view/project")
	api.StaticWeb("/scripts","./view/scripts")
	api.StaticWeb("/test","./view/test")
	api.StaticWeb("/weekly","./view/weekly")
	
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

