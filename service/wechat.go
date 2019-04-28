package service

import (
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/menu"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
	"github.com/chanxuehong/wechat/mp/message/callback/response"
	"github.com/chanxuehong/wechat/mp/oauth2"
	"github.com/chanxuehong/wechat/mp/user"
	oa2 "github.com/chanxuehong/wechat/oauth2"
	"github.com/pantazheng/bci/config"
	"github.com/pelletier/go-toml"
	"log"
)

var (
	ServerAddress  = config.Conf.Get("server.serverAddress").(string)
	WeChatConfigTree    =config.Conf.Get("wechat").(*toml.Tree)
	WeChatOriId         = WeChatConfigTree.Get("OriId").(string)
	WeChatAppId         = WeChatConfigTree.Get("AppId").(string)
	WeChatAppSecret     = WeChatConfigTree.Get("AppSecret").(string)
	WeChatToken         = WeChatConfigTree.Get("Token").(string)
	WeChatEncodedAESKey = WeChatConfigTree.Get("EncodedAESKey").(string)
	defaultClt          = wechatClient()
	tokenEndpoint       =oauth2.Endpoint{AppId: WeChatAppId, AppSecret: WeChatAppSecret}
)

func TextMsgHandler(ctx *core.Context) {
	log.Printf("进入文本消息处理")
	msg := request.GetText(ctx.MixedMsg)
	resp := response.NewText(msg.FromUserName, msg.ToUserName, msg.CreateTime, msg.Content)
	log.Printf("收到文本消息:\n%s\n", ctx.MsgPlaintext)
	if err:=ctx.RawResponse(resp);err!=nil{
		err.Error()
	}
}


func MenuClickEventHandler(ctx *core.Context) {
	log.Printf("收到按钮点击消息:\n%s\n", ctx.MsgPlaintext)
	event := menu.GetClickEvent(ctx.MixedMsg)
	resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, "请先登记个人信息")
	if err:=ctx.RawResponse(resp);err!=nil{
		err.Error()
	}
}

func SubscribeEventHandler(ctx *core.Context){
	event := request.GetSubscribeEvent(ctx.MixedMsg)
	clt := wechatClient()
	info,_:=user.Get(clt,event.FromUserName,"")
	resp := response.NewText(event.FromUserName,event.ToUserName,event.CreateTime, UserInit(info))
	if err:=ctx.RawResponse(resp);err!=nil{
		err.Error()
	}
}

func DefaultEventHandler(ctx *core.Context) {
	log.Printf("收到事件:\n%s\n", ctx.MsgPlaintext)
	_=ctx.NoneResponse()
}

func wechatClient() *core.Client{
	accessTokenTokenServer :=core.NewDefaultAccessTokenServer(WeChatAppId, WeChatAppSecret,nil)
	return core.NewClient(accessTokenTokenServer,nil)
}


func DefaultMenu(defaultMenu *menu.Menu){
	err:=menu.Create(defaultClt,defaultMenu)
	if err!=nil{
		err.Error()
	}
	log.Printf("建立默认菜单\n")
}

func ExchangeToken(token *oa2.Token,code string)(err error){
	exchangeClient:=&oa2.Client{}
	exchangeClient.Endpoint=&tokenEndpoint
	exchangeClient.Token=token
	token,_=exchangeClient.ExchangeToken(code)
	return
}


