package main

import (
	"./controller"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
	"github.com/kataras/iris/middleware/logger"
)

func init(){
	controller.Check()
	controller.DefauleMenu()
}

func newApp() (api *iris.Application){
	api = iris.New()
	api.Use(logger.New())

	api.RegisterView(iris.HTML("./view","html"))
	api.PartyFunc("/anon",func (anon router.Party){
		anon.PartyFunc("/wechat", func(weChat router.Party) {
			weChat.Any("/", controller.WeChat)
		})
		anon.Get("/notice",controller.Enroll)
		//anon.PartyFunc("/user",func(user router.Party){
		//	user.Post("/",)
		//})
	})

	return
}

func main() {
	app := newApp()
	_ = app.Run(iris.Addr(":80"))
}

