package controller

import (
	_ "github.com/kataras/iris"
	_ "github.com/pantazheng/bci/models"
	_ "github.com/pantazheng/bci/service"
)

//func Enroll(ctx iris.Context) {
//	userEnroll :=&models.User{}
//	if err:=ctx.ReadJSON(userEnroll);err!=nil{
//		panic(err.Error())
//	}
//	returnId:=&returnId{}
//	returnId.OpenId=service.Enroll(userEnroll)
//	ctx.StatusCode(iris.StatusOK)
//	if _,err:=ctx.JSON(returnId);err!=nil{
//		panic(err.Error())
//	}
//}




