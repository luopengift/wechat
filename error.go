package wechat

// ErrInfo err info
type ErrInfo struct {
	ErrCode int    `json:"errocde"`
	ErrMsg  string `json:"errmsg"`
}

// Set set
func (err *ErrInfo) Set(code int, msg string) {
	err.ErrCode = code
	err.ErrMsg = msg
}

// GetCode get code
func (err *ErrInfo) GetCode() int {
	return err.ErrCode
}

// GetMsg get msg
func (err *ErrInfo) GetMsg() string {
	return err.ErrMsg
}
