package parsers

import (
    "fmt"
    "strings"
)

func Text(t string) string {
    html := strings.Join(strings.Split(t, "\n\n"), "</p><p>")
    return fmt.Sprintf("<p>%s</p>", html)
}
