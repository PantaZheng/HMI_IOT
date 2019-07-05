package controller

import (
	"github.com/kataras/iris"
	"github.com/pantazheng/bci/models"
)

//MissionInsert
func MissionInsert(ctx iris.Context) {
	m := &models.Mission{}
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
	m := &models.Mission{}
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
	m := &models.Mission{}
	id, err := ctx.Params().GetUint("id")
	if err != nil {
		ErrorProcess(err, ctx)
		return
	}
	if missions, err := m.FindBrief(field, id); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(missions)
	} else {
		ErrorProcess(err, ctx)
	}
	return
}

func MissionsFindByOwnerID(ctx iris.Context) {
	missionsFind("owner", ctx)
}

func MissionsFindByModuleID(ctx iris.Context) {
	missionsFind("module", ctx)
}

func MissionsFindByLeaderID(ctx iris.Context) {
	missionsFind("leader", ctx)
}

func MissionsFindByProjectID(ctx iris.Context) {
	missionsFind("project", ctx)
}

func MissionsFindByManagerID(ctx iris.Context) {
	missionsFind("manager", ctx)
}

func MissionsFindAll(ctx iris.Context) {
	m := &models.Mission{}
	if missions, err := m.FindBrief("all", 0); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(missions)
	} else {
		ErrorProcess(err, ctx)
	}
	return
}

func MissionUpdate(ctx iris.Context) {
	m := &models.Mission{}
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
	m := &models.Mission{}
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
