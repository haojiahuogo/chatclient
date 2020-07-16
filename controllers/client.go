package controllers

import (
	ser "chatclient/service"
	"encoding/json"

	"github.com/astaxie/beego"
)

type ClientController struct {
	beego.Controller
}

//跳转到登录页面
func (this *ClientController) Get() {
	this.TplName = "login.html"
}

//登录到聊天主页
func (this *ClientController) Login() {
	userid := this.GetString("userid")
	if len(userid) == 0 {
		this.Redirect("/", 302)
		return
	}
	url := "http://127.0.0.1:8088/v1/msg/GetConversation"
	postArgs := map[string]string{
		"userid": userid,
	}
	headerArgs := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	res, err := ser.RequestPost("POST", url, postArgs, headerArgs)
	if err == nil {
		beego.Debug(string(res))
		var alluser Res
		err := json.Unmarshal(res, &alluser)
		if err != nil {
			beego.Debug(err)
		}
		this.Data["AllUsers"] = alluser.Data
	}
	this.TplName = "main.html"
	this.Data["userid"] = userid
	this.Data["username"] = userid

}

type Res struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type User struct {
	Cvsid    int
	Msgid    int
	Source   int
	Content  string
	Sendtime int
	Username string
}
