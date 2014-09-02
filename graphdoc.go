package main

import (
    "bufio"
    //"encoding/json"
    "flag"
    //"fmt"
    "log"
    "os"
    //"io/ioutil"
    "github.com/glesica/graphdoc/formatters"
    "github.com/glesica/graphdoc/pgraph"
    "github.com/glesica/graphdoc/parsers"
    "github.com/glesica/graphdoc/schema"
)

func main() {
    //var outFormat = flag.String("outformat", "html", "Output format")
    outPath := flag.String("outpath", "", "Output path, stdout will be used if omitted.")
    inFormat := flag.String("informat", "", "Input file format, default is to auto-detect.")
    flag.Parse()
    inPath := flag.Arg(0)

    inputFile, err := os.Open(inPath)
    if err != nil {
        log.Fatal("Error opening file:", err)
    }
    defer inputFile.Close()

    var outputFile *os.File
    if *outPath == "" {
        outputFile = os.Stdout
    } else {
        outputFile, err = os.Create(*outPath)
        if err != nil {
            log.Fatal("Error opening file:", err)
        }
        defer outputFile.Close()
    }

    scanner := bufio.NewScanner(inputFile)


    graph := pgraph.Graph{}
    var currentNode *pgraph.Node
    var currentRel *pgraph.Rel
    var currentProp *pgraph.Prop

    var depth schema.ModelLayer
    var layer schema.ModelLayer
    var element schema.ModelElement

    for scanner.Scan() {
        line := scanner.Text()
        layer, element = schema.ParseLine(line)

        switch layer {
        case schema.GRAPH:
            graph.Title = element["title"]
        case schema.NODE:
            currentNode = pgraph.NewNode()
            currentNode.Label = element["label"]
            graph.InsertNode(currentNode)
        case schema.REL:
            currentRel = pgraph.NewRel()
            currentRel.Label = element["label"]
            currentRel.Target = element["target"]
            currentRel.Source = currentNode.Label
            currentNode.InsertRel(currentRel)
        case schema.PROP:
            currentProp = pgraph.NewProp()
            currentProp.Name = element["name"]
            currentProp.DataType = element["type"]
            currentNode.InsertProp(currentProp)
        case schema.NONE:
            switch depth {
            case schema.GRAPH:
                graph.AppendDesc(line)
            case schema.NODE:
                currentNode.AppendDesc(line)
            case schema.REL:
                currentRel.AppendDesc(line)
            case schema.PROP:
                currentProp.AppendDesc(line)
            }
        }

        if layer != schema.NONE {
            depth = layer
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal("Error reading file:", scanner.Err())
    }

    var parser parsers.Parser
    switch *inFormat {
    case "txt":
        parser = parsers.Text
    case "md":
        parser = parsers.Markdown
    case "":
        parser = parsers.Text // For now we default to text, later we will auto-detect
    }

    docString := formatters.HTMLDocument(graph, parser)

    outputFile.WriteString(docString)
}
