package controller

import (
	"github.com/kataras/iris"
	"github.com/pantazheng/bci/models"
	"github.com/pantazheng/bci/service"
	"log"
)

func GainCreate(ctx iris.Context){
	gain:=new(models.GainJson)
	if err:=ctx.ReadJSON(gain);err!=nil{
		log.Println(err.Error())
	}
	if gainJson,err:=service.GainCreate(gain);err!=nil{
		ctx.StatusCode(iris.StatusAccepted)
		_,_=ctx.Text(err.Error())
	}else{
		ctx.StatusCode(iris.StatusOK)
		_,_=ctx.JSON(gainJson)
	}
}

func GainFindByID(ctx iris.Context){
	id,_:=ctx.Params().GetUint("id")
	if gainJson,err:=service.GainFindByID(id);err!=nil{
		ctx.StatusCode(iris.StatusAccepted)
		_,_=ctx.Text(err.Error())
	}else{
		ctx.StatusCode(iris.StatusOK)
		_,_=ctx.JSON(gainJson)
	}
}

func GainsFindByOwnerID(ctx iris.Context){
	id,_:=ctx.Params().GetUint("owner_id")
	if gainsJson,err:=service.GainsFindByOwnerID(id);err!=nil{
		ctx.StatusCode(iris.StatusAccepted)
		_,_=ctx.Text(err.Error())
	}else{
		ctx.StatusCode(iris.StatusOK)
		_,_=ctx.JSON(gainsJson)
	}
}

func GainsFindByMissionID(ctx iris.Context){
	id,_:=ctx.Params().GetUint("mission_id")
	if gainsJson,err:=service.GainsFindByMissionID(id);err!=nil{
		ctx.StatusCode(iris.StatusAccepted)
		_,_=ctx.Text(err.Error())
	}else{
		ctx.StatusCode(iris.StatusOK)
		_,_=ctx.JSON(gainsJson)
	}
}

func GainUpdate(ctx iris.Context){
	gain:=new(models.GainJson)
	if err:=ctx.ReadJSON(gain);err!=nil{
		log.Println(err.Error())
	}
	if gainJson,err:=service.GainUpdate(gain);err!=nil{
		ctx.StatusCode(iris.StatusAccepted)
		_,_=ctx.Text(err.Error())
	}else{
		ctx.StatusCode(iris.StatusOK)
		_,_=ctx.JSON(gainJson)
	}
}

func GainDeleteByID(ctx iris.Context){
	id,_:=ctx.Params().GetUint("id")
	if gainJson,err:=service.GainDeleteByID(id);err!=nil{
		ctx.StatusCode(iris.StatusAccepted)
		_,_=ctx.Text(err.Error())
	}else{
		ctx.StatusCode(iris.StatusOK)
		_,_=ctx.JSON(gainJson)
	}
}