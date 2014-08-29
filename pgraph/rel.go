package pgraph

import (
    "fmt"
)

type Rel struct {
    Label string
    Desc string
    Source string
    Target string
}

func NewRel() *Rel {
    r := new(Rel)
    return r
}

func (self *Rel) AppendDesc(desc string) {
    self.Desc += fmt.Sprint("\n", desc)
}

func (r Rel) ToString() string {
    return fmt.Sprint(r.Label, " -> ", r.Target)
}

func (self Rel) ToDOT() string {
    out := fmt.Sprintf("  %s -> %s [label=%s];\n", self.Source, self.Target, self.Label)
    return out
}

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

func (self Rel) ToHTML() string {
    return fmt.Sprintf(relHTMLTemplate, self.Label, self.Source, self.Target, self.Desc)
}
