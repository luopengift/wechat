package wechat

import (
	"github.com/luopengift/log"
	"testing"
)


var wechat *WeChatCtx

func Test_wechat(t *testing.T) {
    corpID := ""
    secret := ""
	log.Info("企业号测试...")
	wechat = NewWeChatCtx(corpID, secret)
	wechat.SetType(QY)
	wechat.GetServerList()
	log.Info("企业号测试成功")
}

/*
func Test_mp(t *testing.T) {
    log.Info("公众号测试...")
    wechat := NewWeChatCtx(id,secret)
    wechat.SetType(MP)
    wechat.GetServerList()
    log.Info("公众号测试成功")
}
*/

func Test_qy_dept(t *testing.T) {
	log.Info("企业号用户测试")
	Dept := &DeptCtx{wechat, nil}
	Dept.GetList(1)
	//mem := MemberCtx{wechat,nil}
	//mem.GetUserByDepartment(1)

}

func Test_qy_member(t *testing.T) {
	log.Info("企业号用户测试")
	member := &MemberCtx{wechat, nil}
	member.GetUser("peng.luo@inveno.cn")
	//member.GetUserListByDepartment(1)
}
