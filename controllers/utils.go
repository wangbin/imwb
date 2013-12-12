package controllers

import (
	"fmt"
	"github.com/russross/blackfriday"
	"html/template"
	"regexp"
)

var (
	titlePattern = regexp.MustCompile(`#{2,}`)
)

func convertTitle(in string) string {
	return fmt.Sprintf("%s#", in)
}

func Markup(in string, needConvert bool) template.HTML {
	if needConvert {
		converted := titlePattern.ReplaceAllStringFunc(in, convertTitle)
		return template.HTML(blackfriday.MarkdownCommon([]byte(converted)))
	} else {
		return template.HTML(blackfriday.MarkdownCommon([]byte(in)))
	}
}
