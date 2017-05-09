package wechat

type ErrInfo struct {
    ErrCode         int         `json:errocde`
    ErrMsg          string      `json"errmsg`
}



func (self *ErrInfo) Set(code int,msg string) {
    self.ErrCode = code
    self.ErrMsg = msg
}

func (self *ErrInfo) GetCode() int {
    return self.ErrCode
}

func (self *ErrInfo) GetMsg() string{
    return self.ErrMsg
}
