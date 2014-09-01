package parsers

import (
    "github.com/russross/blackfriday"
)

func Markdown(md string) string {
    html := blackfriday.MarkdownCommon([]byte(md))
    return string(html)
}
