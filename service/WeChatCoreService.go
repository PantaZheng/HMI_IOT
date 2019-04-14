package service

import (
	"../config"
	"fmt"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/menu"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
	"github.com/chanxuehong/wechat/mp/message/callback/response"
	"github.com/chanxuehong/wechat/mp/oauth2"
	oa2 "github.com/chanxuehong/wechat/oauth2"
	"github.com/chanxuehong/wechat/mp/user"
	"github.com/chanxuehong/wechat/mp/user/tag"
	"github.com/kataras/iris"
	"github.com/pelletier/go-toml"
	"log"
	"strconv"
)

var (
	serverAddress       = config.Conf.Get("server.serverAddress").(string)
	wechatConfigTree    =config.Conf.Get("wechat").(*toml.Tree)
	wechatOriId         = wechatConfigTree.Get("OriId").(string)
	wechatAppId         = wechatConfigTree.Get("AppId").(string)
	wechatAppSecret     = wechatConfigTree.Get("AppSecret").(string)
	wechatToken         = wechatConfigTree.Get("Token").(string)
	wechatEncodedAESKey = wechatConfigTree.Get("EncodedAESKey").(string)
	defaultClt          = wechatClient()
	tokenEndpoint =oauth2.Endpoint{wechatAppId,wechatAppSecret}
	tagTeacher          = 0
	tagStudent          = 0
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
		fmt.Printf("MenuClickEventHandlerERR:%v", err)
	}
}

func SubscribeEventHandler(ctx *core.Context){
	event := request.GetSubscribeEvent(ctx.MixedMsg)
	clt := wechatClient()
	info,_:=user.Get(clt,event.FromUserName,"")
	resp := response.NewText(event.FromUserName,event.ToUserName,event.CreateTime, UserInit(info))
	if err:=ctx.RawResponse(resp);err!=nil{
		fmt.Printf("SubscribeEventHandlerERR:%v",err)
	}

}

func DefaultEventHandler(ctx *core.Context) {
	log.Printf("收到事件:\n%s\n", ctx.MsgPlaintext)
	_=ctx.NoneResponse()
}

func WechatServer(ctx iris.Context) {
	mux := core.NewServeMux()
	mux.DefaultMsgHandleFunc(DefaultEventHandler)
	mux.DefaultEventHandleFunc(DefaultEventHandler)
	mux.MsgHandleFunc(request.MsgTypeText, TextMsgHandler)
	mux.EventHandleFunc(menu.EventTypeClick, MenuClickEventHandler)
	mux.EventHandleFunc(request.EventTypeSubscribe,SubscribeEventHandler)
	msgHandler := mux

	msgServer := core.NewServer(wechatOriId, wechatAppId, wechatToken, wechatEncodedAESKey,msgHandler, nil)
	msgServer.ServeHTTP(ctx.ResponseWriter(), ctx.Request(), nil)
}

func wechatClient() *core.Client{
	accessTokenTokenServer :=core.NewDefaultAccessTokenServer(wechatAppId,wechatAppSecret,nil)
	return core.NewClient(accessTokenTokenServer,nil)
}


func DefaultMenu(){
	btnRelationShip:=menu.Button{}
	btnRelationShip.SetAsViewButton("架构", serverAddress+"/project/index.html")
	btnProjectMission:=menu.Button{}
	btnProjectMission.SetAsViewButton("项目/任务", serverAddress+"/project/index.html")
	btnEnroll:=menu.Button{}
	btnEnroll.SetAsViewButton("登记","https://open.weixin.qq.com/connect/oauth2/authorize?appid="+wechatAppId+"&redirect_uri="+serverAddress+"/test/test.html&response_type=code&scope=snsapi_base&state=12#wechat_redirect")
	defaultButtons:= []menu.Button{btnRelationShip,btnProjectMission,btnEnroll}
	defaultMenu:=menu.Menu{}
	defaultMenu.Buttons= defaultButtons
	err:=menu.Create(defaultClt,&defaultMenu)
	if err!=nil{
		fmt.Printf("DefaultMenu%v\n",err)
	}
	log.Printf("建立默认菜单\n")
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

func DelAllConditionalMenu(){
	_,m2,err:=menu.Get(defaultClt)
	if err!=nil{
		fmt.Printf("%v\n",err)
	}else {
		for _,v1:=range  m2 {
			_=menu.DeleteConditionalMenu(defaultClt,v1.MenuId)
		}
	}
}

func GetAllMenu(){
	m1,m2,err:=menu.Get(defaultClt)
	if err!=nil{
		fmt.Printf("%v\n",err)
	}else{
		fmt.Printf("defaultMenus----\n")
		for _, v:=range m1.Buttons{
			fmt.Printf( v.Name+"\t")
		}
		fmt.Printf("\n")
		fmt.Printf("conditionalMenus----\n")
		for _,v1:=range  m2{
			for _,v3:=range v1.Buttons{
				fmt.Printf( v3.Name+"\t")
			}
			fmt.Printf("\n")
		}
	}
	fmt.Printf("\n")
}

func ExchangeToken(token *oa2.Token,code string)(err error){
	exchangeClient:=&oa2.Client{}
	exchangeClient.Endpoint=&tokenEndpoint
	exchangeClient.Token=token
	token,_=exchangeClient.ExchangeToken(code)
	return
}