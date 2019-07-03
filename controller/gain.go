package controller

import (
	"errors"
	"github.com/kataras/iris"
	"github.com/pantazheng/bci/service"
	"io"
	"os"
	"strconv"
)

//GainInsert
func GainInsert(ctx iris.Context) {
	g := &service.GainJSON{}
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
	g := &service.GainJSON{}
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
	g := &service.GainJSON{}
	if field != "all" {
		if id, err := ctx.Params().GetUint("id"); err != nil {
			ErrorProcess(err, ctx)
			return
		} else {
			if field == "leader_id" {
				g.LeaderID = id
			} else if field == "owner_id" {
				g.OwnerID = id
			} else if field == "mission_id" {
				g.MissionID = id
			} else if field == "all" {
			} else {
				err = errors.New("no this field")
				ErrorProcess(err, ctx)
				return
			}
		}
	}
	if gainsJSON, err := g.Find(field); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(gainsJSON)
	} else {
		ErrorProcess(err, ctx)
	}
	return
}

func GainsFindByLeaderID(ctx iris.Context) {
	gainsFind("leader_id", ctx)
}

func GainsFindByOwnerID(ctx iris.Context) {
	gainsFind("owner_id", ctx)
}

func GainsFindByMissionID(ctx iris.Context) {
	gainsFind("mission_id", ctx)
}

func GainsFindAll(ctx iris.Context) {
	gainsFind("all", ctx)
}

func GainUpdates(ctx iris.Context) {
	g := &service.GainJSON{}
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
	g := &service.GainJSON{}
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
	g := &service.GainJSON{}
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
	g := &service.GainJSON{}
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
