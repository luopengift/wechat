package wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/luopengift/gohttp"
	"github.com/luopengift/log"
)

// Member member
type Member struct {
	UserID        string      `json:"userid"`     //成员UserID。对应管理端的帐号，企业内必须唯一。不区分大小写，长度为1~64个字节
	Name          string      `json:"name"`       //成员名称。长度为1~64个字节
	Department    []int       `json:"department"` //成员所属部门id列表,不超过20个
	Position      string      `json:"position"`   //职位信息。长度为0~64个字节
	Mobile        string      `json:"mobile"`     //手机号码。企业内必须唯一，mobile/weixinid/email三者不能同时为空
	Gender        int         `json:"gender"`     //性别。1表示男性，2表示女性
	Email         string      `json:"email"`      //邮箱。长度为0~64个字节。企业内必须唯一
	WeiXinID      string      `json:"weixinid"`   //微信号。企业内必须唯一。（注意：是微信号，不是微信的名字）
	AvatarMediaID string      `json:"-"`          //`json:"avatar_mediaid"` //成员头像的mediaid，通过多媒体接口上传图片获得的mediaid
	ExtAttr       interface{} `json:"-"`          //`json:"extattr"`       //扩展属性。扩展属性需要在WEB管理端创建后才生效，否则忽略未知属性的赋值
	//": {"attrs":[{"name":"爱好","value":"旅游"},{"name":"卡号","value":"1234567234"}]}
}

// MemberCtx memberctx
type MemberCtx struct {
	*Context `json:"-"`
	*Member
}

// Create 创建成员
func (ctx *MemberCtx) Create() {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/user/create?access_token=%s", ctx.GetToken())
	member, err := json.Marshal(ctx)
	if err != nil {
		log.Error("json err:", err)
	}
	log.Debug("member", string(member))
	body := bytes.NewBuffer(member)
	resp, err := http.Post(url, "application/json;charset=utf-8", body)
	if err != nil {
		log.Error("", err)
	}

	result, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Error("", err)
		return
	}
	log.Info(string(result))
}

//Update 更新成员，如果非必须的字段未指定，则不更新该字段之前的设置值
func (ctx *MemberCtx) Update() {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/update/update?access_token=%s", ctx.GetToken())
	dept, err := json.Marshal(ctx)
	if err != nil {
		log.Error("json err:", err)
	}
	body := bytes.NewBuffer([]byte(dept))
	resp, err := http.Post(url, "application/json;charset=utf-8", body)
	if err != nil {
		log.Error("", err)
	}

	result, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Error("", err)
		return
	}
	log.Info(string(result))

}

// Delete 删除成员
func (ctx *MemberCtx) Delete(id int) (*gohttp.Response, error) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/user/delete?access_token=%s&userid=%d", ctx.GetToken(), id)
	resp, err := gohttp.NewClient().URLString(url).Get()
	return resp, err
}

//批量删除成员

// GetUser 获取成员
func (ctx *MemberCtx) GetUser(userid string) (*gohttp.Response, error) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=%s&userid=%s", ctx.GetToken(), userid)
	resp, err := gohttp.NewClient().URLString(url).Get()
	return resp, err
}

// GetSimpleUserListByDepartment 获取部门成员
// 参数            必须    说明
// access_token    是      调用接口凭证
// department_id   是      获取的部门id
// fetch_child     否      1/0：是否递归获取子部门下面的成员
// status          否      0获取全部成员，1获取已关注成员列表，2获取禁用成员列表，4获取未关注成员列表。status可叠加，未填写则默认为4
func (ctx *MemberCtx) GetSimpleUserListByDepartment(departmentID int) (*gohttp.Response, error) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/user/simplelist?access_token=%s&department_id=%d&fetch_child=%d&status=%d",
		ctx.GetToken(), departmentID, 1, 0)
	resp, err := gohttp.NewClient().URLString(url).Get()
	return resp, err
}

// GetUserListByDepartment 获取部门成员(详情)
func (ctx *MemberCtx) GetUserListByDepartment(departmentID int) (*gohttp.Response, error) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/user/list?access_token=%s&department_id=%d&fetch_child=%d&status=%d",
		ctx.GetToken(), departmentID, 1, 0)
	resp, err := gohttp.NewClient().URLString(url).Get()
	return resp, err
}
