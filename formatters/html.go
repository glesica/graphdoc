// The formatters package provides tools to format documentation based on a
// property graph data model.
package formatters

import (
    "bytes"
    //"fmt"
    "strings"
    "text/template"
    "github.com/glesica/graphdoc/pgraph"
    "github.com/glesica/graphdoc/parsers"
)

// HTMLDoc provides the fields necessary for creating an HTML document based on
// a graph data model.
type HTMLDoc struct {
    HTML string
    DOT string
}

const htmlTemplateString = `
<!doctype html>
<html>
<head>
<meta charset="UTF-8">
<style type="text/css">
body {
    color: #555555;
    margin: auto;
    padding: 0 1.0em;
    max-width: 900px;
}
.graphdoc-graph {
    border: 1px solid #587058;
    padding-bottom: 1.0em;
    margin-top: 1.0em;
}
.graphdoc-graph-title {
    background: #587058;
    color: #fff;
    padding: 0.5em;
    margin-top: 0;
}
.graphdoc-graph-desc {
    margin-left: 1.0em;
    margin-right: 1.0em;
}
.graphdoc-node {
    border: 1px solid #FFD800;
    margin: 1.0em 1.0em 0 1.0em;
}
.graphdoc-node h3, .graphdoc-node-desc {
    margin-left: 1.0em;
    margin-right: 1.0em;
}
.graphdoc-node-label {
    background: #FFD800;
    color: #fff;
    padding: 0.5em;
    margin-top: 0;
}
.graphdoc-rel, .graphdoc-prop {
    border: 1px solid #587498;
    margin: 0 1.0em 1.0em 1.0em;
}
.graphdoc-prop {
    border: 1px solid #E86850;
}
.graphdoc-rel-label, .graphdoc-prop-name {
    color: #fff;
    margin: 0;
    background: #587498;
    padding: 0.5em;
}
.graphdoc-prop-name {
    background: #E86850;
}
.graphdoc-rel-desc, .graphdoc-prop-desc {
    padding: 0.5em;
}
.graphdoc-rel-path, .graphdoc-prop-type {
    float: right;
    font-family: mono;
}
.graphdoc-footer {
    margin: 1.0em 0;
    text-align: center;
}
#graphdoc-graph-viz {
    width: 100%%;
    height: 400px;
}
</style>
<link href="http://visjs.org/dist/vis.css" rel="stylesheet" type="text/css" />
<script src="http://visjs.org/dist/vis.js"></script>
</head>
<body>
{{ .HTML }}
<div class="graphdoc-footer">
    <small>Generated using GraphDoc. A George Lesica joint.</small>
</div>
<script type="text/javascript">
var container = document.getElementById('graphdoc-graph-viz');
var data = {
    dot: '{{ .DOT }}'
};
var options = {};
var graph = new vis.Network(container, data, options);
</script>
</body>
</html>
`
var htmlTemplate = template.Must(template.New("html").Parse(htmlTemplateString))

// HTMLDocument creates and returns an HTML document based on the provided
// graph data model and parser as a string.
func HTMLDocument(graph pgraph.Graph, parser parsers.Parser) string {
    dotStr := strings.Join(strings.Split(graph.ToDOT(), "\n"), "")
    htmlStr := graph.ToHTML(parser)
    var out bytes.Buffer
    htmlTemplate.Execute(&out, HTMLDoc{htmlStr, dotStr})
    return out.String()
}
