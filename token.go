package wechat

import (
    "fmt"
    "github.com/luopengift/gohttp"
    "github.com/luopengift/golibs/logger"
    "encoding/json"
    "time"
)

const (
    QY     = 1 << iota //企业号
    MP                  //公众号(订阅号)
    SR                  //公众号(服务号)

)

type Token struct {
    AppID           string      `json:"appid"`
    AppSecret       string      `json:"secret"`
    AccessToken     string      `json:"access_token"`
    ExpiresIn       int64       `json:"expires_in"`
    ExpiresTime     int64       `json:"expires_time"`
    _type           int64       `json:"-"`              //微信服务类型
}

func (self *Token) SetType(tt int64) {
    self._type = tt
}

func (self *Token) GetType() int64 {
    return self._type
}

func (self *Token) GetToken() string {
    if self.ExpiresTime <= time.Now().Unix() || self.ExpiresIn == 0 {
        url := fmt.Sprintf(Url(self.GetType(),"GetToken"),self.AppID,self.AppSecret)
        resp,_ := gohttp.NewClient().URL(url).Get()//("GET",url,nil,nil,nil)
        bytes,_ := resp.Bytes()
        err := json.Unmarshal(bytes,self)
        if err != nil {
            logger.Error("Token Error:%v,%v",err,resp.String())
            return ""
        }
        self.ExpiresTime = time.Now().Unix() + self.ExpiresIn
        logger.Info("TOKEN is",self.AccessToken)
    }
    return self.AccessToken
}

