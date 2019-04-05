package service

import (
	"../config"
	"fmt"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/menu"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
	"github.com/chanxuehong/wechat/mp/message/callback/response"
	"github.com/chanxuehong/wechat/mp/user"
	"github.com/kataras/iris"
	"github.com/pelletier/go-toml"
	"log"
)

var (
	wechatConfigTree =config.Conf.Get("wechat").(*toml.Tree)
	wechatOriId = wechatConfigTree.Get("OriId").(string)
	wechatAppId = wechatConfigTree.Get("AppId").(string)
	wechatAppSecret = wechatConfigTree.Get("AppSecret").(string)
	wechatToken = wechatConfigTree.Get("Token").(string)
	wechatEncodedAESKey = wechatConfigTree.Get("EncodedAESKey").(string)
)

func TextMsgHandler(ctx *core.Context) {
	log.Printf("收到文本消息:\n%s\n", ctx.MsgPlaintext)

	msg := request.GetText(ctx.MixedMsg)
	resp := response.NewText(msg.FromUserName, msg.ToUserName, msg.CreateTime, msg.Content)
	ctx.RawResponse(resp) // 明文回复

	clt := wechatClient()
	if msg.Content =="menu create"{
		defaultMenu(clt)
	}
	//ctx.AESResponse(resp, 0, "", nil) // aes密文回复
}

func DefaultMsgHandler(ctx *core.Context) {
	log.Printf("收到消息:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

func MenuClickEventHandler(ctx *core.Context) {
	log.Printf("收到菜单 click 事件:\n%s\n", ctx.MsgPlaintext)

	event := menu.GetClickEvent(ctx.MixedMsg)
	resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, "收到 click 类型的事件")
	//ctx.RawResponse(resp) // 明文回复
	ctx.AESResponse(resp, 0, "", nil) // aes密文回复
}

func SubscribeEventHandler(ctx *core.Context){
	log.Printf("收到订阅:\n%s\n", ctx.MsgPlaintext)

	event := request.GetSubscribeEvent(ctx.MixedMsg)
	clt := wechatClient()
	info,_:=user.Get(clt,event.FromUserName,"")
	resp := response.NewText(event.FromUserName,event.ToUserName,event.CreateTime,SubscribeInit(info))
	if err:=ctx.RawResponse(resp);err!=nil{
		fmt.Printf("%v",err)
	}
}


func DefaultEventHandler(ctx *core.Context) {
	log.Printf("收到事件:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

func WechatServer(ctx iris.Context) {
	mux := core.NewServeMux()
	mux.DefaultMsgHandleFunc(DefaultEventHandler)
	mux.DefaultEventHandleFunc(DefaultEventHandler)
	mux.MsgHandleFunc(request.MsgTypeText, TextMsgHandler)
	mux.EventHandleFunc(menu.EventTypeClick, MenuClickEventHandler)

	msgHandler := mux

	msgServer := core.NewServer(wechatOriId, wechatAppId, wechatToken, wechatEncodedAESKey,msgHandler, nil)
	msgServer.ServeHTTP(ctx.ResponseWriter(), ctx.Request(), nil)
}

func wechatClient() *core.Client{
	accessTokenTokenServer :=core.NewDefaultAccessTokenServer(wechatAppId,wechatAppSecret,nil)
	return core.NewClient(accessTokenTokenServer,nil)
}

func defaultMenu(clt *core.Client){
	btnProjectMission:=menu.Button{}
	btnProjectMission.SetAsClickButton("Project/Mission","www.baidu.com")
	buttonsDefault:=make([]menu.Button,1)
	buttonsDefault[0]=btnProjectMission
	defaultMenu:=menu.Menu{}
	defaultMenu.Buttons=buttonsDefault
	err:=menu.Create(clt,&defaultMenu)
	if err!=nil{
		fmt.Printf("%v",err)
	}
}