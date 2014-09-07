package schema

import (
    "regexp"
)

type ModelLayer int
const (
    GRAPH ModelLayer = iota
    NODE
    REL
    REL_FROM
    REL_TO
    PROP
    PROP_TYPE
    PROP_IND
    PROP_REQ
    PROP_UNIQ
    NONE
)

type ModelElement map[string]string


const graphExpStr = `@Graph +([A-Za-z0-9]+)`
var graphExp = regexp.MustCompile(graphExpStr)

const nodeExpStr = `@Node +([A-Z][A-Za-z0-9]*)`
var nodeExp = regexp.MustCompile(nodeExpStr)

const relExpStr = `@Rel +([A-Z][A-Z_]*)`
var relExp = regexp.MustCompile(relExpStr)

const relFromExpStr = `@From +([A-Z][A-Za-z0-9]*)`
var relFromExp = regexp.MustCompile(relFromExpStr)

const relToExpStr = `@To +([A-Z][A-Za-z0-9]*)`
var relToExp = regexp.MustCompile(relToExpStr)

const propExpStr = `@Prop +([a-z][A-Za-z0-9_]*)`
var propExp = regexp.MustCompile(propExpStr)

const propTypeExpStr = `@Type +(num|str|bool|date|time|timestamp|any)`
var propTypeExp = regexp.MustCompile(propTypeExpStr)

const propIndExpStr = `@Index`
var propIndExp = regexp.MustCompile(propIndExpStr)

const propReqExpStr = `@Required`
var propReqExp = regexp.MustCompile(propReqExpStr)

const propUniqExpStr = `@Unique`
var propUniqExp = regexp.MustCompile(propUniqExpStr)


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
        }
    }

    matches = relFromExp.FindStringSubmatch(line)
    if matches != nil {
        return REL_FROM, ModelElement{
            "label": matches[1],
        }
    }

    matches = relToExp.FindStringSubmatch(line)
    if matches != nil {
        return REL_TO, ModelElement{
            "label": matches[1],
        }
    }

    matches = propExp.FindStringSubmatch(line)
    if matches != nil {
        return PROP, ModelElement{
            "name": matches[1],
        }
    }

    matches = propTypeExp.FindStringSubmatch(line)
    if matches != nil {
        return PROP_TYPE, ModelElement{
            "type": matches[1],
        }
    }

    matches = propIndExp.FindStringSubmatch(line)
    if matches != nil {
        return PROP_IND, nil
    }

    matches = propReqExp.FindStringSubmatch(line)
    if matches != nil {
        return PROP_REQ, nil
    }

    matches = propUniqExp.FindStringSubmatch(line)
    if matches != nil {
        return PROP_UNIQ, nil
    }

    return NONE, nil
}
