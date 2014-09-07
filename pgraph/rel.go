package pgraph

import (
    "fmt"
)

// Rel represents a particular type of relationship that can be present in a
// graph data model.
type Rel struct {
    Label string `json:"label"`
    Desc string `json:"desc"`
    Source string `json:"src"`
    Target string `json:"targ"`
    Props map[string]*Prop `json:"props"`
}

func NewRel() *Rel {
    r := new(Rel)
    r.Props = make(map[string]*Prop)
    return r
}

// TODO Add constructor that can parse a rel from a string or return nil (or
// something, maybe err) if it doesn't match the proper format

// AppendDesc appends a string to the relationship description.
func (r *Rel) AppendDesc(desc string) {
    r.Desc += fmt.Sprint("\n", desc)
}

// InsertProp adds a property to the relationship.
func (r *Rel) InsertProp(prop *Prop) {
    r.Props[prop.Name] = prop
}

// Return the relationship formatted as a Cypher relationship as it might be in
// a MATCH query.
// TODO Make variable names reflect labels
func (r Rel) Cypher() string {
    return fmt.Sprintf("(s:%s)-[r:%s]->(t:%s)", r.Source, r.Label, r.Target)
}

func (r Rel) String() string {
    return r.Cypher()
}
