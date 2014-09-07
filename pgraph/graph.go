// The pgraph package provides tools for building a simulated property graph
// model that can be used to build documentation.
package pgraph

import (
    "fmt"
)

// Graph is a representation of a graph data model.
type Graph struct {
    Title string `json:"title"`
    Desc string `json:"desc"`
    Nodes map[string]*Node `json:"nodes"`
    Rels map[string]*Rel `json:"rels"`
}

func NewGraph() *Graph {
    g := new(Graph)
    g.Nodes = make(map[string]*Node)
    g.Rels = make(map[string]*Rel)
    return g
}

// AppendDesc appends a string to the description of the graph data model. No
// other formatting is done to the string and no newlines are added or removed.
func (g *Graph) AppendDesc(desc string) {
    g.Desc += desc
}

// InsertNode inserts a new node type into the graph data model.
func (g *Graph) InsertNode(node *Node) {
    g.Nodes[node.Label] = node
}

func (g *Graph) InsertRel(rel *Rel) {
    g.Rels[rel.Label] = rel
}

// DOT returns a representation of the graph data model in DOT format.
func (g Graph) DOT() string {
    out := fmt.Sprintf("digraph %s { ", g.Title)
    for _, rel := range g.Rels {
        out += fmt.Sprintf(" %s -> %s [label=%s];", rel.Source, rel.Target, rel.Label)
    }
    out += fmt.Sprint(" }")
    return out
}
