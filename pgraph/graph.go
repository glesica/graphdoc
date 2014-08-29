package pgraph

import (
    "fmt"
)

type Graph struct {
    Title string
    Desc string
    Nodes map[string]*Node
}

func NewGraph() Graph {
    g := Graph{}
    g.Nodes = make(map[string]*Node)
    return g
}

func (self *Graph) AppendDesc(desc string) {
    self.Desc += fmt.Sprint("\n", desc)
}

func (self *Graph) InsertNode(node *Node) {
    self.Nodes[node.Label] = node
}

func (self Graph) ToDOT() string {
    out := fmt.Sprintf("digraph %s {\n", self.Title)
    for _, node := range self.Nodes {
        for _, rel := range node.Rels {
            out += rel.ToDOT()
        }
    }
    out += fmt.Sprint("}")
    return out
}

func (self Graph) ToMarkdown() string {
    out := fmt.Sprintln("# ", self.Title)
    out += fmt.Sprintln(self.Desc)
    for _, node := range self.Nodes {
        out += node.ToMarkdown()
    }
    return out
}

const graphHTMLTemplate = `
<div class="graphdoc-graph">
    <h1 class="graphdoc-graph-title">%s</h1>
    <div class="graphdoc-graph-desc">%s</div>
    <div class="graphdoc-graph-nodes">%s</div>
</div>
`

func (self Graph) ToHTML() string {
    nodesHTML := ""
    for _, node := range self.Nodes {
        nodesHTML += node.ToHTML()
    }
    return fmt.Sprintf(graphHTMLTemplate, self.Title, self.Desc, nodesHTML)
}
