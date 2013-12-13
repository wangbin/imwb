package controllers

import (
	"github.com/astaxie/beego"
	"github.com/wangbin/imwb/forms"
	"github.com/wangbin/imwb/models"
)

const (
	SessionKey = "_auth_user_id"
)

type LoginController struct {
	beego.Controller
	userCache *models.User
}

func (this *LoginController) Prepare() {
	userId := this.GetSession(SessionKey)
	if userId == nil {
		this.userCache = models.AnonymousUser()
	} else {
		this.userCache = models.GetUser(userId.(string))
	}
}

func (this *LoginController) Get() {
	this.TplNames = "login.tpl"
	form := forms.NewLoginForm(this.userCache)
	this.Data["Form"] = form
	this.Data["User"] = this.userCache
}

func (this *LoginController) Post() {
	form := forms.NewLoginForm(this.userCache)
	err := this.ParseForm(form)
	if err != nil {
		form.SetNonFieldError(err)
	}
	if !form.IsValid() {
		this.TplNames = "login.tpl"
		this.Data["User"] = this.userCache
		this.Data["Form"] = form
	} else {
		this.login(form.User())
		this.Ctx.Redirect(302, "/login/")
	}
}

func (this *LoginController) login(user *models.User) {
	this.userCache = user
	this.DelSession(SessionKey)
	this.SetSession(SessionKey, user.Id)
}

type LogoutController struct {
	beego.Controller
}

func (this *LogoutController) Get() {
	this.DelSession(SessionKey)
	this.Ctx.Redirect(302, "/login/")
}
