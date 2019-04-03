package controller

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"github.com/silenceper/wechat/menu"
	"github.com/silenceper/wechat/message"
)

const weburl="101.132.125.102"



var config = &wechat.Config{
	AppID: "wx6bb6950cf39d79ee",
	AppSecret: "25e017d8ab0f6711b5080be1ae317421",
	Token: "HMIIoT",
	EncodingAESKey:"iesxoHBsnaKVry5E8xd8gavmJLTVVNcd8aS7w3KYOaU",
	Cache: cache.NewMemory(),
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
	flag := false
	wc := wechat.NewWechat(config)
	wechatServer := wc.GetServer(ctx.Request(),ctx.ResponseWriter())

	wechatServer.SetMessageHandler(func(v message.MixMessage) *message.Reply {
		textResp := "HMI-IoT Not FOUND"
		switch v.MsgType {
		//文本消息
		case message.MsgTypeText:
			flag=true
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

	if flag {
		wechatMenu := wc.GetMenu()
		//btnPlaceholder := new (menu.Button)
		//btnPlaceholder.SetViewButton("项目/任务","www.baidu.com")
		//btnCreate := new (menu.Button)
		//btnCreate.SetViewButton("创建用户","www.github.com")
		//btnWeekly := new (menu.Button)
		//btnWeekly.SetViewButton("周报","www.sohu.com")
		//buttonsSub := make([]* menu.Button,2)
		//buttonsSub[0]=btnCreate
		//buttonsSub[1]=btnWeekly
		//btnPerson := new (menu.Button)
		//btnPerson.SetClickButton("个人","person")
		//btnPerson.SetSubButton("subButton",buttonsSub)
		//buttonsDefault := make ([]* menu.Button, 2)
		//buttonsDefault[0]=btnPlaceholder
		//buttonsDefault[1]=btnPerson
		buttons := make([]*menu.Button, 1)
		btn := new(menu.Button)

		//创建click类型菜单
		btn.SetClickButton("name", "key123")
		buttons[0] = btn

		//设置btn为二级菜单
		btn2 := new(menu.Button)
		btn2.SetSubButton("subButton", buttons)

		buttons2 := make([]*menu.Button, 1)
		buttons2[0] = btn

		err2 := wechatMenu.SetMenu(buttons)
		if err2 != nil {
			fmt.Printf("***err= %v", err2)
		}
	}

}



