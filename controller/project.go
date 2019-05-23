package controller

import (
	"github.com/kataras/iris"
	"github.com/pantazheng/bci/service"
	"log"
)

func ProjectCreate(ctx iris.Context) {
	p := new(service.ProjectJSON)
	err := *new(error)
	if err = ctx.ReadJSON(p); err == nil {
		if err = p.Create(); err == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(p)
		}
	}
	if err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		info := err.Error()
		_, _ = ctx.Text(info)
		log.Println(info)
	}
}

func ProjectFindByID(ctx iris.Context) {
	err := *new(error)
	if id, err1 := ctx.Params().GetUint("id"); err1 == nil {
		if p, err2 := service.ProjectFindByID(id); err2 == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(p)
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

func ProjectFramPaceByID(ctx iris.Context) {
	err := *new(error)
	if id, err1 := ctx.Params().GetUint("id"); err1 == nil {
		if p, err2 := service.ProjectFramePaceByID(id); err2 == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(p)
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

func ProjectsFindAll(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 23:52
	*/
	if projectsJSON, err := service.ProjectsFindAll(); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(projectsJSON)
	} else {
		ctx.StatusCode(iris.StatusAccepted)
		info := err.Error()
		_, _ = ctx.Text(info)
		log.Println(info)
	}
}

func ProjectsFindByCreatorID(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 22:21
	*/
	err := *new(error)
	if id, err1 := ctx.Params().GetUint("id"); err1 == nil {
		if projects, err2 := service.ProjectsFindByCreatorID(id); err2 == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(projects)
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

func ProjectsFindByLeaderID(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 23:55
	*/
	err := *new(error)
	if id, err1 := ctx.Params().GetUint("id"); err1 == nil {
		if projectsJSON, err2 := service.ProjectsFindByLeaderID(id); err2 == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(projectsJSON)
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

func ProjectsFindByParticipantID(ctx iris.Context) {
	err := *new(error)
	if id, err1 := ctx.Params().GetUint("id"); err1 == nil {
		if projectsJSON, err2 := service.ProjectsFindByParticipantID(id); err2 == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(projectsJSON)
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

func ProjectUpdate(ctx iris.Context) {
	p := new(service.ProjectJSON)
	err := *new(error)
	if err = ctx.ReadJSON(p); err == nil {
		if err = p.Updates(); err == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(p)
		}
	}
	if err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		info := err.Error()
		_, _ = ctx.Text(info)
		log.Println(info)
	}
}

func ProjectDeleteByID(ctx iris.Context) {
	err := *new(error)
	if id, err1 := ctx.Params().GetUint("id"); err1 == nil {
		if p, err2 := service.ProjectDeleteByID(id); err2 == nil {
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(p)
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
