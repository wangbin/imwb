package controllers

import (
	"github.com/astaxie/beego"
	r "github.com/christopherhesse/rethinkgo"
	"github.com/wangbin/imwb/models/blog"
	"github.com/wangbin/imwb/settings"
)

const (
	NumberOfPostsPerPage = 10
)

func init() {
	beego.AddFuncMap("markup", Markup)
}

type PostListController struct {
	beego.Controller
	rs *r.Session
}

func (this *PostListController) Prepare() {
	this.rs, _ = r.Connect(settings.DbUri, settings.DbName)
}

func (this *PostListController) Finish() {
	this.rs.Close()
}

func (this *PostListController) Get() {
	this.TplNames = "post-list.html"
	var posts []*blog.Post
	page, err := this.GetInt("page")
	if err != nil || page < 0 {
		page = 0
	}
	r.Table(blog.PostTable).OrderBy(r.Desc("created")).Skip(page * NumberOfPostsPerPage).Limit(
		NumberOfPostsPerPage).Run(this.rs).All(&posts)
	this.Data["Posts"] = posts
}
