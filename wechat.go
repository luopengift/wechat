package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/luopengift/gohttp"
	"github.com/luopengift/log"
)

type WeChatCtx struct {
	*Token
	*ErrInfo
	*Department
	*Member
	*UrlMap
	ServerList []string `json:"serverlist"`
}

func (self *WeChatCtx) GetServerList() []string {
	url := fmt.Sprintf(Url(self.GetType(), "GetServerList"), self.GetToken())
	//url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/getcallbackip?access_token=%s",self.GetToken())
	resp, err := gohttp.NewClient().Url(url).Get()
	if err != nil || resp.Code() != 200 {
		log.Error("NewClient:%v,%v", err, resp.String())
		return []string{}
	}
	list := map[string][]string{}
	bytes := resp.Bytes()
	err = json.Unmarshal(bytes, &list)
	if err != nil {
		log.Error("serverlist:%v,%v", err, resp.String())
		return self.ServerList
	}
	if value, ok := list["ip_list"]; ok {
		self.ServerList = value
	}
	log.Info("%v", self.ServerList)
	return self.ServerList
}

func NewWeChatCtx(appid, appsecret string) *WeChatCtx {
	return &WeChatCtx{&Token{AppID: appid, AppSecret: appsecret}, &ErrInfo{}, &Department{}, &Member{}, &UrlMap{}, []string{}}
}
