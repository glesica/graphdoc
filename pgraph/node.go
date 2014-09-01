package pgraph

import (
    "fmt"
    "github.com/glesica/graphdoc/parsers"
)

// Node represents a given node type in a graph data model.
type Node struct {
    Label string `json:"id"`
    Desc string `json:"description"`
    Props map[string]*Prop `json:"properties"`
    Rels []*Rel `json:"edges"`
}

// NewNode returns a new node with the Props field initialized to an empty map.
func NewNode() *Node {
    n := new(Node)
    n.Props = make(map[string]*Prop)
    return n
}

// AppendDesc appends a string to the node description.
func (self *Node) AppendDesc(desc string) {
    self.Desc += fmt.Sprint("\n", desc)
}

// InsertRel adds a new relationship type pointing away from the current node.
func (self *Node) InsertRel(rel *Rel) {
    self.Rels = append(self.Rels, rel)
}

// InsertProp adds a new property to the node.
func (self *Node) InsertProp(prop *Prop) {
    self.Props[prop.Name] = prop
}

// ToMarkdown returns a representation of the node in Markdown format.
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

// ToHTML returns an HTML representation of the node, with relationship and
// property descriptions (and its own description) parsed using the provided
// parser.
func (self Node) ToHTML(parser parsers.Parser) string {
    propsHTML := ""
    for _, prop := range self.Props {
        propsHTML += prop.ToHTML(parser)
    }
    relsHTML := ""
    for _, rel := range self.Rels {
        relsHTML += rel.ToHTML(parser)
    }
    return fmt.Sprintf(nodeHTMLTemplate, self.Label, parser(self.Desc), propsHTML, relsHTML)
}

func (self Node) ToJS() string {
    relsJS := ""
    for _, rel := range self.Rels {
        relsJS += rel.ToJS()
    }
    return relsJS
}
