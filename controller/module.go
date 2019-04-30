package controller

import (
	"github.com/kataras/iris"
	"github.com/pantazheng/bci/models"
	"github.com/pantazheng/bci/service"
	"log"
)

func ModuleCreate(ctx iris.Context){
	module := new(models.ModuleJson)
	if err:=ctx.ReadJSON(module);err!=nil{
		log.Println(err.Error())
	}
	if moduleJson,err:=service.ModuleCreate(module);err!=nil{
		ctx.StatusCode(iris.StatusAccepted)
		_,_=ctx.Text(err.Error())
	}else{
		ctx.StatusCode(iris.StatusOK)
		_,_=ctx.JSON(moduleJson)
	}
}

func ModuleFindByID(ctx iris.Context){
	id,_:=ctx.Params().GetUint("id")
	if moduleJson,err:=service.ModuleFindByID(id);err!=nil{
		ctx.StatusCode(iris.StatusAccepted)
		_,_=ctx.Text(err.Error())
	}else{
		ctx.StatusCode(iris.StatusOK)
		_,_=ctx.JSON(moduleJson)
	}
}

func ModlesFindByLeaderID(ctx iris.Context){
	id,_:=ctx.Params().GetUint("id")
	if modulesJson,err:=service.ModulesFindByLeaderID(id);err!=nil{
		ctx.StatusCode(iris.StatusAccepted)
		_,_=ctx.Text(err.Error())
	}else{
		ctx.StatusCode(iris.StatusOK)
		_,_=ctx.JSON(modulesJson)
	}
}


func ModulesFindByProjectID(ctx iris.Context){
	id,_:=ctx.Params().GetUint("id")
	if modulesJson,err:=service.ModulesFindByProjectID(id);err!=nil{
		ctx.StatusCode(iris.StatusAccepted)
		_,_=ctx.Text(err.Error())
	}else{
		ctx.StatusCode(iris.StatusOK)
		_,_=ctx.JSON(modulesJson)
	}
}

func ModuleUpdate(ctx iris.Context){
	module := new(models.ModuleJson)
	if err:=ctx.ReadJSON(module);err!=nil{
		log.Println(err.Error())
	}
	if moduleJson,err:=service.ModuleUpdate(module);err!=nil{
		ctx.StatusCode(iris.StatusAccepted)
		_,_=ctx.Text(err.Error())
	}else{
		ctx.StatusCode(iris.StatusOK)
		_,_=ctx.JSON(moduleJson)
	}
}

func ModuleDeleteByID(ctx iris.Context){
	id,_:=ctx.Params().GetUint("id")
	if moduleJson,err:=service.ModuleDeleteByID(id);err!=nil{
		ctx.StatusCode(iris.StatusAccepted)
		_,_=ctx.Text(err.Error())
	}else{
		ctx.StatusCode(iris.StatusOK)
		_,_=ctx.JSON(moduleJson)
	}
}
