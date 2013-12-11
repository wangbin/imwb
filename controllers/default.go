package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	r "github.com/christopherhesse/rethinkgo"
	"github.com/wangbin/imwb/models/auth"
)

const (
	SessionKey = "_auth_user_id"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}

type LoginForm struct {
	Id       int    `form:"-"`
	Name     string `form:"username"`
	Password string `form:"password,password,"`
}

type LoginController struct {
	beego.Controller
	userCache *auth.User
	rs        *r.Session
}

func (this *LoginController) Prepare() {
	rs = 
	userId := this.GetSession(SessionKey)
	if userId == nil {
		this.userCache = auth.NewAnonymousUser()
	} else {
		this.userCache = auth.GetUser(userId)
	}
}

func (this *LoginController) Get() {
	this.TplNames = "login.tpl"
	this.Data["Form"] = &LoginForm{}
	ss := this.StartSession()
	defer ss.SessionRelease()
	ss.Set("name", "wangbin")
}

func (this *LoginController) Post() {
	form := LoginForm{}
	if err := this.ParseForm(&form); err != nil {
		fmt.Println(err)
	}
	this.SetSession("auth_id", form.Password)
	this.SetSession("user", form.Password)
	fmt.Println(form)
	this.Ctx.Redirect(302, "/login/")
}
