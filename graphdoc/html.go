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
.graphdoc-graph-title {
    background: #554600;
    color: #FFF0AA;
    padding: 0.5em;
}
.graphdoc-graph-desc {
    margin-left: 1.0em;
    margin-right: 1.0em;
}
.graphdoc-node {
    border: 1px solid #440026;
    margin-top: 1.0em;
}
.graphdoc-node h3, .graphdoc-node-desc {
    margin-left: 1.0em;
    margin-right: 1.0em;
}
.graphdoc-node-label {
    color: #CD88AF;
    padding: 0.5em;
    background: #440026;
    margin-top: 0;
}
.graphdoc-rel, .graphdoc-prop {
    border: 1px solid #236467;
    margin: 0 1.0em 1.0em 1.0em;
}
.graphdoc-rel-label, .graphdoc-prop-name {
    color: #67989A;
    margin: 0;
    background: #003133;
    padding: 0.5em;
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
