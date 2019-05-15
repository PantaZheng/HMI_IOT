package controller

import (
	"errors"
	"github.com/kataras/iris"
	"github.com/pantazheng/bci/service"
	"log"
)

func UserCreate(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 23:48
	*/
	u := new(service.UserJSON)
	err := *new(error)
	if err = ctx.ReadJSON(u); err == nil {
		if u.WechatName != "" || u.OpenID != "" {
			err = errors.New("UserCreate接口不支持微信信息")
		} else if err = u.Create(); err == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(u)
		}
	}
	if err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		info := err.Error()
		_, _ = ctx.Text(info)
		log.Println(info)
	}
}

func UserBind(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 15:41
	*/
	u := new(service.UserJSON)
	err := *new(error)
	if err = ctx.ReadJSON(u); err == nil {
		if err = u.Bind(); err == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(u)
		}
	}
	if err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		info := err.Error()
		_, _ = ctx.Text(info)
		log.Println(info)
	}
}

func UserFindByID(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 15:41
	*/
	err := *new(error)
	if id, err1 := ctx.Params().GetUint("id"); err == nil {
		if u, err2 := service.UserFindByID(id); err2 == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(u)
		} else {
			err = err2
		}
	} else {
		err = err1
	}
	if err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		info := err.Error()
		_, _ = ctx.Text(info)
		log.Println(info)
	}
}

func UserFindByIDCard(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 15:41
	*/
	idCard := ctx.Params().GetString("id_card")
	if u, err := service.UserFindByIDCard(idCard); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(u)
	} else {
		ctx.StatusCode(iris.StatusAccepted)
		info := err.Error()
		_, _ = ctx.Text(info)
		log.Println(info)
	}
}

func UserFindByOpenID(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 15:41
	*/
	openid := ctx.Params().GetString("openid")
	if u, err := service.UserFindByOpenID(openid); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(u)
	} else {
		ctx.StatusCode(iris.StatusAccepted)
		info := err.Error()
		_, _ = ctx.Text(info)
		log.Println(info)
	}
}

func UsersFindByLevel(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 15:41
	*/
	err := *new(error)
	if level, err1 := ctx.Params().GetInt("level"); err1 == nil {
		if usersJson, err2 := service.UsersFindByLevel(level); err2 == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(usersJson)
		} else {
			err = err2
		}
	} else {
		err = err1
	}
	if err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		info := err.Error()
		_, _ = ctx.Text(info)
		log.Println(info)
	}
}

func UserDeleteByID(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 15:41
	*/
	err := *new(error)
	if id, err1 := ctx.Params().GetUint("id"); err1 == nil {
		if userJson, err2 := service.UserDeleteByID(id); err2 == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(userJson)
		} else {
			err = err2
		}
	} else {
		err = err1
	}
	if err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		info := err.Error()
		_, _ = ctx.Text(info)
		log.Println(info)
	}
}

func UserDeleteByIDCard(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 15:41
	*/
	idCard := ctx.Params().GetString("id_card")
	if userJson, err := service.UserDeleteByIDCard(idCard); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(userJson)
	} else {
		ctx.StatusCode(iris.StatusAccepted)
		info := err.Error()
		_, _ = ctx.Text(info)
		log.Println(info)
	}
}

func UserUpdates(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 15:41
	*/
	u := new(service.UserJSON)
	err := *new(error)
	if err = ctx.ReadJSON(u); err == nil {
		if err = u.Updates(); err == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(u)
		}
	}
	if err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		info := err.Error()
		_, _ = ctx.Text(info)
		log.Println(info)
	}
}
