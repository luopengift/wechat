package wechat

import (
	//"net/http"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/luopengift/log"
	"io/ioutil"
	"net/http"
)

type Department struct {
	Name     string `json:"name"`     //部门名称。长度限制为32个字（汉字或英文字母），字符不能包括\:*?"<>｜
	ParentId int    `json:"parentid"` //父亲部门id。根部门id为1
	Order    int    `json:"Order"`    //在父部门中的次序值。order值小的排序靠前。
	Id       int    `json:"id"`       //部门id，整型。指定时必须大于1，不指定时则自动生成

}

type DeptCtx struct {
	*WeChatCtx `json:"-"`
	*Department
}

//创建部门
func (self *DeptCtx) Create() {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/department/create?access_token=%s", self.GetToken())
	dept, err := json.Marshal(self)
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

//更新部门信息
func (self *DeptCtx) Update() {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/department/update?access_token=%s", self.GetToken())
	dept, err := json.Marshal(self)
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

//删除部门;管理组须拥有指定部门的管理权限。;不能删除根部门；不能删除含有子部门、成员的部门
func (self *DeptCtx) Delete(id int) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/department/delete?access_token=%s&id=%d", self.GetToken(), id)
	resp, err := http.Get(url)
	if err != nil {
		log.Error("error", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Info(string(body))
}

//获取部门列表
func (self *DeptCtx) GetList(id int) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=%s&id=%d", self.GetToken(), id)
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
