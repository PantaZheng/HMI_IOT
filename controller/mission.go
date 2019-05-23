package controller

import (
	"github.com/kataras/iris"
	"github.com/pantazheng/bci/service"
	"io"
	"log"
	"os"
	"strconv"
)

func MissionCreate(ctx iris.Context) {
	var err error
	m := new(service.MissionJSON)
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
	var err error
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

func MissionsFindAll(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/21 12:34
	*/
	if missionsJSON, err := service.MissionsFindAll(); err == nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(missionsJSON)
	} else {
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
	var err error
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
	var err error
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
	var err error
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

func MissionUpFileByID(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/23 10:00
	*/
	var err error
	if id, err1 := ctx.Params().GetUint("id"); err1 == nil {
		if m, err2 := service.MissionFindByID(id); err2 == nil {
			fileName := m.File
			file, _, err3 := ctx.FormFile(fileName)
			if err3 == nil {
				out, err4 := os.OpenFile("./files/mission"+strconv.Itoa(int(m.ID))+"_"+fileName,
					os.O_WRONLY|os.O_CREATE, 0666)
				if err4 == nil {
					_, err = io.Copy(out, file)
				} else {
					err = err2
				}
				defer func() {
					err = out.Close()
				}()
			} else {
				err = err3
			}
			defer func() {
				err = file.Close()
			}()
			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.Text("文件传输完成")
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

func MissionDownFileByID(ctx iris.Context) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/20 11:17
	*/
	var err error
	if id, err1 := ctx.Params().GetUint("id"); err1 == nil {
		if m, err2 := service.MissionFindByID(id); err2 == nil {
			fileName := m.File
			ctx.StatusCode(iris.StatusOK)
			_ = ctx.SendFile("./files/mission"+strconv.Itoa(int(m.ID))+"_"+fileName, fileName)
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
