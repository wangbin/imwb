package controllers

import (
	"github.com/astaxie/beego"
	r "github.com/christopherhesse/rethinkgo"
	"github.com/wangbin/imwb/forms"
	"github.com/wangbin/imwb/models"
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
	this.Data["User"] = this.userCache
}

func (this *LoginController) Post() {
	form := forms.NewLoginForm()
	err := this.ParseForm(form)
	if err != nil {
		form.SetNonFieldError(err)
	}
	form.SetRs(this.rs)
	if !form.IsValid() {
		this.TplNames = "login.tpl"
		this.Data["User"] = this.userCache
		this.Data["Form"] = form
	} else {
		this.login(form.User().Id)
		this.Ctx.Redirect(302, "/login/")
	}
}

func (this *LoginController) login(userId string) {
	this.DelSession(SessionKey)
	this.SetSession(SessionKey, userId)
}

type LogoutController struct {
	beego.Controller
}

func (this *LogoutController) Get() {
	this.DelSession(SessionKey)
	this.Ctx.Redirect(302, "/login/")
}
