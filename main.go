package main

import (
	"./controller/anon"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
	"github.com/kataras/iris/middleware/logger"
)

func newApp() (api *iris.Application){
	api = iris.New()
	api.Use(logger.New())

	api.PartyFunc("/anon",func (anon router.Party){
		anon.PartyFunc("/wechat", func(wechat router.Party) {
			wechat.Any("/", controller.Wechat)
		})
		anon.PartyFunc("/user",func(user router.Party){
			user.Post("/",)
		})
	})

	return
}

func main() {
	app := newApp()
	_ = app.Run(iris.Addr(":80"))
}

