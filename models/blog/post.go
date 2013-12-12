package blog

import (
	"time"
)

const (
	PostTable = "blog_post"
)

type Post struct {
	Id         string    `json:"id,omitempty"`
	Title      string    `json:"title"`
	Slug       string    `json:"slug"`
	Created    time.Time `json:"created"`
	Modified   time.Time `json:"modified"`
	Content    string    `json:"content"`
	RenderType string    `json:"render_type"`
	Tags       []string  `json:"tags"`
	User       string    `json:"user_id"`
}
