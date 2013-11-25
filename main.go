package main

import (
	"github.com/wangbin/imwb/controllers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SessionOn = true
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login/", &controllers.LoginController{})
	beego.Run()
}

