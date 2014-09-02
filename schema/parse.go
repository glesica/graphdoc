package schema

import (
    "regexp"
)

type ModelLayer int
const (
    GRAPH ModelLayer = iota
    NODE
    PROP
    REL
    NONE
)

type ModelElement map[string]string

const graphExpStr = `:: ?Graph ([A-Za-z0-9]+) ?::`
var graphExp = regexp.MustCompile(graphExpStr)

const nodeExpStr = `:: ?Node ([A-Z][a-z0-9]*) ?::`
var nodeExp = regexp.MustCompile(nodeExpStr)

const relExpStr = `:: ?Rel ([_A-Z]+) ?-> ?([A-Z][a-z0-9]*) ?::`
var relExp = regexp.MustCompile(relExpStr)

const propExpStr = `:: ?Prop ([_A-Za-z0-9]+)(?: ?: ?(num|str|any|bool))? ?::`
var propExp = regexp.MustCompile(propExpStr)

func ParseLine(line string) (ModelLayer, ModelElement) {
    var matches []string

    matches = graphExp.FindStringSubmatch(line)
    if matches != nil {
        return GRAPH, ModelElement{
            "title": matches[1],
        }
    }

    matches = nodeExp.FindStringSubmatch(line)
    if matches != nil {
        return NODE, ModelElement{
            "label": matches[1],
        }
    }

    matches = relExp.FindStringSubmatch(line)
    if matches != nil {
        return REL, ModelElement{
            "label": matches[1],
            "target": matches[2],
        }
    }

    matches = propExp.FindStringSubmatch(line)
    if matches != nil {
        return PROP, ModelElement{
            "name": matches[1],
            "type": matches[2],
        }
    }

    return NONE, nil
}
