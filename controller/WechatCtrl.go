package controller

import (
	"../service"
	"github.com/kataras/iris"
)


// wxCallbackHandler 是处理回调请求的 http handler.
//  1. 不同的 web 框架有不同的实现
//  2. 一般一个 handler 处理一个公众号的回调请求(当然也可以处理多个, 这里我只处理一个)
func WeChat(ctx iris.Context) {
	service.WechatServer(ctx)
}

func Menu(){
	service.CreateTag()
	service.GetTag()
	service.DefaultMenu()
	service.TeacherMenu()
	service.StudentMenu()
	service.TestMenu()
}




