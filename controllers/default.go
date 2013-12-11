package controllers

import (
	"github.com/astaxie/beego"
	r "github.com/christopherhesse/rethinkgo"
	"github.com/wangbin/imwb/forms"
	"github.com/wangbin/imwb/models/auth"
	"github.com/wangbin/imwb/settings"
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

type LoginController struct {
	beego.Controller
	userCache *auth.User
	rs        *r.Session
}

func (this *LoginController) Prepare() {
	this.rs, _ = r.Connect(settings.DbUri, settings.DbName)
	userId := this.GetSession(SessionKey)
	if userId == nil {
		this.userCache = auth.NewAnonymousUser()
	} else {
		this.userCache = auth.GetUser(this.rs, userId.(string))
	}
}

func (this *LoginController) Finish() {
	this.rs.Close()
}

func (this *LoginController) Get() {
	this.TplNames = "login.tpl"
	form := forms.NewLoginForm()
	this.Data["Form"] = form
}

func (this *LoginController) Post() {
	form := forms.NewLoginForm()
	err := this.ParseForm(form)
	if err != nil {
		form.SetNonFieldError(err)
	}
	if !form.IsValid() {
		this.TplNames = "login.tpl"
		this.Data["Form"] = form
	} else {
		this.Ctx.Redirect(302, "/login/")
	}
}
