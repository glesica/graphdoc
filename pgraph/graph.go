package pgraph

import (
    "fmt"
    "github.com/glesica/graphdoc/parsers"
)

type Graph struct {
    Title string `json:"title"`
    Desc string `json:"description"`
    Nodes []*Node `json:"nodes"`
}

func (self *Graph) AppendDesc(desc string) {
    self.Desc += desc
}

func (self *Graph) InsertNode(node *Node) {
    self.Nodes = append(self.Nodes, node)
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
    <div id="graphdoc-graph-viz"></div>
    <div class="graphdoc-graph-nodes">%s</div>
</div>
`

func (self Graph) ToHTML(parser parsers.Parser) string {
    nodesHTML := ""
    for _, node := range self.Nodes {
        nodesHTML += node.ToHTML(parser)
    }
    return fmt.Sprintf(graphHTMLTemplate, self.Title, parser(self.Desc), nodesHTML)
}

const graphJSTemplate = `
var g = new Graph();
%s
`

func (self Graph) ToJS() string {
    nodesJS := ""
    for _, node := range self.Nodes {
        nodesJS += node.ToJS()
    }
    return fmt.Sprintf(graphJSTemplate, nodesJS)
}
