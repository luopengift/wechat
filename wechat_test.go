package wechat

import (
    "testing"
    "github.com/luopengift/golibs/logger"
)


var wechat *WeChatCtx

func Test_wechat(t *testing.T) {
    logger.Info("企业号测试...")
    wechat = NewWeChatCtx(id,secret)
    wechat.SetType(QY)
    wechat.GetServerList()
    logger.Info("企业号测试成功")

}


func Test_mp(t *testing.T) {
    logger.Info("公众号测试...")
    wechat := NewWeChatCtx(id,secret)
    wechat.SetType(MP)
    wechat.GetServerList()
    logger.Info("公众号测试成功")
}


func Test_qy_member(t *testing.T) {
    logger.Info("企业号用户测试")
    Dept := &DeptCtx{wechat, nil}
    Dept.GetList(2)
    //mem := MemberCtx{wechat,nil}
    //mem.GetUserByDepartment(1)

}
