package service

import (
	"../config"
	"fmt"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/menu"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
	"github.com/chanxuehong/wechat/mp/message/callback/response"
	"github.com/chanxuehong/wechat/mp/oauth2"
	"github.com/chanxuehong/wechat/mp/user"
	oa2 "github.com/chanxuehong/wechat/oauth2"
	"github.com/pelletier/go-toml"
	"log"
)

var (
	ServerAddress       = config.Conf.Get("server.ServerAddress").(string)
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
		log.Printf("文本消息处理出错：%v",err)
	}
}


func MenuClickEventHandler(ctx *core.Context) {
	log.Printf("收到按钮点击消息:\n%s\n", ctx.MsgPlaintext)
	event := menu.GetClickEvent(ctx.MixedMsg)
	resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, "请先登记个人信息")
	if err:=ctx.RawResponse(resp);err!=nil{
		panic(err.Error())
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
		fmt.Printf("DefaultMenu%v\n",err)
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

//func TestMenu(){
//	M,err:=menu.TryMatch(defaultClt,"oPKFh5lM9MA6_Svd39Km-84no7c8")
//	if err!=nil{
//		fmt.Printf("%v\n",err)
//	}else{
//		for _, v:=range M.Buttons{
//			fmt.Printf( v.Name+"\t")
//		}
//	}
//	fmt.Printf("\n")
//}

//func DelAllConditionalMenu(){
//	_,m2,err:=menu.Get(defaultClt)
//	if err!=nil{
//		fmt.Printf("%v\n",err)
//	}else {
//		for _,v1:=range  m2 {
//			_=menu.DeleteConditionalMenu(defaultClt,v1.MenuId)
//		}
//	}
//}

//func GetAllMenu(){
//	m1,m2,err:=menu.Get(defaultClt)
//	if err!=nil{
//		fmt.Printf("%v\n",err)
//	}else{
//		fmt.Printf("defaultMenus----\n")
//		for _, v:=range m1.Buttons{
//			fmt.Printf( v.Name+"\t")
//		}
//		fmt.Printf("\n")
//		fmt.Printf("conditionalMenus----\n")
//		for _,v1:=range  m2{
//			for _,v3:=range v1.Buttons{
//				fmt.Printf( v3.Name+"\t")
//			}
//			fmt.Printf("\n")
//		}
//	}
//	fmt.Printf("\n")
//}

