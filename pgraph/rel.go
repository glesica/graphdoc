package pgraph

import (
    "fmt"
    "github.com/glesica/graphdoc/parsers"
)

// Rel represents a particular type of relationship that can be present in a
// graph data model.
type Rel struct {
    Label string `json:"id"`
    Desc string `json:"description"`
    Source string `json:"source"`
    Target string `json:"target"`
}

func NewRel() *Rel {
    r := new(Rel)
    return r
}

// AppendDesc appends a string to the relationship description.
func (self *Rel) AppendDesc(desc string) {
    self.Desc += fmt.Sprint("\n", desc)
}

func (r Rel) ToString() string {
    return fmt.Sprint(r.Label, " -> ", r.Target)
}

// ToDOT returns the relationship in DOT format.
func (self Rel) ToDOT() string {
    var target string
    if self.Source == self.Target {
        target = fmt.Sprintf("%s%s", self.Target, self.Label)
    } else {
        target = self.Target
    }
    out := fmt.Sprintf("  %s -> %s [label=%s];\n", self.Source, target, self.Label)
    return out
}

// ToMarkdown returns the relationship in Markdown format.
func (self Rel) ToMarkdown() string {
    out := fmt.Sprintf("#### %s \u279E %s\n", self.Label, self.Target)
    out += fmt.Sprintln(self.Desc)
    return out
}

const relHTMLTemplate = `
<div class="graphdoc-rel">
    <h4 class="graphdoc-rel-label">
        %s
        <small class="graphdoc-rel-path">%s ➔ %s</small>
    </h4>
    <div class="graphdoc-prop-desc">%s</div>
</div>
`

// ToHTML returns the relationship in HTML format with its description
// formatted using the provided parser.
func (self Rel) ToHTML(parser parsers.Parser) string {
    return fmt.Sprintf(relHTMLTemplate, self.Label, self.Source, self.Target, parser(self.Desc))
}

const relJSTemplate = `
g.addEdge("%s", "%s");
`

func (self Rel) ToJS() string {
    return fmt.Sprintf(relJSTemplate, self.Source, self.Target)
}
