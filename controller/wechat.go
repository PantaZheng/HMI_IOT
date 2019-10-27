package controller

import (
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/menu"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
	"github.com/kataras/iris"
	"github.com/pantazheng/bci/config"
	"github.com/pantazheng/bci/service"
)

var (
	wechatConfig = config.Conf.Wechat
	vueAddress   = config.Conf.APP.Address + "vue"
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
	mux.EventHandleFunc(request.EventTypeSubscribe, service.SubscribeEventHandler)
	mux.EventHandleFunc(request.EventTypeUnsubscribe, service.UnsubscribeEventHandler)
	msgHandler := mux

	msgServer := core.NewServer(wechatConfig.OriID, wechatConfig.AppID, wechatConfig.Token, wechatConfig.EncodedAESKEY, msgHandler, nil)
	msgServer.ServeHTTP(ctx.ResponseWriter(), ctx.Request(), nil)
}

func Menu() {
	btnBinding := menu.Button{}
	btnBinding.SetAsViewButton("绑定", "https://open.weixin.qq.com/connect/oauth2/authorize?appid="+wechatConfig.AppID+"&response_type=code&scope=snsapi_base&redirect_uri="+vueAddress+"/index/#/&state=12#wechat_redirect")
	btnFrame := menu.Button{}
	btnFrame.SetAsViewButton("架构", vueAddress+"/index/#/frame")
	btnPerson := menu.Button{}
	btnPerson.SetAsSubMenuButton("成员", []menu.Button{btnBinding, btnFrame})

	btnArrangeProject := menu.Button{}
	btnArrangeProject .SetAsViewButton("项目", vueAddress+"/index/#/checkProject")
	btnArrangeTask := menu.Button{}
	btnArrangeTask.SetAsViewButton("课题", vueAddress+"/index/#/checkTask")
	btnArrange := menu.Button{}
	btnArrange.SetAsSubMenuButton("安排", []menu.Button{btnArrangeProject,btnArrangeTask})

	btnProject := menu.Button{}
	btnProject.SetAsViewButton("项目", vueAddress+"/index/#/project")
	btnTask := menu.Button{}
	btnTask.SetAsViewButton("课题", vueAddress+"/index/#/taskList")
	btnMission := menu.Button{}
	btnMission.SetAsViewButton("任务", vueAddress+"/index/#/mission")
	btnView := menu.Button{}
	btnView.SetAsSubMenuButton("查看", []menu.Button{btnProject,btnTask,btnMission})
	service.DefaultMenu(&menu.Menu{Buttons: []menu.Button{btnPerson, btnArrange, btnView}})
}
