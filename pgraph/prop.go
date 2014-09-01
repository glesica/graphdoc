package pgraph

import (
    "fmt"
    "github.com/glesica/graphdoc/parsers"
)

// Prop represents a node property within a graph data model.
type Prop struct {
    Name string `json:"name"`
    Desc string `json:"description"`
    DataType string `json:"type"`
}

func NewProp() *Prop {
    p := new(Prop)
    return p
}

// AppendDesc appends the given string to the property description.
func (self *Prop) AppendDesc(desc string) {
    self.Desc += fmt.Sprint("\n", desc)
}

func (p Prop) ToString() string {
    return fmt.Sprint(p.Name, " (", p.DataType, ")")
}

// ToMarkdown returns the property formatted as Markdown.
func (self Prop) ToMarkdown() string {
    out := fmt.Sprintf("#### %s (%s)\n", self.Name, self.DataType)
    out += fmt.Sprintln(self.Desc)
    return out
}

const propHTMLTemplate = `
<div class="graphdoc-prop">
    <h4 class="graphdoc-prop-name">
        %s
        <small class="graphdoc-prop-type">%s</small>
    </h4>
    <div class="graphdoc-prop-desc">%s</div>
</div>
`

// ToHTML returns the property in HTML format with its description parsed using
// the provided parser.
func (self Prop) ToHTML(parser parsers.Parser) string {
    return fmt.Sprintf(propHTMLTemplate, self.Name, self.DataType, parser(self.Desc))
}
