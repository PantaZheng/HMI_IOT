package controller

import (
	"github.com/kataras/iris"
	"github.com/pantazheng/bci/models"
	"github.com/pantazheng/bci/service"
)


type returnId struct {
	OpenId string `json:"openid"`
}

func Enroll(ctx iris.Context) {
	userEnroll :=&models.User{}
	if err:=ctx.ReadJSON(userEnroll);err!=nil{
		panic(err.Error())
	}
	returnId:=&returnId{}
	returnId.OpenId=service.Enroll(userEnroll)
	ctx.StatusCode(iris.StatusOK)
	if _,err:=ctx.JSON(returnId);err!=nil{
		panic(err.Error())
	}
}

//func List(ctx iris.Context){
//	role:= ctx.Params().GetString("role")
//	memberList :=service.GetMembers(role)
//	if _,err:=ctx.JSON(memberList);err!=nil{
//		panic(err.Error())
//	}
//}


