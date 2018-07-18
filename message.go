package wechat

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/luopengift/gohttp"
)

// Agentid agentid
var Agentid = map[string]int{"企业小助手": 0, "监控告警": 1}

// AgentID agentid
var AgentID = map[int]string{0: "企业小助手", 1: "监控告警"}

// Image var msgType = [...]string{"text", "image", "voice", "video", "music", "file", "news", "mpnews"}
// Image image
type Image struct {
	MediaID string `xml:"MediaId"`
}

// Voice voice
type Voice struct {
	MediaID string `xml:"MediaId"`
}

// Video video
type Video struct {
	MediaID     string `xml:"MediaId"`
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
}

// Music music
type Music struct {
	Title        string `xml:"Title"`        //音乐标题
	Description  string `xml:"Description"`  //音乐描述
	MusicURL     string `xml:"MusicUrl"`     //音乐链接
	HQMusicURL   string `xml:"HQMusicUrl"`   //高质量音乐链接，WIFI环境优先使用该链接播放音乐
	ThumbMediaID string `xml:"ThumbMediaId"` //缩略图的媒体id，通过素材管理接口上传多媒体文件，得到的id
}

// Article article
type Article struct { //xml item
	Title       string `xml:"Title"`       //图文消息标题
	Description string `xml:"Description"` //图文消息描述
	PicURL      string `xml:"PicUrl"`      //图片链接，支持JPG、PNG格式，较好的效果为大图360*200，小图200*200
	URL         string `xml:"Url"`         //点击图文消息跳转链接
}

// MPMessage mp message
type MPMessage struct {
	*Context     `xml:"-"`
	ToUser       string    `xml:"ToUserName"`
	FromUser     string    `xml:"FromUserName"`
	CreateTime   int64     `xml:"CreateTime"`
	MsgType      string    `xml:"MsgType"`
	Text         string    `xml:"Content"`      //文本消息
	Image        Image     `xml:"Image"`        //图片消息
	Voice        Voice     `xml:"Voice"`        //语音消息
	Vodeo        Video     `xml:"Vodeo"`        //视频消息
	Music        Music     `xml:"Music"`        //音乐消息
	ArticleCount int       `xml:"ArticleCount"` //图文消息个数，限制为10条以内
	Articles     []Article `xml:"Articles"`
}

// SendText send text
func (ctx *MPMessage) SendText() {}

type text struct {
	Text string `xml:",cdata"` //文本消息
}

type msgType struct {
	Type string `xml:",cdata"`
}

type user struct {
	Name string `xml:",cdata"`
}

// Message message
type Message struct {
	XMLName    xml.Name `xml:"xml"`
	ToUser     user     `xml:"ToUserName"`
	FromUser   user     `xml:"FromUserName"`
	CreateTime int64    `xml:"CreateTime"`
	MsgType    msgType  `xml:"MsgType"`
	Content    text     `xml:"Content"` //文本消息
	Test       string   `xml:"TEST"`
}

// NewTextMsg new text msg
func NewTextMsg(from, to, content string) *Message {
	return &Message{
		ToUser:     user{to},
		FromUser:   user{from},
		CreateTime: time.Now().Unix(),
		MsgType:    msgType{"text"},
		Content:    text{content},
	}
}

// NewMusicMsg new music msg
func NewMusicMsg(from, to string) *Message {
	return nil
}

// ToUserName to user name
func (ctx *Message) ToUserName() string {
	return ctx.ToUser.Name
}

// FromUserName FromUserName
func (ctx *Message) FromUserName() string {
	return ctx.FromUser.Name
}

// Text text
func (ctx *Message) Text() string {
	return ctx.Content.Text
}

// Type type
func (ctx *Message) Type() string {
	return ctx.MsgType.Type
}

// QYMessage qy message
type QYMessage struct {
	*Context `json:"-" xml:"-"`
	ToUser   string      `json:"touser"`  //"UserID1|UserID2|UserID3",
	ToParty  string      `json:"toparty"` //" PartyID1 | PartyID2 ", QY
	ToTag    string      `json:"totag"`   //" TagID1 | TagID2 ",
	MsgType  string      `json:"msgtype"` //"text",
	AgentID  int         `json:"agentid"` //企业应用的id，整型。可在应用的设置页面查看
	Safe     int         `json:"safe"`    //表示是否是保密消息，0表示否，1表示是，默认0
	Text     interface{} `json:"text"`    //文本消息 {"content": "Holiday Request For Pony(http://xxxxx)"},
	Image    interface{} `json:"image"`   //图片消息 {"media_id": "MEDIA_ID"},
	Voice    interface{} `json:"voice"`   //语音消息 {"media_id": "MEDIA_ID"},
	Video    interface{} `json:"video"`   //视频消息 "video": {"media_id": "MEDIA_ID","title": "Title","description": "Description"},
	File     interface{} `json:"file"`    //File消息,企业号专有
	News     interface{} `json:"news"`    //图文消息
	MpNews   interface{} `json:"mpnews"`  //mpnews消息与news消息类似，不同的是图文消息内容存储在微信后台,企业号专有
}

var msg = QYMessage{
	ToUser:  "",
	ToParty: "",
	ToTag:   "",
	MsgType: "text",
	AgentID: 1,
	Safe:    0,
	Text:    map[string]string{"content": "状态:PROBLEM\n告警级别:Warning\n名称:PredictServer\nIP地址:172.31.30.165\n区域:ap-southeast-1\n监控项:DiskSpaceLess 20%  /data\n详细信息:Free disk space on /data (percentage):19.99 %\n错误信息:\n时间:2016.10.20.10:03:51"},
}

// SendText send text
func (ctx *QYMessage) SendText() string {
	ctx.MsgType = "text"
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", ctx.GetToken())
	resp, _ := gohttp.NewClient().URLString(url).Body(ctx).Header("Content-Type", "application/json;charset=utf-8").Post()
	return resp.String()
}
