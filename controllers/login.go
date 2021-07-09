package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

type RegisterController struct {
	beego.Controller
}

// 登录页面 get
func (this *RegisterController) ShowLogin() {
	this.TplName = "login.html"
}

// 登录页面 post
func (this *RegisterController) HandleLogin() {
	// 拿到浏览器数据，并去除两边的空格
	Name := strings.TrimSpace(this.GetString("userName"))
	Pwd := strings.TrimSpace(this.GetString("passWord"))
	beego.Info("账号:", Name, "密码:", Pwd)

	//数据处理
	if Name == "" || Pwd == "" {
		beego.Info("登录失败！！")
		this.TplName = "login.html"
		this.Data["errmsg"] = "登录失败！！！！！"
		return
	}
	if Name== "zengyang" && Pwd=="zeng0428"{
	
		// 测试返回视图
		//this.Ctx.WriteString("登录成功")
		// 实际情况注册成功返回到登录页面
		this.Redirect("/login/blogs", 302)
	}else{
		beego.Info("登录失败！！")
		this.TplName = "login.html"
		this.Data["errmsg"] = "账户或密码错误"
		return
	}
}

