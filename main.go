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

	api.OnErrorCode(iris.StatusNotFound,func(ctx iris.Context){
		if _,err:=ctx.Writef("404 not found");err!=nil{
			fmt.Printf("%v",err)
		}
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
		gain.Put("/",controller.GainUpdate)
		gain.Delete("/{id:uint}",controller.GainDeleteByID)
	})

	api.PartyFunc("/mission",func(mission router.Party){
		mission.Post("/",controller.MissionCreate)
		mission.Get("/id/{id:uint}",controller.MissionFindByID)
		mission.Get("/name/{name:string}",controller.MissionFindByName)
		mission.Get("/module/{id:uint}",controller.MissionsFindByModuleID)
		mission.Put("/",controller.MissionUpdate)
		mission.Delete("/id/{id:uint}",controller.MissionDeleteByID)
		mission.Delete("/name/{name:string}",controller.MissionDeleteByName)
	})

	api.PartyFunc("/module",func(module router.Party){
		module.Post("/",controller.ModuleCreate)
		module.Get("/id/{id:uint}",controller.ModuleFindByID)
		module.Get("/leader/{id:uint}",controller.ModulesFindByLeaderID)
		module.Get("/project/{id:uint}",controller.ModulesFindByProjectID)
		module.Put("/",controller.ModuleUpdate)
		module.Delete("/id/{id:uint}",controller.ModuleDeleteByID)
	})

	return
}

func main() {
	app := newApp()
	_ = app.Run(iris.Addr(":80"))
}

