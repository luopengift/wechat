package wechat

import (
	//"net/http"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/luopengift/log"
)

// Department dept
type Department struct {
	Name     string `json:"name"`     //部门名称。长度限制为32个字（汉字或英文字母），字符不能包括\:*?"<>｜
	ParentID int    `json:"parentid"` //父亲部门id。根部门id为1
	Order    int    `json:"Order"`    //在父部门中的次序值。order值小的排序靠前。
	ID       int    `json:"id"`       //部门id，整型。指定时必须大于1，不指定时则自动生成

}

// DeptCtx dept ctx
type DeptCtx struct {
	*Context `json:"-"`
	*Department
}

//Create 创建部门
func (ctx *DeptCtx) Create() {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/department/create?access_token=%s", ctx.GetToken())
	dept, err := json.Marshal(ctx)
	if err != nil {
		log.Error("json err:", err)
	}
	body := bytes.NewBuffer([]byte(dept))
	resp, err := http.Post(url, "application/json;charset=utf-8", body)
	if err != nil {
		log.Error("error", err)
	}

	result, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Error("error", err)
		return
	}
	log.Info(string(result))

}

//Update 更新部门信息
func (ctx *DeptCtx) Update() {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/department/update?access_token=%s", ctx.GetToken())
	dept, err := json.Marshal(ctx)
	if err != nil {
		log.Error("json err:", err)
	}
	body := bytes.NewBuffer([]byte(dept))
	resp, err := http.Post(url, "application/json;charset=utf-8", body)
	if err != nil {
		log.Error("error", err)
		return
	}

	result, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Error("error", err)
		return
	}
	log.Info(string(result))

}

//Delete 删除部门;管理组须拥有指定部门的管理权限。;不能删除根部门；不能删除含有子部门、成员的部门
func (ctx *DeptCtx) Delete(id int) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/department/delete?access_token=%s&id=%d", ctx.GetToken(), id)
	resp, err := http.Get(url)
	if err != nil {
		log.Error("error", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Info(string(body))
}

// GetList 获取部门列表
func (ctx *DeptCtx) GetList(id int) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=%s&id=%d", ctx.GetToken(), id)
	log.Info(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Error("error", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Info("部门列表: %s", string(body))
}
