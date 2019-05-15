package controller

import (
	"github.com/kataras/iris"
	"github.com/pantazheng/bci/service"
	"log"
)

func ModuleCreate(ctx iris.Context) {
	m := new(service.ModuleJSON)
	err := *new(error)
	if err = ctx.ReadJSON(m); err != nil {
		if err = m.Create(); err == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(m)
		}
	}
	if err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		info := err.Error()
		_, _ = ctx.Text(info)
		log.Println(info)
	}
}

func ModuleFindByID(ctx iris.Context) {
	err := *new(error)
	if id, err1 := ctx.Params().GetUint("id"); err1 == nil {
		if m, err2 := service.ModuleFindByID(id); err2 == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(m)
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

func ModulesFindByCreatorID(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 22:21
	*/
	err := *new(error)
	if id, err1 := ctx.Params().GetUint("id"); err1 == nil {
		if modules, err2 := service.ModulesFindByCreatorID(id); err2 == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(modules)
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

func ModulesFindByLeaderID(ctx iris.Context) {
	err := *new(error)
	if id, err1 := ctx.Params().GetUint("id"); err1 == nil {
		if modules, err2 := service.ModulesFindByLeaderID(id); err2 == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(modules)
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

func ModulesFindByParticipantID(ctx iris.Context) {
	err := *new(error)
	if id, err1 := ctx.Params().GetUint("id"); err1 == nil {
		if modules, err2 := service.ModulesFindByParticipantID(id); err2 == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(modules)
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

func ModulesFindByProjectID(ctx iris.Context) {
	err := *new(error)
	if id, err1 := ctx.Params().GetUint("id"); err1 == nil {
		if modules, err2 := service.ModulesFindByProjectID(id); err2 == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(modules)
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

func ModuleUpdate(ctx iris.Context) {
	m := new(service.ModuleJSON)
	err := *new(error)
	if err = ctx.ReadJSON(m); err == nil {
		if err = m.Updates(); err == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(m)
		}
	}
	if err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		info := err.Error()
		_, _ = ctx.Text(info)
		log.Println(info)
	}
}

func ModuleDeleteByID(ctx iris.Context) {
	err := *new(error)
	if id, err1 := ctx.Params().GetUint("id"); err1 == nil {
		if missionJSON, err2 := service.ModuleDeleteByID(id); err2 == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(missionJSON)
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
