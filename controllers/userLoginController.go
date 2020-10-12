package controllers

import (
	"DataCertProject/models"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}
//处理直接访问login.html页面
func (l *LoginController) Get(){
//	设置login.html为模板文件
//	template 模板
	l.TplName="login.html"
}
//用户登录接口
func (l *LoginController)Post(){
	//解析登录的数据
	var user models.User
	err:=l.ParseForm(&user)
	if err!=nil{
		l.Ctx.WriteString("抱歉，登录页面数据解析失败，请重试")
		return
	}
//查询数据库的用户信息
_,err=user.QueryUser()
if err !=nil{
	l.Ctx.WriteString("抱歉，用户登录失败，请重试")
	return
}
l.TplName="upload.html"
}