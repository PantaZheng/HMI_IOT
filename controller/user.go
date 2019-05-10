package controller

import (
	"github.com/kataras/iris"
	"github.com/pantazheng/bci/service"
	"log"
)

const title = "controller.user."

func UserCreate(ctx iris.Context) {
	u := new(service.UserJSON)
	if err := ctx.ReadJSON(u); err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		info := err.Error()
		_, _ = ctx.Text(info)
		log.Println(info)
	} else {
		if u.WechatName != "" || u.OpenID != "" {
			ctx.StatusCode(iris.StatusAccepted)
			info := title + ":\tCreate接口不支持微信信息"
			_, _ = ctx.Text(info)
			log.Println(info)
		} else {
			if err := u.Create(); err != nil {
				ctx.StatusCode(iris.StatusAccepted)
				info := err.Error()
				_, _ = ctx.Text(info)
				log.Println(info)
			} else {
				ctx.StatusCode(iris.StatusOK)
				_, _ = ctx.JSON(u)
			}
		}
	}
}

func UserBind(ctx iris.Context) {
	u := new(service.UserJSON)
	if err := ctx.ReadJSON(u); err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		info := err.Error()
		_, _ = ctx.Text(info)
		log.Println(info)
	} else {
		if err := u.Bind(); err != nil {
			ctx.StatusCode(iris.StatusAccepted)
			info := err.Error()
			_, _ = ctx.Text(info)
			log.Println(info)
		} else {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(u)
		}
	}

}

func UserFindByID(ctx iris.Context) {
	if id, err := ctx.Params().GetUint("id"); err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		info := err.Error()
		_, _ = ctx.Text(info)
		log.Println(info)
	} else {
		if u, err := service.UserFindByID(id); err != nil {
			ctx.StatusCode(iris.StatusAccepted)
			info := err.Error()
			_, _ = ctx.Text(info)
			log.Println(info)
		} else {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(u)
		}
	}
}

func UserFindByIDCard(ctx iris.Context) {
	idCard := ctx.Params().GetString("id_card")
	if u, err := service.UserFindByIDCard(idCard); err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		info := err.Error()
		_, _ = ctx.Text(info)
		log.Println(info)
	} else {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(u)
	}
}

func UserFindByOpenID(ctx iris.Context) {
	openid := ctx.Params().GetString("openid")
	if u, err := service.UserFindByOpenID(openid); err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		info := err.Error()
		_, _ = ctx.Text(info)
		log.Println(info)
	} else {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(u)
	}
}

func UsersFindByLevel(ctx iris.Context) {
	if level, err := ctx.Params().GetInt("level"); err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		info := err.Error()
		_, _ = ctx.Text(info)
		log.Println(info)
	} else {
		if usersJson, err := service.UsersFindByLevel(level); err != nil {
			ctx.StatusCode(iris.StatusAccepted)
			info := err.Error()
			_, _ = ctx.Text(info)
			log.Println(info)
		} else {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(usersJson)
		}
	}
}

func UserDeleteByID(ctx iris.Context) {
	if id, err := ctx.Params().GetUint("id"); err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		info := err.Error()
		_, _ = ctx.Text(info)
		log.Println(info)
	} else {
		if userJson, err := service.UserDeleteByID(id); err != nil {
			ctx.StatusCode(iris.StatusAccepted)
			info := err.Error()
			_, _ = ctx.Text(info)
			log.Println(info)
		} else {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(userJson)
		}
	}

}

func UserDeleteByIDCard(ctx iris.Context) {
	idCard := ctx.Params().GetString("id_card")
	if userJson, err := service.UserDeleteByIDCard(idCard); err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		info := err.Error()
		_, _ = ctx.Text(info)
		log.Println(info)
	} else {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(userJson)
	}
}

func UserUpdates(ctx iris.Context) {
	u := new(service.UserJSON)
	if err := ctx.ReadJSON(u); err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		info := err.Error()
		_, _ = ctx.Text(info)
		log.Println(info)
	} else {
		if err := u.Updates(); err != nil {
			ctx.StatusCode(iris.StatusAccepted)
			info := err.Error()
			_, _ = ctx.Text(info)
			log.Println(info)
		} else {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(u)
		}
	}

}
