package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//默认显示的是注册页面
	c.TplName = "register.html"
}
