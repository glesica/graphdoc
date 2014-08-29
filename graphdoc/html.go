package graphdoc

import (
    "fmt"
    "github.com/glesica/graphdoc/pgraph"
)

const htmlTemplate = `
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
</style>
</head>
<body>
%s
<div class="graphdoc-footer">
    <small>Generated using GraphDoc. A George Lesica joint.</small>
</div>
</body>
</html>
`

func HTMLDocument(graph pgraph.Graph) string {
    return fmt.Sprintf(htmlTemplate, graph.ToHTML())
}
