package models

import (
	r "github.com/dancannon/gorethink"
	"time"
)

const (
	PostTable            = "blog_post"
	PostTableSecondIndex = "slug"
)

type Post struct {
	Id         string    `gorethink:"id,omitempty"`
	Title      string    `gorethink:"title"`
	Slug       string    `gorethink:"slug"`
	Created    time.Time `gorethink:"created"`
	Modified   time.Time `gorethink:"modified"`
	Content    string    `gorethink:"content"`
	RenderType string    `gorethink:"render_type"`
	Tags       []string  `gorethink:"tags"`
	User       string    `gorethink:"user_id"`
}

func paginate(query r.RqlTerm, limit, offset int64) (posts []*Post) {
	query = query.OrderBy(r.Desc("created"))
	if offset > 0 {
		query = query.Skip(offset)
	}
	if limit > 0 {
		query = query.Limit(limit)
	}
	rows, _ := query.Run(Conn)
	rows.ScanAll(&posts)
	return
}

func Posts(limit, offset int64) []*Post {
	query := r.Table(PostTable)
	return paginate(query, limit, offset)
}

func PostsByTag(tag string, limit, offset int64) []*Post {
	query := r.Table(PostTable)
	if len(tag) > 0 {
		query = query.Filter(r.Row.Field("tags").Contains(tag))
	}
	return paginate(query, limit, offset)
}

func PostsByYear(year, limit, offset int64) []*Post {
	query := r.Table(PostTable).Filter(r.Row.Field("created").Year().Eq(year))
	return paginate(query, limit, offset)
}
