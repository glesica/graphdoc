package pgraph

import (
    "fmt"
)

type Prop struct {
    Name string
    Desc string
    DataType string
}

func NewProp() *Prop {
    p := new(Prop)
    return p
}

func (self *Prop) AppendDesc(desc string) {
    self.Desc += fmt.Sprint("\n", desc)
}

func (p Prop) ToString() string {
    return fmt.Sprint(p.Name, " (", p.DataType, ")")
}

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

func (self Prop) ToHTML() string {
    return fmt.Sprintf(propHTMLTemplate, self.Name, self.DataType, self.Desc)
}
