package wechat

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/luopengift/gohttp"
	"github.com/luopengift/log"
)

const (
	// QY qy 企业号
	QY = 1 << iota //企业号
	// MP mp
	MP //公众号(订阅号)
	// SR sr
	SR //公众号(服务号)

)

// Token token
type Token struct {
	AppID       string `json:"appid"`
	AppSecret   string `json:"secret"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	ExpiresTime int64  `json:"expires_time"`
	Type        int64  `json:"-"` //微信服务类型
}

// SetType set type
func (token *Token) SetType(tt int64) {
	token.Type = tt
}

// GetType get type
func (token *Token) GetType() int64 {
	return token.Type
}

// GetToken token
func (token *Token) GetToken() string {
	if token.ExpiresTime <= time.Now().Unix() || token.ExpiresIn == 0 {
		url := fmt.Sprintf(URL(token.GetType(), "GetToken"), token.AppID, token.AppSecret)
		resp, _ := gohttp.NewClient().URLString(url).Get() //("GET",url,nil,nil,nil)
		bytes := resp.Bytes()
		err := json.Unmarshal(bytes, token)
		if err != nil {
			log.Error("Token Error:%v,%v", err, resp.String())
			return ""
		}
		token.ExpiresTime = time.Now().Unix() + token.ExpiresIn
		log.Info("TOKEN is %s", token.AccessToken)
	}
	return token.AccessToken
}
