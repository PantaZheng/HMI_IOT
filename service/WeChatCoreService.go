package service

import (
	"fmt"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/menu"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
	"github.com/chanxuehong/wechat/mp/message/callback/response"
	"log"
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

func DefaultEventHandler(ctx *core.Context) {
	log.Printf("收到事件:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

func wechatClient() *core.Client{
	accessTokenTokenServer :=core.NewDefaultAccessTokenServer("wx6bb6950cf39d79ee","25e017d8ab0f6711b5080be1ae317421",nil)
	return core.NewClient(accessTokenTokenServer,nil)
}

func defaultMenu(clt *core.Client){
	btnProjectMission:=menu.Button{}
	btnProjectMission.SetAsClickButton("ProjectMission","www.baidu.com")
	buttonsDefault:=make([]menu.Button,1)
	buttonsDefault[0]=btnProjectMission
	defaultMenu:=menu.Menu{}
	defaultMenu.Buttons=buttonsDefault
	err:=menu.Create(clt,&defaultMenu)
	if err!=nil{
		fmt.Printf("%v",err)
	}
}