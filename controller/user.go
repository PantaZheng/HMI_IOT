package controller

import (
	"github.com/kataras/iris"
	"github.com/pantazheng/bci/models"
	"github.com/pantazheng/bci/service"
	"log"
)

func UserCreate(ctx iris.Context) {
	newUser := new(models.UserJSON)
	if err := ctx.ReadJSON(newUser); err != nil {
		log.Println(err.Error())
	}
	if userJson, err := service.UserCreate(newUser); err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		_, _ = ctx.Text(err.Error())
	} else {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(userJson)
	}
}

func UserFindByID(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")
	if userJson, err := service.UserFindByID(id); err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		_, _ = ctx.Text(err.Error())
	} else {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(userJson)
	}
}

func UserFindByIDCard(ctx iris.Context) {
	idCard := ctx.Params().GetString("id_card")
	if userJson, err := service.UserFindByIDCard(idCard); err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		_, _ = ctx.Text(err.Error())
	} else {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(userJson)
	}
}

func UserFindByOpenID(ctx iris.Context) {
	openid := ctx.Params().GetString("openid")
	if userJson, err := service.UserFindByOpenID(openid); err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		_, _ = ctx.Text(err.Error())
	} else {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(userJson)
	}
}

func UsersFindByLevel(ctx iris.Context) {
	level, _ := ctx.Params().GetInt("level")
	if usersJson, err := service.UsersFindByLevel(level); err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		_, _ = ctx.Text(err.Error())
	} else {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(usersJson)
	}
}

func UserDeleteByID(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")
	if userJson, err := service.UserDeleteByID(id); err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		_, _ = ctx.Text(err.Error())
	} else {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(userJson)
	}
}

func UserDeleteByOpenID(ctx iris.Context) {
	id := ctx.Params().GetString("openid")
	if userJson, err := service.UserDeleteByOpenID(id); err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		_, _ = ctx.Text(err.Error())
	} else {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(userJson)
	}
}

func UserUpdate(ctx iris.Context) {
	newUser := new(models.UserJSON)
	if err := ctx.ReadJSON(newUser); err != nil {
		log.Println(err.Error())
	}
	if userJson, err := service.UserUpdate(newUser); err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		_, _ = ctx.Text(err.Error())
	} else {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(userJson)
	}
}

func UserBind(ctx iris.Context) {
	newUser := new(models.UserJSON)
	if err := ctx.ReadJSON(newUser); err != nil {
		log.Println(err.Error())
	}
	if userJson, err := service.UserBind(newUser); err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		_, _ = ctx.Text(err.Error())
	} else {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(userJson)
	}
}
