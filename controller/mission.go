package controller

import (
	"github.com/kataras/iris"
	"github.com/pantazheng/bci/service"
	"log"
)

func MissionCreate(ctx iris.Context) {
	m := new(service.MissionJSON)
	err := *new(error)
	if err = ctx.ReadJSON(m); err == nil {
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

func MissionFindByID(ctx iris.Context) {
	err := *new(error)
	if id, err1 := ctx.Params().GetUint("id"); err1 == nil {
		if g, err2 := service.MissionFindByID(id); err2 == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(g)
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

func MissionsFindByCreatorID(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 15:40
	*/
	err := *new(error)
	if id, err1 := ctx.Params().GetUint("id"); err1 == nil {
		if gainsJson, err2 := service.MissionsFindByCreatorID(id); err2 == nil {
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

func MissionsFindByParticipantID(ctx iris.Context) {
	err := *new(error)
	if id, err1 := ctx.Params().GetUint("id"); err1 == nil {
		if gainsJson, err2 := service.MissionsFindByParticipantID(id); err2 == nil {
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

func MissionsFindByModuleID(ctx iris.Context) {
	err := *new(error)
	if id, err1 := ctx.Params().GetUint("id"); err1 == nil {
		if gainsJson, err2 := service.MissionsFindByModuleID(id); err2 == nil {
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

func MissionUpdate(ctx iris.Context) {
	m := new(service.MissionJSON)
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

func MissionDeleteByID(ctx iris.Context) {
	err := *new(error)
	if id, err1 := ctx.Params().GetUint("id"); err1 == nil {
		if missionJSON, err2 := service.MissionDeleteByID(id); err2 == nil {
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
