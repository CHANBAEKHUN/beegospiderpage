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
}
