package controllers

import (
	"github.com/astaxie/beego"
	"github.com/wangbin/imwb/models"
	"strconv"
)

const (
	NumberOfPostsPerPage = 10
)

func init() {
	beego.AddFuncMap("markup", Markup)
}

type PostListController struct {
	beego.Controller
}

func (this *PostListController) Get() {
	this.TplNames = "post-list.html"
	page, err := this.GetInt("page")
	if err != nil || page < 0 {
		page = 0
	}
	tag := this.Ctx.Input.Param(":tag")
	year, err := strconv.ParseInt(this.Ctx.Input.Param(":year"), 0, 64)
	if len(tag) > 0 {
		this.Data["Posts"] = models.PostsByTag(tag, NumberOfPostsPerPage,
			page*NumberOfPostsPerPage)
	} else if year > 0 {
		this.Data["Posts"] = models.PostsByYear(year, NumberOfPostsPerPage,
			page*NumberOfPostsPerPage)
	} else {
		this.Data["Posts"] = models.Posts(NumberOfPostsPerPage,
			page*NumberOfPostsPerPage)
	}
}
