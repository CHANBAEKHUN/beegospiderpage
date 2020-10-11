package controllers

import (
    "DataCertProject/models"
    "github.com/astaxie/beego"
)

type RegisterController struct{
    beego.Controller
}
func (r *RegisterController) Post(){
//1:解析请求数据
    var user models.User
   err:= r.ParseForm(&user)
   if err !=nil{
       r.Ctx.WriteString("抱歉，解析错误，请重试")
       return
   }
//	2：保存用户信息到数据库

//	3：返回前端结果(成功登陆，失败就弹出错误）

}