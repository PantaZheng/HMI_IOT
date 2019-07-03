package controller

import (
	"github.com/kataras/iris"
	"log"
)

/**
*@Author: PantaZheng
*@CreatedAt: 2019/6/30 22:23
*@Title: common.go
*@Package: controller
*@Description: (用一句话描述该文件该做什么)
@Software: GoLand
*/

func ErrorProcess(err error, ctx iris.Context) {
	ctx.StatusCode(iris.StatusAccepted)
	info := err.Error()
	_, _ = ctx.Text(info)
	log.Println(info)
}
