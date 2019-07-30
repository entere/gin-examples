package delivery

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"github.com/silenceper/wechat/message"
	"github.com/silenceper/wechat/qr"
	"github.com/silenceper/wechat/template"
)

// WechatDelivery ...
var WechatDelivery = newWechatDelivery()

const (
	baseURL = "http://fv6qb2.natappfree.cc"
)

// Scene ...
type Scene struct {
	SceneStr string `json:"scene_str,omitempty"`
	SceneID  int    `json:"scene_id,omitempty"`
}

type wechatDelivery struct {
	Wechat *wechat.Wechat
}

func newWechatDelivery() *wechatDelivery {
	//配置微信参数
	memCache := cache.NewMemcache("127.0.0.1:11211") // memCache必须开启，否则发送模板消息等不可作用
	config := &wechat.Config{
		AppID:          "wxd579834b7f52b7fa",
		AppSecret:      "c22c22d1104493814944e870349a4303",
		Token:          "codeisland_token",
		EncodingAESKey: "your encoding aes key",
		Cache:          memCache,
	}
	wc := wechat.NewWechat(config)
	return &wechatDelivery{
		Wechat: wc,
	}
}

func (wd *wechatDelivery) OauthHandler(c *gin.Context) {

	oauth := wd.Wechat.GetOauth()
	err := oauth.Redirect(c.Writer, c.Request, baseURL+"/v1/wechat/oauth/redirect", "snsapi_userinfo", "123dd123")
	if err != nil {
		fmt.Println(err)
	}

}

func (wd *wechatDelivery) OauthRedirectHandler(c *gin.Context) {
	oauth := wd.Wechat.GetOauth()
	code := c.Query("code")
	resToken, err := oauth.GetUserAccessToken(code)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("%+v\n", resToken)
	userInfo, err := oauth.GetUserInfo(resToken.AccessToken, resToken.OpenID)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("%+v\n", userInfo)

	c.JSON(http.StatusOK, userInfo)

}

// 微信接口配置这个地址来处理消息 开发-基本配置-服务器配置-服务器地址(URL)
func (wd *wechatDelivery) ServerHandler(c *gin.Context) {
	// 传入request和responseWriter
	user := wd.Wechat.GetUser()
	server := wd.Wechat.GetServer(c.Request, c.Writer)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		// log.Printf("%+v\n", msg)
		switch msg.MsgType {

		//文本消息
		case message.MsgTypeText:
			//do something
			//回复消息：演示回复用户发送的消息
			text := message.NewText(msg.Content)
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

			//图片消息
		case message.MsgTypeImage:
			//do something

			//语音消息
		case message.MsgTypeVoice:
			//do something

			//视频消息
		case message.MsgTypeVideo:
			//do something

			//小视频消息
		case message.MsgTypeShortVideo:
			//do something

			//地理位置消息
		case message.MsgTypeLocation:
			//do something

			//链接消息
		case message.MsgTypeLink:
			//do something

			//事件推送消息
		case message.MsgTypeEvent:
			switch msg.Event {
			//EventSubscribe 订阅
			case message.EventSubscribe:
				openID := server.GetOpenID()
				userInfo, err := user.GetUserInfo(openID)
				if err != nil {
					log.Println(err)
					return nil
				}
				log.Printf("关注者openID为：" + openID)
				log.Printf("%+v", userInfo)

				return &message.Reply{MsgType: message.MsgTypeText, MsgData: "关注公号登录成功"}
				//do something

				//取消订阅
			case message.EventUnsubscribe:
				log.Printf("取消关注")
				return &message.Reply{MsgType: message.MsgTypeText, MsgData: "取消关注"}
				//do something

				//用户已经关注公众号，则微信会将带场景值扫描事件推送给开发者
			case message.EventScan:
				//do something

				// 上报地理位置事件
			case message.EventLocation:
				//do something

				// 点击菜单拉取消息时的事件推送
			case message.EventClick:
				//do something

				// 点击菜单跳转链接时的事件推送
			case message.EventView:
				//do something

				// 扫码推事件的事件推送
			case message.EventScancodePush:
				//do something

				// 扫码推事件且弹出“消息接收中”提示框的事件推送
			case message.EventScancodeWaitmsg:
				//do something

				// 弹出系统拍照发图的事件推送
			case message.EventPicSysphoto:
				//do something

				// 弹出拍照或者相册发图的事件推送
			case message.EventPicPhotoOrAlbum:
				//do something

				// 弹出微信相册发图器的事件推送
			case message.EventPicWeixin:
				break
				//do something

				// 弹出地理位置选择器的事件推送
			case message.EventLocationSelect:
				break
				//do something

			}

		}

		return nil

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

// 发送模板消息，配置好后在浏览器执行 /v1/wechat/template即可
func (wd *wechatDelivery) TemplateHandler(c *gin.Context) {
	// 传入request和responseWriter
	itemData := map[string]*template.DataItem{
		"First": &template.DataItem{
			Value: "恭喜你购买成功",
			Color: "#173177",
		},
		"Product": &template.DataItem{
			Value: "Python入门",
			Color: "#173177",
		},
		"Amount": &template.DataItem{
			Value: "9.9元",
			Color: "#173177",
		},
		"Time": &template.DataItem{
			Value: "2019年7月25日",
			Color: "#173177",
		},
		"Remark": &template.DataItem{
			Value: "加入幻想编程岛，一起改变世界，GOGOGO！",
			Color: "#173177",
		},
	}
	msg := &template.Message{
		ToUser:     "oO8C8vkXDDjt9OqBP3GpcPi_RKAU", // openid
		TemplateID: "OGw41hl2pxfbGmmvK4801ih6y5EtVeyh3MJHu9pUmMY",
		URL:        "http://www.baidu.com",
		// Color:      "",
		Data: itemData,
	}
	template := wd.Wechat.GetTemplate()
	template.Send(msg)

}

// @Summary 获取临时二维码
// @Description 注册
// @Accept  json
// @Produce  json
// @Param   scene_id scene_str exp
// @Success 0 {string} string    "ok"
// @Router /wechat/qrcode/tmp?scene_id=100001&scene_str=unknow [get]
func (wd *wechatDelivery) QRCodeTmpHandler(c *gin.Context) {
	// 传入request和responseWriter
	sceneIDStr := c.DefaultQuery("scene_id", "100001")
	sceneStr := c.DefaultQuery("scene_str", "unknow")
	sceneID, err := strconv.Atoi(sceneIDStr)
	if err != nil {
		log.Println(err)
		return
	}
	// scene := Scene{
	// 	SceneStr: sceneStr,
	// 	SceneID:  sceneID,
	// }
	// request := qr.NewTmpQrRequest(30, scene) // 功能和下面注解一样 生成临时二维码的reqeust
	var request qr.Request
	// ActionName	二维码类型，QR_SCENE为临时的整型参数值，QR_STR_SCENE为临时的字符串参数值，QR_LIMIT_SCENE为永久的整型参数值，QR_LIMIT_STR_SCENE为永久的字符串参数值
	request = qr.Request{
		ExpireSeconds: 30,
		ActionName:    "QR_STR_SCENE",
	}
	request.ActionInfo.Scene.SceneStr = sceneStr
	request.ActionInfo.Scene.SceneID = sceneID
	qrcode := wd.Wechat.GetQR()
	ticket, err := qrcode.GetQRTicket(&request)
	log.Println(ticket)
	if err != nil {
		log.Println(err)
	}
	url := qr.ShowQRCode(ticket)
	c.JSON(200, url)

}

// @Summary 获取永久二维码
// @Description 获取永久二维码
// @Accept  json
// @Produce  json
// @Param   scene_id scene_str
// @Success 0 {string} string    "ok"
// @Router /wechat/qrcode/limit?scene_id=1&scene_str=hxbcd [get]
func (wd *wechatDelivery) QRCodeLimitHandler(c *gin.Context) {
	// 传入request和responseWriter

	sceneIDStr := c.DefaultQuery("scene_id", "2")
	sceneStr := c.DefaultQuery("scene_str", "msqx")
	sceneID, err := strconv.Atoi(sceneIDStr)
	if err != nil {
		log.Println(err)
		return
	}
	// scene := Scene{
	// 	SceneStr: sceneStr,
	// 	SceneID:  sceneID,
	// }
	// request := qr.NewTmpQrRequest(30, scene) // 功能和下面注解一样 生成临时二维码的reqeust
	var request qr.Request
	// ActionName	二维码类型，QR_SCENE为临时的整型参数值，QR_STR_SCENE为临时的字符串参数值，QR_LIMIT_SCENE为永久的整型参数值，QR_LIMIT_STR_SCENE为永久的字符串参数值
	request = qr.Request{
		ActionName: "QR_LIMIT_STR_SCENE",
	}
	request.ActionInfo.Scene.SceneStr = sceneStr
	request.ActionInfo.Scene.SceneID = sceneID
	qrcode := wd.Wechat.GetQR()
	ticket, err := qrcode.GetQRTicket(&request)
	log.Println(ticket)
	if err != nil {
		log.Println(err)
	}
	url := qr.ShowQRCode(ticket)
	c.JSON(200, url)

}
