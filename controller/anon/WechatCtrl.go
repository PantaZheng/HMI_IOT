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

	server := wc.GetServer(ctx.Request(),ctx.ResponseWriter())
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {

		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()


}


