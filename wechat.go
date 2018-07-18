package wechat

import (
	"encoding/json"
	"fmt"

	"github.com/luopengift/gohttp"
	"github.com/luopengift/log"
)

// Context wechat ctx
type Context struct {
	*Token
	*ErrInfo
	*Department
	*Member
	*URLMap
	ServerList []string `json:"serverlist"`
}

// GetServerList get server list
func (ctx *Context) GetServerList() []string {
	url := fmt.Sprintf(URL(ctx.GetType(), "GetServerList"), ctx.GetToken())
	//url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/getcallbackip?access_token=%s",self.GetToken())
	resp, err := gohttp.NewClient().URLString(url).Get()
	if err != nil || resp.Code() != 200 {
		log.Error("NewClient:%v,%v", err, resp.String())
		return []string{}
	}
	list := map[string][]string{}
	bytes := resp.Bytes()
	err = json.Unmarshal(bytes, &list)
	if err != nil {
		log.Error("serverlist:%v,%v", err, resp.String())
		return ctx.ServerList
	}
	if value, ok := list["ip_list"]; ok {
		ctx.ServerList = value
	}
	return ctx.ServerList
}

// NewWeChatCtx new wechat ctx
func NewWeChatCtx(appid, appsecret string) *Context {
	return &Context{&Token{AppID: appid, AppSecret: appsecret}, &ErrInfo{}, &Department{}, &Member{}, &URLMap{}, []string{}}
}
