package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
	"github.com/kataras/iris/middleware/logger"
	"github.com/pantazheng/bci/controller"
)

func init(){
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
	api.Get("/bind/index",func(ctx iris.Context){
		_=ctx.View("/bind/index.html")
	})
	api.Get("/framework/",func(ctx iris.Context){
		_=ctx.View("/framework/index.html")
	})
	api.Get("/new/", func(ctx iris.Context) {
		_=ctx.View("/new/index.html")
	})
	api.Get("/project/", func(ctx iris.Context) {
		_=ctx.View("/project/index.html")
	})
	api.Get("/pace/", func(ctx iris.Context) {
		_=ctx.View("/pace/index.html")
	})

	api.PartyFunc("/anon",func (anon router.Party){
		anon.PartyFunc("/wechat", func(weChat router.Party) {
			weChat.Any("/", controller.WeChat)
		})
	})

	api.PartyFunc("/user",func(user router.Party){
		user.Get("/index",func(ctx iris.Context){
			_=ctx.View("/user/index.html")
		})
		user.Post("/bind",func(ctx iris.Context){

		})
	})

	api.PartyFunc("/gain",func(gain router.Party){
		gain.Post("/",controller.GainCreate)
		gain.Get("/id/{id:uint}",controller.GainFindByID)
		gain.Get("/owner/{id:uint}",controller.GainsFindByOwnerID)
		gain.Get("/mission/{id:uint}",controller.GainsFindByMissionID)
		gain.Put("/update",controller.GainUpdate)
		gain.Delete("/{id:uint}",controller.GainDeleteByID)
	})

	api.PartyFunc("/mission",func(mission router.Party){
		mission.Get("/index",func(ctx iris.Context){
			_=ctx.View("/mission/index.html")
		})
		mission.Post("/create",controller.MissionCreate)
		mission.Get("/find?id={id:uint}",controller.MissionFindByID)
		mission.Get("/find?name={name:string}",controller.MissionFindByName)
		mission.Put("/update",controller.MissionUpdate)
		mission.Delete("/delete?id={id:uint}",controller.MissionDeleteByID)
		mission.Delete("/delete?name={name:string}",controller.MissionDeleteByName)
	})
	return
}

func main() {
	app := newApp()
	_ = app.Run(iris.Addr(":80"))
}

