package controller

import (
	"../models"
	"../service"
	"github.com/kataras/iris"
)

func Check() {
	service.CheckTableUser()
}

func Enroll(ctx iris.Context) {
	userEnroll :=&models.User{}
	if err:=ctx.ReadJSON(userEnroll);err!=nil{
		panic(err.Error())
	}
	user:=&models.User{}
	user.OpenId=service.Enroll(userEnroll)
	if _,err:=ctx.JSON(user.OpenId);err!=nil{
		panic(err.Error())
	}
}

func List(ctx iris.Context){
	role:= ctx.Params().GetString("role")
	memberList :=service.GetMembers(role)
	if _,err:=ctx.JSON(memberList);err!=nil{
		panic(err.Error())
	}
}


