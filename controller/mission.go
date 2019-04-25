package controller

import (
	"github.com/kataras/iris"
	"github.com/pantazheng/bci/models"
	"github.com/pantazheng/bci/service"
	"log"
)

func MissionCreate(ctx iris.Context){
	mission := new(models.MissionJson)
	if err:=ctx.ReadJSON(mission);err!=nil{
		log.Println(err.Error())
	}
	if missionBriefJson,err:=service.MissionCreate(mission);err!=nil{
		ctx.StatusCode(iris.StatusAccepted)
		_,_=ctx.Text(err.Error())
	}else{
		ctx.StatusCode(iris.StatusOK)
		_,_=ctx.JSON(missionBriefJson)
	}
}

func MissionFindByID(ctx iris.Context){
	id,_:=ctx.Params().GetUint("id")
	if missionBriefJson,err:=service.MissionFindByID(id);err!=nil{
		ctx.StatusCode(iris.StatusAccepted)
		_,_=ctx.Text(err.Error())
	}else{
		ctx.StatusCode(iris.StatusOK)
		_,_=ctx.JSON(missionBriefJson)
	}
}

func MissionFindByName(ctx iris.Context){
	name:=ctx.Params().GetString("name")
	if missionBriefJson,err:=service.MissionFindByName(name);err!=nil{
		ctx.StatusCode(iris.StatusAccepted)
		_,_=ctx.Text(err.Error())
	}else{
		ctx.StatusCode(iris.StatusOK)
		_,_=ctx.JSON(missionBriefJson)
	}
}

func MissionUpdate(ctx iris.Context){
	mission := new(models.MissionJson)
	if err:=ctx.ReadJSON(mission);err!=nil{
		log.Println(err.Error())
	}
	if missionBriefJson,err:=service.MissionUpdate(mission);err!=nil{
		ctx.StatusCode(iris.StatusAccepted)
		_,_=ctx.Text(err.Error())
	}else{
		ctx.StatusCode(iris.StatusOK)
		_,_=ctx.JSON(missionBriefJson)
	}
}

func MissionDeleteByID(ctx iris.Context){
	id,_:=ctx.Params().GetUint("id")
	if missionBriefJson,err:=service.MissionDeleteByID(id);err!=nil{
		ctx.StatusCode(iris.StatusAccepted)
		_,_=ctx.Text(err.Error())
	}else{
		ctx.StatusCode(iris.StatusOK)
		_,_=ctx.JSON(missionBriefJson)
	}
}

func MissionDeleteByName(ctx iris.Context){
	name:=ctx.Params().GetString("name")
	if missionBriefJson,err:=service.MissionDeleteByName(name);err!=nil{
		ctx.StatusCode(iris.StatusAccepted)
		_,_=ctx.Text(err.Error())
	}else{
		ctx.StatusCode(iris.StatusOK)
		_,_=ctx.JSON(missionBriefJson)
	}
}