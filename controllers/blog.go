package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	r "github.com/christopherhesse/rethinkgo"
	"github.com/wangbin/imwb/models/blog"
	"github.com/wangbin/imwb/settings"
	"time"
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
	query := r.Table(blog.PostTable)
	page, err := this.GetInt("page")
	if err != nil || page < 0 {
		page = 0
	}
	tag := this.Ctx.Input.Param(":tag")
	year := this.Ctx.Input.Param(":year")
	if len(tag) > 0 {
		query = query.Filter(r.Row.Attr("tags").Contains(tag))
	}

	if len(year) > 0 {
		t, _ := time.Parse("2006-01-02T15:04:05Z", "2009-09-02T23:34:00Z")
		fmt.Println(t)
		query = query.Filter(r.Row.Attr("created").TypeOf().Eq("STRING"))
	}

	query.OrderBy(r.Desc("created")).Skip(
		page * NumberOfPostsPerPage).Limit(NumberOfPostsPerPage).Run(this.rs).All(
		&posts)

	this.Data["Posts"] = posts
}
