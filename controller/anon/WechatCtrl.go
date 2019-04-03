package controller

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
)

//101.132.125.102

func Login(ctx iris.Context) {
	// wechat parameters
	config := &wechat.Config{
		AppID: "wx6bb6950cf39d79ee",
		AppSecret: "25e017d8ab0f6711b5080be1ae317421",
		Token: "HMIIoT",
		EncodingAESKey:"iesxoHBsnaKVry5E8xd8gavmJLTVVNcd8aS7w3KYOaU",
	}

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
	config := &wechat.Config{
		AppID: "wx6bb6950cf39d79ee",
		AppSecret: "25e017d8ab0f6711b5080be1ae317421",
		Token: "HMIIoT",
		EncodingAESKey:"iesxoHBsnaKVry5E8xd8gavmJLTVVNcd8aS7w3KYOaU",
	}

	wc := wechat.NewWechat(config)
	wechatServer := wc.GetServer(ctx.Request(),ctx.ResponseWriter())

	wechatServer.SetMessageHandler(func(v message.MixMessage) *message.Reply {
		textResp := "HMI-IoT Not FOUND"
		switch v.MsgType {
		//文本消息
		case message.MsgTypeText:
			textResp = "HMI-IoT"
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

		return &message.Reply{MsgType: message.MsgTypeText, MsgData: textResp}
	})

}



