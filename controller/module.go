package controller

import (
	"errors"
	"github.com/kataras/iris"
	"github.com/pantazheng/bci/service"
)

//ModuleInsert
func ModuleInsert(ctx iris.Context) {
	m := service.ModuleJSON{}
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

//ModuleFindByID
func ModuleFindByID(ctx iris.Context) {
	m := service.ModuleJSON{}
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

func modulesFind(field string, ctx iris.Context) {
	m := service.ModuleJSON{}
	if id, err := ctx.Params().GetUint("id"); err != nil {
		ErrorProcess(err, ctx)
		return
	} else {
		if field == "leader_id" {
			m.LeaderID = id
		} else if field == "creator_id" {
			m.CreatorID = id
		} else if field == "project_id" {
			m.ProjectID = id
		} else if field == "all" {
		} else {
			err = errors.New("no this field")
			ErrorProcess(err, ctx)
			return
		}
	}
	if modulesJSON, err := m.Find(field); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(modulesJSON)
	} else {
		ErrorProcess(err, ctx)
	}
	return
}

func ModulesFindByLeaderID(ctx iris.Context) {
	modulesFind("leader_id", ctx)
}

func ModulesFindByCreatorID(ctx iris.Context) {
	modulesFind("creator_id", ctx)
}

func ModulesFindByProjectID(ctx iris.Context) {
	modulesFind("project_id", ctx)
}

func ModulesFindAll(ctx iris.Context) {
	modulesFind("all", ctx)
}

func ModuleUpdate(ctx iris.Context) {
	m := service.ModuleJSON{}
	if err := ctx.ReadJSON(m); err != nil {
		ErrorProcess(err, ctx)
		return
	}
	if err := m.Update(); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(m)
	} else {
		ErrorProcess(err, ctx)
	}
	return

}

func ModuleDeleteByID(ctx iris.Context) {
	m := service.ModuleJSON{}
	if id, err := ctx.Params().GetUint("id"); err != nil {
		ErrorProcess(err, ctx)
		return
	} else {
		m.ID = id
	}
	if err := m.Delete(); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(m)
	} else {
		ErrorProcess(err, ctx)
	}
	return
}
