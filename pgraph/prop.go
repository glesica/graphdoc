package pgraph

import (
    "fmt"
)

// Prop represents a node property within a graph data model.
type Prop struct {
    Name string `json:"name"`
    Desc string `json:"description"`
    Type string `json:"type"`
    Ind bool `json:"ind"`
    Req bool `json:"req"`
    Uniq bool `json:"uniq"`
}

func NewProp() *Prop {
    p := new(Prop)
    return p
}

// AppendDesc appends the given string to the property description.
func (self *Prop) AppendDesc(desc string) {
    self.Desc += fmt.Sprint("\n", desc)
}

func (p Prop) String() string {
    return p.Name
}
