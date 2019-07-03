package controller

import (
	"errors"
	"github.com/kataras/iris"
	"github.com/pantazheng/bci/service"
)

func ProjectInsert(ctx iris.Context) {
	p := service.ProjectJSON{}
	if err := ctx.ReadJSON(p); err == nil {
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
	p := service.ProjectJSON{}
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
	p := service.ProjectJSON{}
	if id, err := ctx.Params().GetUint("id"); err != nil {
		ErrorProcess(err, ctx)
		return
	} else {
		if field == "leader_id" || field == "participant_id" || field == "member_id" {
			p.LeaderID = id
		} else if field == "creator_id" {
			p.CreatorID = id
		} else if field == "all" {
		} else {
			err = errors.New("no this field")
			ErrorProcess(err, ctx)
			return
		}
	}
	if projectsJSON, err := p.Find(field); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(projectsJSON)
	} else {
		ErrorProcess(err, ctx)
	}
	return
}

func ProjectsFindByLeaderID(ctx iris.Context) {
	projectsFind("leader_id", ctx)
}

func ProjectsFindByCreatorID(ctx iris.Context) {
	projectsFind("creator_id", ctx)
}

func ProjectsFindByParticipantID(ctx iris.Context) {
	projectsFind("participant_id", ctx)
}

func ProjectsFindByMemberID(ctx iris.Context) {
	projectsFind("member_id", ctx)
}

func ProjectsFindAll(ctx iris.Context) {
	projectsFind("all", ctx)
}

func ProjectUpdate(ctx iris.Context) {
	p := service.ProjectJSON{}
	if err := ctx.ReadJSON(p); err == nil {
		ErrorProcess(err, ctx)
		return
	}
	if err := p.Update(); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(p)
	} else {
		ErrorProcess(err, ctx)
	}
	return
}

func ProjectDeleteByID(ctx iris.Context) {
	p := service.ProjectJSON{}
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
