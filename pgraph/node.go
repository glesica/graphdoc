package pgraph

import (
    "fmt"
)

type Node struct {
    Label string
    Desc string
    Props map[string]*Prop
    Rels []*Rel
}

func NewNode() *Node {
    n := new(Node)
    n.Props = make(map[string]*Prop)
    return n
}

func (self *Node) AppendDesc(desc string) {
    self.Desc += fmt.Sprint("\n", desc)
}

func (self *Node) InsertRel(rel *Rel) {
    self.Rels = append(self.Rels, rel)
}

func (self *Node) InsertProp(prop *Prop) {
    self.Props[prop.Name] = prop
}

func (self Node) ToMarkdown() string {
    out := fmt.Sprintln("## ", self.Label)
    out += fmt.Sprintln(self.Desc)
    out += fmt.Sprintln("### Properties")
    for _, prop := range self.Props {
        out += prop.ToMarkdown()
    }
    out += fmt.Sprintln("### Relationships")
    for _, rel := range self.Rels {
        out += rel.ToMarkdown()
    }
    return out
}

const nodeHTMLTemplate = `
<div class="graphdoc-node">
    <h2 class="graphdoc-node-label">%s</h2>
    <div class="graphdoc-node-desc">%s</div>
    <div class="graphdoc-node-props">
        <h3>Properties</h3>
        %s
    </div>
    <div class="graphdoc-node-rels">
        <h3>Relationships</h3>
        %s
    </div>
</div>
`

func (self Node) ToHTML() string {
    propsHTML := ""
    for _, prop := range self.Props {
        propsHTML += prop.ToHTML()
    }
    relsHTML := ""
    for _, rel := range self.Rels {
        relsHTML += rel.ToHTML()
    }
    return fmt.Sprintf(nodeHTMLTemplate, self.Label, self.Desc, propsHTML, relsHTML)
}
