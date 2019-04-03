package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
	"github.com/kataras/iris/middleware/logger"
	"./controller/anon"
)

func newApp() (api *iris.Application){
	api = iris.New()
	api.Use(logger.New())

	api.PartyFunc("/anon",func (anon router.Party){
		anon.PartyFunc("/wechat", func(wechat router.Party) {
			wechat.Get("/",controller.Login)
			wechat.Post("/",controller.Port)
		})
	})

	return
}

func main() {
	app := newApp()
	_ = app.Run(iris.Addr(":80"))
	fmt.Printf("Hello 8080")
}

