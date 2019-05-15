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
	err := *new(error)
	if err = ctx.ReadJSON(g); err == nil {
		if err = g.Create(); err == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(g)
		}
	}
	if err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		info := err.Error()
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
	err := *new(error)
	if id, err1 := ctx.Params().GetUint("id"); err1 == nil {
		if g, err2 := service.GainFindByID(id); err2 == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(g)
		} else {
			err=err2
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

func GainsFindByOwnerID(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 15:40
	*/
	err := *new(error)
	if id, err1 := ctx.Params().GetUint("id"); err1 == nil {
		if gainsJson, err2 := service.GainsFindByOwnerID(id); err2 == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(gainsJson)
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

func GainsFindByMissionID(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 15:40
	*/
	err := *new(error)
	if id, err1 := ctx.Params().GetUint("id"); err1 == nil {
		if gainsJson, err2 := service.GainsFindByMissionID(id); err2 == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(gainsJson)
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

func GainUpdate(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 15:40
	*/
	g := new(service.GainJSON)
	err := *new(error)
	if err = ctx.ReadJSON(g); err == nil {
		if err = g.Updates(); err == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(g)
		}
	}
	if err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		info := err.Error()
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
	err := *new(error)
	if id, err1 := ctx.Params().GetUint("id"); err1 == nil {
		if gainJson, err2 := service.GainDeleteByID(id); err2 == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(gainJson)
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
