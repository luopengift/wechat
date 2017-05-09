//素材
package wechat

import (
    "github.com/luopengift/gohttp"
    "fmt"
)

func (self *WeChatCtx) upload() {
    //gohttp.MyClient("POST",fmt.Sprintf(Url(self.GetType(),"UploadMedia"),self.GetToken()),nil,nil,nil)
    url := fmt.Sprintf(Url(self.GetType(),"UploadMedia"),self.GetToken())
    gohttp.NewClient().URL(url).Post()
}


