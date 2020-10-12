package routers

import (
	"DataCertProject/controllers"
	"github.com/astaxie/beego"
)
//router：作用接受并并发接收到的浏览器的请求
func init() {
    beego.Router("/", &controllers.MainController{})
//    用户注册的接口请求
beego.Router("/user_register",&controllers.RegisterController{})
//访问登录页面的接口
beego.Router("/login.html",&controllers.LoginController{})
//用户登录接口
beego.Router("/user_login",&controllers.LoginController{})
}
