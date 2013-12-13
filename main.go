package main

import (
	"github.com/astaxie/beego"
	"github.com/wangbin/imwb/controllers"
)

func main() {
	beego.SessionOn = true
	beego.Router("/login/", &controllers.LoginController{})
	beego.Router("/logout/", &controllers.LogoutController{})
	beego.Router("/blog/", &controllers.PostListController{})
	beego.Router(`/blog/tag/:tag([\w-]+)/`, &controllers.PostListController{})
	beego.Router(`/blog/:year(\d{4})/`, &controllers.PostListController{})
	beego.Run()
}
