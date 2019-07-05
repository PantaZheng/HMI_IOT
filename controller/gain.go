package controller

import (
	"github.com/kataras/iris"
	"github.com/pantazheng/bci/models"
	"io"
	"os"
	"strconv"
)

//GainInsert
func GainInsert(ctx iris.Context) {
	g := &models.Gain{}
	if err := ctx.ReadJSON(g); err != nil {
		ErrorProcess(err, ctx)
		return
	}
	if err := g.Insert(); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(g)
	} else {
		ErrorProcess(err, ctx)
	}
	return
}

func GainFindByID(ctx iris.Context) {
	g := &models.Gain{}
	if id, err := ctx.Params().GetUint("id"); err != nil {
		ErrorProcess(err, ctx)
		return
	} else {
		g.ID = id
	}
	if err := g.First(); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(g)
	} else {
		ErrorProcess(err, ctx)
	}
	return
}

func gainsFind(field string, ctx iris.Context) {
	g := &models.Gain{}
	id, err := ctx.Params().GetUint("id")
	if err != nil {
		ErrorProcess(err, ctx)
		return
	}
	if gains, err := g.FindBrief(field, id); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(gains)
	} else {
		ErrorProcess(err, ctx)
	}
	return
}

func GainsFindByMissionID(ctx iris.Context) {
	gainsFind("mission", ctx)
}

func GainsFindByOwnerID(ctx iris.Context) {
	gainsFind("owner", ctx)
}

func GainsFindByModuleID(ctx iris.Context) {
	gainsFind("module", ctx)
}

func GainsFindByLeaderID(ctx iris.Context) {
	gainsFind("leader", ctx)
}

func GainsFindByProjectID(ctx iris.Context) {
	gainsFind("project", ctx)
}

func GainsFindByManagerID(ctx iris.Context) {
	gainsFind("manager", ctx)
}

func GainsFindAll(ctx iris.Context) {
	g := &models.Gain{}
	if gains, err := g.FindBrief("all", 0); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(gains)
	} else {
		ErrorProcess(err, ctx)
	}
	return
}

func GainUpdates(ctx iris.Context) {
	g := &models.Gain{}
	if err := ctx.ReadJSON(g); err != nil {
		ErrorProcess(err, ctx)
		return
	}
	if err := g.Updates(); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(g)
	} else {
		ErrorProcess(err, ctx)
	}
	return
}

func GainDeleteByID(ctx iris.Context) {
	g := &models.Gain{}
	if id, err := ctx.Params().GetUint("id"); err != nil {
		ErrorProcess(err, ctx)
		return
	} else {
		g.ID = id
	}
	if err := g.Delete(); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(g)
	} else {
		ErrorProcess(err, ctx)
	}
	return
}

func GainUpFileByID(ctx iris.Context) {
	g := &models.Gain{}
	if id, err := ctx.Params().GetUint("id"); err != nil {
		ErrorProcess(err, ctx)
		return
	} else {
		g.ID = id
	}
	if err := g.First(); err != nil {
		ErrorProcess(err, ctx)
		return
	}
	fileName := g.File
	if file, _, err := ctx.FormFile(fileName); err != nil {
		ErrorProcess(err, ctx)
		return
	} else {
		out, err := os.OpenFile("./files/gain"+strconv.Itoa(int(g.ID))+"_"+fileName,
			os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			ErrorProcess(err, ctx)
			return
		}
		if _, err := io.Copy(out, file); err != nil {
			ErrorProcess(err, ctx)
			return
		}
		defer out.Close()
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.Text("文件传输完成")
	}
	return
}

func GainDownFileByID(ctx iris.Context) {
	g := &models.Gain{}
	if id, err := ctx.Params().GetUint("id"); err != nil {
		ErrorProcess(err, ctx)
		return
	} else {
		g.ID = id
	}
	if err := g.First(); err != nil {
		ErrorProcess(err, ctx)
		return
	}
	fileName := g.File
	ctx.StatusCode(iris.StatusOK)
	_ = ctx.SendFile("./files/gain"+strconv.Itoa(int(g.ID))+"_"+fileName, fileName)
	return
}
