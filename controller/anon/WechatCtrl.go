package controller

import (
	"../../service"
	"../../config"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/menu"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
	"github.com/kataras/iris"
)


// wxCallbackHandler 是处理回调请求的 http handler.
//  1. 不同的 web 框架有不同的实现
//  2. 一般一个 handler 处理一个公众号的回调请求(当然也可以处理多个, 这里我只处理一个)
func Wechat(ctx iris.Context) {


	mux := core.NewServeMux()
	mux.DefaultMsgHandleFunc(service.DefaultEventHandler)
	mux.DefaultEventHandleFunc(service.DefaultEventHandler)
	mux.MsgHandleFunc(request.MsgTypeText, service.TextMsgHandler)
	mux.EventHandleFunc(menu.EventTypeClick, service.MenuClickEventHandler)

	msgHandler := mux
	wechatOriId := config.Conf.Get("wechat.OriId").(string)
	wechatAppId := config.Conf.Get("wechat.AppId").(string)
	wechatToken := config.Conf.Get("wechat.Token").(string)
	wechatEncodedAESKey := config.Conf.Get("wechat.EncodedAESKey").(string)
	msgServer := core.NewServer(wechatOriId, wechatAppId, wechatToken, wechatEncodedAESKey,msgHandler, nil)
	msgServer.ServeHTTP(ctx.ResponseWriter(), ctx.Request(), nil)

}





