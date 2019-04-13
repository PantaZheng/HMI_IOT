package service

import (
	"../config"
	"fmt"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/menu"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
	"github.com/chanxuehong/wechat/mp/message/callback/response"
	"github.com/chanxuehong/wechat/mp/user"
	"github.com/chanxuehong/wechat/mp/user/tag"
	"github.com/kataras/iris"
	"github.com/pelletier/go-toml"
	"log"
	"strconv"
)

var (
	wechatConfigTree =config.Conf.Get("wechat").(*toml.Tree)
	wechatOriId = wechatConfigTree.Get("OriId").(string)
	wechatAppId = wechatConfigTree.Get("AppId").(string)
	wechatAppSecret = wechatConfigTree.Get("AppSecret").(string)
	wechatToken = wechatConfigTree.Get("Token").(string)
	wechatEncodedAESKey = wechatConfigTree.Get("EncodedAESKey").(string)
	defaultClt = wechatClient()
	tagTeacher = 0
	tagStudent = 0
)

func TextMsgHandler(ctx *core.Context) {
	msg := request.GetText(ctx.MixedMsg)
	resp := response.NewText(msg.FromUserName, msg.ToUserName, msg.CreateTime, msg.Content)
	log.Printf("收到文本消息:\n%s\n", ctx.MsgPlaintext)
	_=ctx.RawResponse(resp)
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

//向微信服务器创建Tag
func CreateTag(){
	log.Printf("CreateTag\n")
	_,err:=tag.Create(defaultClt,"student")
	if err!=nil{
		fmt.Printf("\nCreateTag%v\n",err)
	}
	_,err=tag.Create(defaultClt,"teacher")
	if err!=nil{
		fmt.Printf("\nCreateTag%v\n",err)
	}
}

//获得Tag的ID
func GetTagList(){
	log.Printf("GetTagList\n")
	tagList,err:=tag.List(defaultClt)
	if err!=nil{
		fmt.Printf("%v",err)
	}
	for _,v :=range tagList{
		if v.Name == "student"{
			tagStudent=v.Id
		}
		if v.Name == "teacher"{
			tagTeacher=v.Id
		}
		fmt.Printf("\nid: "+strconv.Itoa(v.Id)+"\tname: "+v.Name+""+"\tcount: "+strconv.Itoa(v.UserCount)+"\n")
	}
}

func AddRoleTag(weChatOpenId string, tagId int){
	if err:=tag.BatchTag(defaultClt,[]string{weChatOpenId},tagId);err!=nil{
		fmt.Printf("AddRoleTagError:%v\n",err)
	}
	log.Printf("AddRoleTag\t"+weChatOpenId+"\ttagId"+strconv.Itoa(tagId)+"\n")
	GetTagList()
}

func DelRoleTag(weChatOpenId string, tagId int){
	if err:=tag.BatchUntag(defaultClt,[]string{weChatOpenId},tagId);err!=nil{
		fmt.Printf("AddRoleTagError:%v\n",err)
	}
	log.Printf("DelRoleTag+\t"+weChatOpenId+"\ttagID:"+strconv.Itoa(tagId)+"\n")
	GetTagList()
}

func DefaultMenu(){
	btnRelationShip:=menu.Button{}
	btnRelationShip.SetAsClickButton("架构","RelationShip")
	btnProjectMission:=menu.Button{}
	btnProjectMission.SetAsClickButton("项目/任务","ProjectMission")
	btnEnroll:=menu.Button{}
	btnEnroll.SetAsViewButton("登记","https://open.weixin.qq.com/connect/oauth2/authorize?appid="+wechatAppId+"&redirect_uri=http://bci.renjiwulian.com/test&response_type=code&scope=snsapi_base&state=12#wechat_redirect")
	defaultButtons:= []menu.Button{btnRelationShip,btnProjectMission,btnEnroll}
	defaultMenu:=menu.Menu{}
	defaultMenu.Buttons= defaultButtons
	err:=menu.Create(defaultClt,&defaultMenu)
	if err!=nil{
		fmt.Printf("DefaultMenu%v\n",err)
	}
}

func TeacherMenu(){
	btnRelationShip:=menu.Button{}
	btnRelationShip.SetAsViewButton("架构","http://bci.renjiwulian.com/project")
	btnProject:=menu.Button{}
	btnProject.SetAsViewButton("项目","http://bci.renjiwulian.com/project")
	btnPersonal:=menu.Button{}
	btnPersonal.SetAsViewButton("个人","http://bci.renjiwulian.com/project")
	teacherButtons := []menu.Button{btnRelationShip, btnProject,btnPersonal}
	teacherRule :=menu.MatchRule{}
	teacherRule.TagId=strconv.Itoa(tagTeacher)
	teacherMenu :=menu.Menu{}
	teacherMenu.Buttons= teacherButtons
	teacherMenu.MatchRule=&teacherRule
	_,err:=menu.AddConditionalMenu(defaultClt,&teacherMenu)
	if err!=nil{
		fmt.Printf("\nTeacherMenu:%v\n",err)
	}
}

func StudentMenu(){
	btnRelationShip:=menu.Button{}
	btnRelationShip.SetAsViewButton("架构","http://bci.renjiwulian.com/project")
	btnMission :=menu.Button{}
	btnMission.SetAsViewButton("任务","http://bci.renjiwulian.com/project")
	btnPersonal:=menu.Button{}
	btnPersonal.SetAsViewButton("个人","http://bci.renjiwulian.com/project")
	studentButtons := []menu.Button{btnRelationShip, btnMission,btnPersonal}
	studentRule := menu.MatchRule{}
	studentRule.TagId=strconv.Itoa(tagStudent)
	studentMenu := menu.Menu{}
	studentMenu.Buttons= studentButtons
	studentMenu.MatchRule = &studentRule
	_,err:=menu.AddConditionalMenu(defaultClt,&studentMenu)
	if err!=nil{
		fmt.Printf("\nStudentMenu%v\n",err)
	}
}

func TestMenu(){
	M,err:=menu.TryMatch(defaultClt,"oPKFh5lM9MA6_Svd39Km-84no7c8")
	if err!=nil{
		fmt.Printf("%v\n",err)
	}else{
		for _, v:=range M.Buttons{
			fmt.Printf( v.Name+"\t")
		}
	}
	fmt.Printf("\n")
}

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