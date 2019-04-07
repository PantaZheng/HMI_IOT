package controller

import "github.com/kataras/iris"

func Enroll(ctx iris.Context) {
	_ = ctx.View("notice/index.html")
}
