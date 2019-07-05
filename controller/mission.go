package controller

import (
	"errors"
	"github.com/kataras/iris"
	"github.com/pantazheng/bci/service"
	"io"
	"os"
	"strconv"
)

//MissionInsert
func MissionInsert(ctx iris.Context) {
	m := &service.MissionJSON{}
	if err := ctx.ReadJSON(m); err != nil {
		ErrorProcess(err, ctx)
		return
	}
	if err := m.Insert(); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(m)
	} else {
		ErrorProcess(err, ctx)
	}
	return
}

//MissionFindByID
func MissionFindByID(ctx iris.Context) {
	m := &service.MissionJSON{}
	if id, err := ctx.Params().GetUint("id"); err != nil {
		ErrorProcess(err, ctx)
		return
	} else {
		m.ID = id
	}
	if err := m.First(); err != nil {
		ErrorProcess(err, ctx)
	} else {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(m)
	}
	return
}

func missionsFind(field string, ctx iris.Context) {
	m := &service.MissionJSON{}
	if field != "all" {
		if id, err := ctx.Params().GetUint("id"); err != nil {
			ErrorProcess(err, ctx)
			return
		} else {
			if field == "leader_id" {
				m.LeaderID = id
			} else if field == "owner_id" {
				m.OwnerID = id
			} else if field == "module_id" {
				m.ModuleID = id
			} else if field == "all" {
			} else {
				err = errors.New("no this field")
				ErrorProcess(err, ctx)
				return
			}
		}
	}
	if missionsJSON, err := m.Find(field); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(missionsJSON)
	} else {
		ErrorProcess(err, ctx)
	}
	return
}

func MissionsFindByLeaderID(ctx iris.Context) {
	missionsFind("leader_id", ctx)
}

func MissionsFindByOwnerID(ctx iris.Context) {
	missionsFind("owner_id", ctx)
}

func MissionsFindByModuleID(ctx iris.Context) {
	missionsFind("module_id", ctx)
}

func MissionsFindAll(ctx iris.Context) {
	missionsFind("all", ctx)
}

func MissionUpdate(ctx iris.Context) {
	m := &service.MissionJSON{}
	if err := ctx.ReadJSON(m); err != nil {
		ErrorProcess(err, ctx)
		return
	}
	if err := m.Updates(); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(m)
	} else {
		ErrorProcess(err, ctx)
	}
	return
}

func MissionDeleteByID(ctx iris.Context) {
	m := &service.MissionJSON{}
	if id, err := ctx.Params().GetUint("id"); err != nil {
		ErrorProcess(err, ctx)
		return
	} else {
		m.ID = id
	}
	if err := m.Delete(); err != nil {
		ErrorProcess(err, ctx)
	} else {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(m)
	}
	return
}
