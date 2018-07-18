package wechat

import (
	"fmt"

	"github.com/luopengift/gohttp"
)

func (ctx *Context) upload() {
	//gohttp.MyClient("POST",fmt.Sprintf(Url(self.GetType(),"UploadMedia"),self.GetToken()),nil,nil,nil)
	url := fmt.Sprintf(URL(ctx.GetType(), "UploadMedia"), ctx.GetToken())
	gohttp.NewClient().URLString(url).Post()
}
