package main

import (
    "bufio"
    //"encoding/json"
    //"flag"
    "fmt"
    "log"
    "os"
    //"io/ioutil"
    "regexp"
    "github.com/glesica/graphdoc/pgraph"
    "github.com/glesica/graphdoc/graphdoc"
)

var graphExp = regexp.MustCompile(":: ?Graph ([A-Za-z0-9]+) ?::")
var nodeExp = regexp.MustCompile(":: ?Node ([A-Z][a-z0-9]*) ?::")
var relExp = regexp.MustCompile(":: ?Rel ([A-Z_]+) ?-> ?([A-Z][a-z0-9]*) ?::")
var propExp = regexp.MustCompile(":: ?Prop ([A-Za-z0-9]+) ?: ?(num|str|any) ?::")

func main() {
    //var outFormat = flag.String("outformat", "html", "Output format")
    //flag.Parse()

    //fmt.Println(*outFormat)

    inputFile, err := os.Open("testDoc0.md")
    if err != nil {
        log.Fatal("Error opening file:", err)
    }
    defer inputFile.Close()

    outputFile, err := os.Create("testDoc0_filtered.md")
    if err != nil {
        log.Fatal("Error opening file:", err)
    }
    defer outputFile.Close()

    scanner := bufio.NewScanner(inputFile)

    type Depth int
    const (
        GRAPH Depth = iota
        NODE
        PROP
        REL
    )

    graph := pgraph.NewGraph()
    var currentNode *pgraph.Node
    var currentRel *pgraph.Rel
    var currentProp *pgraph.Prop
    var matches []string
    depth := GRAPH
    for scanner.Scan() {
        line := scanner.Text()

        matches = graphExp.FindStringSubmatch(line)
        if matches != nil {
            // TODO Support more than one graph
            graph.Title = matches[1]
            depth = GRAPH
            continue
        }

        matches = nodeExp.FindStringSubmatch(line)
        if matches != nil {
            currentNode = pgraph.NewNode()
            currentNode.Label = matches[1]
            graph.InsertNode(currentNode)
            depth = NODE
            continue
        }

        matches = relExp.FindStringSubmatch(line)
        if matches != nil {
            currentRel = pgraph.NewRel()
            currentRel.Label = matches[1]
            currentRel.Target = matches[2]
            currentRel.Source = currentNode.Label
            currentNode.InsertRel(currentRel)
            depth = REL
            continue
        }

        matches = propExp.FindStringSubmatch(line)
        if matches != nil {
            currentProp = pgraph.NewProp()
            currentProp.Name = matches[1]
            currentProp.DataType = matches[2]
            currentNode.InsertProp(currentProp)
            depth = PROP
            continue
        }

        switch depth {
        case GRAPH:
            graph.AppendDesc(line)
        case NODE:
            currentNode.AppendDesc(line)
        case REL:
            currentRel.AppendDesc(line)
        case PROP:
            currentProp.AppendDesc(line)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal("Error reading file:", scanner.Err())
    }

    fmt.Print(graphdoc.HTMLDocument(graph))
}
