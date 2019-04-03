package controller

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/menu"
	"github.com/silenceper/wechat/message"
)

const weburl="101.132.125.102"

var config = &wechat.Config{
	AppID: "wx6bb6950cf39d79ee",
	AppSecret: "25e017d8ab0f6711b5080be1ae317421",
	Token: "HMIIoT",
	EncodingAESKey:"iesxoHBsnaKVry5E8xd8gavmJLTVVNcd8aS7w3KYOaU",
}

func Login(ctx iris.Context) {
	// wechat parameters

	wc := wechat.NewWechat(config)

	wechatServer := wc.GetServer(ctx.Request(),ctx.ResponseWriter())

	//处理消息接收以及回复
	err := wechatServer.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	_=wechatServer.Send()

}

func Port(ctx iris.Context){

	wc := wechat.NewWechat(config)
	wechatServer := wc.GetServer(ctx.Request(),ctx.ResponseWriter())

	wechatServer.SetMessageHandler(func(v message.MixMessage) *message.Reply {
		textResp := "HMI-IoT Not FOUND"
		switch v.MsgType {
		//文本消息
		case message.MsgTypeText:

			wechatMenu := wc.GetMenu()

			btnPlaceholder := new (menu.Button)
			btnPlaceholder.SetViewButton("项目/任务","")
			btnCreate := new (menu.Button)
			btnCreate.SetViewButton("创建用户","")
			btnWeekly := new (menu.Button)
			btnWeekly.SetViewButton("周报","")
			buttonsSub := make([]* menu.Button,2)
			buttonsSub[0]=btnCreate
			buttonsSub[1]=btnWeekly
			btnPerson := new (menu.Button)
			btnPerson.SetClickButton("个人","person")
			btnPerson.SetSubButton("subButton",buttonsSub)
			defaultButtons := make ([]* menu.Button, 2)
			defaultButtons[0]=btnPlaceholder
			defaultButtons[1]=btnPerson

			err := wechatMenu.SetMenu(buttonsSub)
			if err != nil {
				fmt.Printf("err= %v", err)
			}

		case message.MsgTypeEvent:
			switch v.Event {
				//EventSubscribe 订阅
			case message.EventSubscribe:
				textResp = "感谢关注 HMI-IoT 人机物联"
				// 点击菜单跳转链接时的事件推送
			case message.EventView:
				//do something

			}
		}

		text := message.NewText(textResp)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := wechatServer.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	_=wechatServer.Send()
}



