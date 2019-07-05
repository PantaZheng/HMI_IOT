package controller

import (
	"github.com/kataras/iris"
	"github.com/pantazheng/bci/models"
)

//ModuleInsert
func ModuleInsert(ctx iris.Context) {
	m := &models.Module{}
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
	m := &models.Module{}
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
	m := &models.Module{}
	id, err := ctx.Params().GetUint("id")
	if err != nil {
		ErrorProcess(err, ctx)
		return
	}
	if modules, err := m.FindBrief(field, id); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(modules)
	} else {
		ErrorProcess(err, ctx)
	}
	return
}

func ModulesFindByLeaderID(ctx iris.Context) {
	modulesFind("leader", ctx)
}

func ModulesFindByProjectID(ctx iris.Context) {
	modulesFind("project", ctx)
}

func ModulesFindByManagerID(ctx iris.Context) {
	modulesFind("manager", ctx)
}

func ModulesFindByMemberID(ctx iris.Context) {
	modulesFind("member", ctx)
}

func ModulesFindAll(ctx iris.Context) {
	m := &models.Module{}
	if modules, err := m.FindBrief("all", 0); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(modules)
	} else {
		ErrorProcess(err, ctx)
	}
	return
}

func ModuleUpdate(ctx iris.Context) {
	m := &models.Module{}
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

func ModuleDeleteByID(ctx iris.Context) {
	m := &models.Module{}
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
