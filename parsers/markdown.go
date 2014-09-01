package parsers

import (
    "github.com/russross/blackfriday"
)

// Markdown processes a string of Markdown formatted text into a string of html
// formatted text. This is mostly just a shortcut for now, and no special logic
// is applied.
func Markdown(md string) string {
    html := blackfriday.MarkdownCommon([]byte(md))
    return string(html)
}
