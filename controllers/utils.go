package controllers

import (
	"bytes"
	"fmt"
	"github.com/wangbin/blackfriday"
	"html/template"
	"regexp"
	"strings"
)

var (
	titlePattern = regexp.MustCompile(`#{2,}`)
)

func convertTitle(in string) string {
	return fmt.Sprintf("%s#", in)
}

func doubleSpace(out *bytes.Buffer) {
	if out.Len() > 0 {
		out.WriteByte('\n')
	}
}

type Html struct {
	*blackfriday.Html
}

func (options *Html) BlockCode(out *bytes.Buffer, text []byte, lang string) {
	options.BlockCodeNormal(out, text, lang)
}

func (options *Html) BlockCodeNormal(out *bytes.Buffer, text []byte, lang string) {
	doubleSpace(out)
	// parse out the language names/classes
	count := 0
	for _, elt := range strings.Fields(lang) {
		if elt[0] == '.' {
			elt = elt[1:]
		}
		if len(elt) == 0 {
			continue
		}
		if count == 0 {
			out.WriteString("<pre><code data-language=\"")
		} else {
			out.WriteByte(' ')
		}
		attrEscape(out, []byte(elt))
		count++
	}

	if count == 0 {
		out.WriteString("<pre><code>")
	} else {
		out.WriteString("\">")
	}

	attrEscape(out, text)
	out.WriteString("</code></pre>\n")
}

func MarkdownCustom(input []byte) []byte {
	// set up the HTML renderer
	htmlFlags := 0
	htmlFlags |= blackfriday.HTML_USE_XHTML
	htmlFlags |= blackfriday.HTML_USE_SMARTYPANTS
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_FRACTIONS
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_LATEX_DASHES
	htmlFlags |= blackfriday.HTML_SKIP_SCRIPT
	renderer := &Html{blackfriday.HtmlRenderer(htmlFlags, "", "").(*blackfriday.Html)}

	// set up the parser
	extensions := 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS

	return blackfriday.Markdown(input, renderer, extensions)
}

func attrEscape(out *bytes.Buffer, src []byte) {
	org := 0
	for i, ch := range src {
		// using if statements is a bit faster than a switch statement.
		// as the compiler improves, this should be unnecessary
		// this is only worthwhile because attrEscape is the single
		// largest CPU user in normal use
		if ch == '"' {
			if i > org {
				// copy all the normal characters since the last escape
				out.Write(src[org:i])
			}
			org = i + 1
			out.WriteString("&quot;")
			continue
		}
		if ch == '&' {
			if i > org {
				out.Write(src[org:i])
			}
			org = i + 1
			out.WriteString("&amp;")
			continue
		}
		if ch == '<' {
			if i > org {
				out.Write(src[org:i])
			}
			org = i + 1
			out.WriteString("&lt;")
			continue
		}
		if ch == '>' {
			if i > org {
				out.Write(src[org:i])
			}
			org = i + 1
			out.WriteString("&gt;")
			continue
		}
	}
	if org < len(src) {
		out.Write(src[org:])
	}
}

func Markup(in string, needConvert bool) template.HTML {
	if needConvert {
		converted := titlePattern.ReplaceAllStringFunc(in, convertTitle)
		return template.HTML(MarkdownCustom([]byte(converted)))
	} else {
		return template.HTML(MarkdownCustom([]byte(in)))
	}
}
