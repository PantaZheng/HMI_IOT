package controller

import (
	"github.com/kataras/iris"
	"github.com/pantazheng/bci/models"
)

func ProjectInsert(ctx iris.Context) {
	p := &models.Project{}
	if err := ctx.ReadJSON(p); err != nil {
		ErrorProcess(err, ctx)
		return
	}
	if err := p.Insert(); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(p)
	} else {
		ErrorProcess(err, ctx)
	}
	return
}

func ProjectFindByID(ctx iris.Context) {
	p := &models.Project{}
	if id, err := ctx.Params().GetUint("id"); err != nil {
		ErrorProcess(err, ctx)
		return
	} else {
		p.ID = id
	}
	if err := p.First(); err != nil {
		ErrorProcess(err, ctx)
	} else {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(p)
	}
	return
}

func projectsFind(field string, ctx iris.Context) {
	p := &models.Project{}
	id, err := ctx.Params().GetUint("id")
	if err != nil {
		ErrorProcess(err, ctx)
		return
	}
	if projects, err := p.FindBrief(field, id); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(projects)
	} else {
		ErrorProcess(err, ctx)
	}
	return
}

func ProjectsFindByManagerID(ctx iris.Context) {
	projectsFind("manager", ctx)
}

func ProjectsFindByMemberID(ctx iris.Context) {
	projectsFind("member", ctx)
}

func ProjectsFindAll(ctx iris.Context) {
	p := &models.Project{}
	if projects, err := p.FindBrief("all", 0); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(projects)
	} else {
		ErrorProcess(err, ctx)
	}
}

func ProjectUpdate(ctx iris.Context) {
	p := &models.Project{}
	if err := ctx.ReadJSON(p); err == nil {
		ErrorProcess(err, ctx)
		return
	}
	if err := p.Updates(); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(p)
	} else {
		ErrorProcess(err, ctx)
	}
	return
}

func ProjectDeleteByID(ctx iris.Context) {
	p := &models.Project{}
	if id, err := ctx.Params().GetUint("id"); err == nil {
		ErrorProcess(err, ctx)
		return
	} else {
		p.ID = id
	}
	if err := p.Delete(); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(p)
	} else {
		ErrorProcess(err, ctx)
	}
	return
}
