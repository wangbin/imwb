package main

import (
	"github.com/astaxie/beego"
	"github.com/wangbin/imwb/controllers"
)

func main() {
	beego.SessionOn = true
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login/", &controllers.LoginController{})
	beego.Run()
}
