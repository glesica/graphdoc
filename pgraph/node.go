package pgraph

import (
    "fmt"
)

// Node represents a given node type in a graph data model.
type Node struct {
    Label string `json:"id"`
    Desc string `json:"description"`
    Props map[string]*Prop `json:"properties"`
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

// InsertProp adds a new property to the node.
func (self *Node) InsertProp(prop *Prop) {
    self.Props[prop.Name] = prop
}
