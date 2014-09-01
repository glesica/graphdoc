package parsers

import (
    "fmt"
    "strings"
)

// Text parses plain text into HTML by applying one simple rule, double line
// breaks are replaced with a set of closing and opening paragraph tags, then
// the entire string is surrounded by paragraph tags.
func Text(t string) string {
    html := strings.Join(strings.Split(t, "\n\n"), "</p><p>")
    return fmt.Sprintf("<p>%s</p>", html)
}
