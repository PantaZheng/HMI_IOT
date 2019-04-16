package controller

import (
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/menu"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
	"github.com/kataras/iris"
	"github.com/pantazheng/HMI-IOT/service"
)


// wxCallbackHandler 是处理回调请求的 http handler.
//  1. 不同的 web 框架有不同的实现
//  2. 一般一个 handler 处理一个公众号的回调请求(当然也可以处理多个, 这里我只处理一个)
func WeChat(ctx iris.Context) {
	mux := core.NewServeMux()
	mux.DefaultMsgHandleFunc(service.DefaultEventHandler)
	mux.DefaultEventHandleFunc(service.DefaultEventHandler)
	mux.MsgHandleFunc(request.MsgTypeText, service.TextMsgHandler)
	mux.EventHandleFunc(menu.EventTypeClick, service.MenuClickEventHandler)
	mux.EventHandleFunc(request.EventTypeSubscribe,service.SubscribeEventHandler)
	msgHandler := mux

	msgServer := core.NewServer(service.WeChatOriId, service.WeChatAppId,service.WeChatToken , service.WeChatEncodedAESKey,msgHandler, nil)
	msgServer.ServeHTTP(ctx.ResponseWriter(), ctx.Request(), nil)
}

func Menu(){
	btnRelationShip:=menu.Button{}
	btnRelationShip.SetAsViewButton("架构", service.ServerAddress+"/framework/frame.html")
	btnProjectMission:=menu.Button{}
	btnProjectMission.SetAsViewButton("项目/任务", service.ServerAddress+"/project/index.html")
	btnEnroll:=menu.Button{}
	btnEnroll.SetAsViewButton("个人信息","https://open.weixin.qq.com/connect/oauth2/authorize?appid="+service.WeChatAppId+"&redirect_uri="+service.ServerAddress+"/createUser&response_type=code&scope=snsapi_base&state=12#wechat_redirect")
	defaultButtons:= []menu.Button{btnRelationShip,btnProjectMission,btnEnroll}
	service.DefaultMenu(&menu.Menu{Buttons:defaultButtons})
}




