package models

import (
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
