package controller

//import (
//	"github.com/kataras/iris"
//	"github.com/pantazheng/bci/models"
//	"github.com/pantazheng/bci/service"
//	"log"
//)
//
//func ProjectCreate(ctx iris.Context) {
//	project := new(models.ProjectJson)
//	if err := ctx.ReadJSON(project); err != nil {
//		log.Println(err.Error())
//	}
//	if projectJson, err := service.ProjectCreate(project); err != nil {
//		ctx.StatusCode(iris.StatusAccepted)
//		_, _ = ctx.Text(err.Error())
//	} else {
//		ctx.StatusCode(iris.StatusOK)
//		_, _ = ctx.JSON(projectJson)
//	}
//}
//
//func ProjectFindByID(ctx iris.Context) {
//	id, _ := ctx.Params().GetUint("id")
//	if projectJson, err := service.ProjectFindByID(id); err != nil {
//		ctx.StatusCode(iris.StatusAccepted)
//		_, _ = ctx.Text(err.Error())
//	} else {
//		ctx.StatusCode(iris.StatusOK)
//		_, _ = ctx.JSON(projectJson)
//	}
//}
//
//func ProjectsFindByLeaderID(ctx iris.Context) {
//	id, _ := ctx.Params().GetUint("id")
//	if projectsJson, err := service.ProjectsFindByLeaderID(id); err != nil {
//		ctx.StatusCode(iris.StatusAccepted)
//		_, _ = ctx.Text(err.Error())
//	} else {
//		ctx.StatusCode(iris.StatusOK)
//		_, _ = ctx.JSON(projectsJson)
//	}
//}
//
//func ProjectsFindByParticipantID(ctx iris.Context) {
//	id, _ := ctx.Params().GetUint("id")
//	if projectsJson, err := service.ProjectsFindByParticipantID(id); err != nil {
//		ctx.StatusCode(iris.StatusAccepted)
//		_, _ = ctx.Text(err.Error())
//	} else {
//		ctx.StatusCode(iris.StatusOK)
//		_, _ = ctx.JSON(projectsJson)
//	}
//}
//
//func ProjectUpdate(ctx iris.Context) {
//	project := new(models.ProjectJson)
//	if err := ctx.ReadJSON(project); err != nil {
//		log.Println(err.Error())
//	}
//	if projectJson, err := service.ProjectUpdate(project); err != nil {
//		ctx.StatusCode(iris.StatusAccepted)
//		_, _ = ctx.Text(err.Error())
//	} else {
//		ctx.StatusCode(iris.StatusOK)
//		_, _ = ctx.JSON(projectJson)
//	}
//}
//
//func ProjectDeleteByID(ctx iris.Context) {
//	id, _ := ctx.Params().GetUint("id")
//	if projectJson, err := service.ProjectDeleteByID(id); err != nil {
//		ctx.StatusCode(iris.StatusAccepted)
//		_, _ = ctx.Text(err.Error())
//	} else {
//		ctx.StatusCode(iris.StatusOK)
//		_, _ = ctx.JSON(projectJson)
//	}
//}
