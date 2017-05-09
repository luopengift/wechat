package wechat

import (
    "fmt"
    "github.com/luopengift/gohttp"
    "github.com/luopengift/golibs/logger"
    "encoding/json"
)


type WeChatCtx struct {
    *Token
    *ErrInfo
    *Department
    *UrlMap
    ServerList      []string    `json:"serverlist"`
}


func (self *WeChatCtx) GetServerList() []string {
    url := fmt.Sprintf(Url(self.GetType(),"GetServerList"),self.GetToken())
    //url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/getcallbackip?access_token=%s",self.GetToken())
    resp,err := gohttp.NewClient().URL(url).Get()
    if err != nil || resp.Code() != 200 {
        logger.Error("NewClient:%v,%v",err,resp.String())
        return []string{}
    }
    list := map[string][]string{}
    bytes,_ := resp.Bytes()
    err = json.Unmarshal(bytes,&list)
    if err != nil {
        logger.Error("serverlist:%v,%v",err,resp.String())
        return self.ServerList
    }
    if value,ok := list["ip_list"]; ok {
        self.ServerList = value
    }
    logger.Info("",self.ServerList)
    return self.ServerList
}



func NewWeChatCtx(appid, appsecret string) *WeChatCtx {
    return &WeChatCtx{&Token{AppID:appid,AppSecret:appsecret},&ErrInfo{},&Department{},&UrlMap{},[]string{}}
}


