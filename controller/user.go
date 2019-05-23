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
	var err error
	u := new(service.UserJSON)
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
	var err error
	u := new(service.UserJSON)
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
	var err error
	if id, err1 := ctx.Params().GetUint("id"); err1 == nil {
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

func UserFindByTelephone(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 15:41
	*/
	tel := ctx.Params().GetString("telephone")
	if u, err := service.UserFindByTelephone(tel); err == nil {
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
	var err error
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
	var err error
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

func UserDeleteByTelephone(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 15:41
	*/
	tel := ctx.Params().GetString("tel")
	if userJson, err := service.UserDeleteByTelephone(tel); err == nil {
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
	var err error
	u := new(service.UserJSON)
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
