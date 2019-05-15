package controller

import (
	"github.com/kataras/iris"
	"github.com/pantazheng/bci/service"
	"log"
)

func GainCreate(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 15:40
	*/
	g := new(service.GainJSON)
	info := ""
	if err := ctx.ReadJSON(g); err == nil {
		if err := g.Create(); err == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(g)
		} else {
			info = err.Error()
		}
	} else {
		info = err.Error()
	}
	if info != "" {
		ctx.StatusCode(iris.StatusAccepted)
		_, _ = ctx.Text(info)
		log.Println(info)
	}
}

func GainFindByID(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 15:40
	*/
	info := ""
	if id, err := ctx.Params().GetUint("id"); err == nil {
		if g, err := service.GainFindByID(id); err == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(g)
		} else {
			info = err.Error()
		}
	} else {
		info = err.Error()
	}
	if info != "" {
		ctx.StatusCode(iris.StatusAccepted)
		_, _ = ctx.Text(info)
		log.Println(info)
	}
}

func GainsFindByOwnerID(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 15:40
	*/
	info := ""
	if id, err := ctx.Params().GetUint("id"); err == nil {
		if gainsJson, err := service.GainsFindByOwnerID(id); err == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(gainsJson)
		} else {
			info = err.Error()
		}
	} else {
		info = err.Error()
	}
	if info != "" {
		ctx.StatusCode(iris.StatusAccepted)
		_, _ = ctx.Text(info)
		log.Println(info)
	}
}

func GainsFindByMissionID(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 15:40
	*/
	info := ""
	if id, err := ctx.Params().GetUint("id"); err == nil {
		if gainsJson, err := service.GainsFindByMissionID(id); err == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(gainsJson)
		} else {
			info = err.Error()
		}
	} else {
		info = err.Error()
	}
	if info != "" {
		ctx.StatusCode(iris.StatusAccepted)
		_, _ = ctx.Text(info)
		log.Println(info)
	}
}

func GainUpdate(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 15:40
	*/
	g := new(service.GainJSON)
	info := ""
	if err := ctx.ReadJSON(g); err == nil {
		if err := g.Updates(); err == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(g)
		} else {
			info = err.Error()
		}
	} else {
		info = err.Error()
	}
	if info != "" {
		ctx.StatusCode(iris.StatusAccepted)
		_, _ = ctx.Text(info)
		log.Println(info)
	}
}

func GainDeleteByID(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 15:43
	*/
	info := ""
	if id, err := ctx.Params().GetUint("id"); err == nil {
		if gainJson, err := service.GainDeleteByID(id); err == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(gainJson)
		} else {
			info = err.Error()
		}
	} else {
		info = err.Error()
	}
	if info != "" {
		ctx.StatusCode(iris.StatusAccepted)
		_, _ = ctx.Text(info)
		log.Println(info)
	}
}
