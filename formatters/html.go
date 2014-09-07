// The formatters package provides tools to format documentation based on a
// property graph data model.
package formatters

import (
    "bytes"
    //"fmt"
    //"strings"
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
    padding: 0;
    margin: 0;
}

.graphdoc-sidebar {
    position: fixed;
    width: 275px;
    height: 100%;
}
.graphdoc-sidebar h2 {
    margin: 0.5em;
}
.graphdoc-sidebar-content {
}
.graphdoc-sidebar-content ul {
    padding-left: 0;
    list-style: none;
}
.graphdoc-sidebar-content ul li {
}
.graphdoc-sidebar-content ul a {
    color: #333;
    display: block;
    padding: 0.25em;
    padding-left: 0.5em;
}
.graphdoc-sidebar-content ul a:hover {
    background: #eee;
}
.graphdoc-sidebar-content ul ul {
}
.graphdoc-sidebar-content ul ul a {
    padding-left: 1.0em;
}
.graphdoc-sidebar-content ul ul a:hover {
}
.graphdoc-sidebar-content ul ul ul {
}
.graphdoc-sidebar-content ul ul ul a {
    padding-left: 1.5em;
}
.graphdoc-sidebar-content ul ul ul a:hover {
}

.graphdoc-graph {
    border: 1px solid #587058;
    padding-bottom: 1.0em;
    margin-left: 275px;
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

.graphdoc-rel {
    border: 1px solid #587498;
    margin: 1.0em 1.0em 0 1.0em;
}
.graphdoc-rel h3, .graphdoc-rel-desc {
    margin-left: 1.0em;
    margin-right: 1.0em;
}
.graphdoc-rel-path {
    float: right;
    font-family: mono;
}
.graphdoc-rel-label {
    background: #587498;
    color: #fff;
    padding: 0.5em;
    margin-top: 0;
}

.graphdoc-prop-none {
    margin: 0 1.0em 1.0em 1.0em;
}
.graphdoc-prop-desc {
    padding: 0.5em;
}
.graphdoc-prop-type {
    float: right;
    font-family: mono;
}
.graphdoc-prop {
    border: 1px solid #587498;
    margin: 0 1.0em 1.0em 1.0em;
}
.graphdoc-prop {
    border: 1px solid #E86850;
}
.graphdoc-prop-name {
    background: #E86850;
    color: #fff;
    margin: 0;
    padding: 0.5em;
}
.graphdoc-footer {
    margin: 1.0em 0;
    text-align: center;
}
#graphdoc-graph-viz {
    width: 100%;
    height: 400px;
}
</style>
<link href="http://visjs.org/dist/vis.css" rel="stylesheet" type="text/css" />
<script src="http://visjs.org/dist/vis.js"></script>
</head>
<body>
<div class="graphdoc-sidebar">
    <div class="graphdoc-sidebar-content">
        <h2>Menu</h2>
        <ul>
            <li><a href="#graphdoc-graph-{{ .Title }}">{{ .Title }}</a></li>
            <ul>
                {{ range .Nodes }}
                <li><a href="#graphdoc-node-{{ .Label }}">{{ .Label }}</a><li>
                <ul>
                    {{ $nodeLabel := .Label }}
                    {{ range .Props }}
                    <li><a href="#graphdoc-node-{{ $nodeLabel }}-prop-{{ .Name }}">{{ .Name }}</a></li>
                    {{ end }}
                </ul>
                {{ end }}
            </ul>
            <ul>
                {{ range .Rels }}
                <li><a href="#graphdoc-rel-{{ .Label }}">{{ .Label }}</a></li>
                <ul>
                    {{ $relLabel := .Label }}
                    {{ range .Props }}
                    <li><a href="#graphdoc-rel-{{ $relLabel }}-prop-{{ .Name }}">{{ .Name }}</a></li>
                    {{ end }}
                </ul>
                {{ end }}
            </ul>
        </ul>
    </div>
</div>
<div class="graphdoc-graph">
    <h1 class="graphdoc-graph-title" id="graphdoc-graph-{{ .Title }}">{{ .Title }}</h1>
    <div class="graphdoc-graph-desc">{{ .Desc }}</div>
    <div id="graphdoc-graph-viz"></div>
    <div class="graphdoc-graph-nodes">
        {{ range .Nodes }}
        <div class="graphdoc-node">
            <h2 class="graphdoc-node-label" id="graphdoc-node-{{ .Label }}">{{ .Label }}</h2>
            <div class="graphdoc-node-desc">{{ .Desc }}</div>
            <div class="graphdoc-node-props">
                <h3>Properties</h3>
                {{ $nodeLabel := .Label }}
                {{ range .Props }}
                <div class="graphdoc-prop">
                    <h4 class="graphdoc-prop-name" id="graphdoc-node-{{ $nodeLabel }}-prop-{{ .Name }}">
                        {{ .Name }}
                        <small class="graphdoc-prop-type">{{ .Type }}</small>
                    </h4>
                    <div class="graphdoc-prop-desc">{{ .Desc }}</div>
                </div>
                {{ else }}
                <p class="graphdoc-prop-none">There are no properties on this relationship.</p>
                {{ end }}
            </div>
        </div>
        {{ end }}
    </div>
    <div class="graphdoc-graph-rels">
        {{ range .Rels }}
        <div class="graphdoc-rel">
            <h2 class="graphdoc-rel-label" id="graphdoc-rel-{{ .Label }}">
                {{ .Label }}
                <small class="graphdoc-rel-path">{{ .Source }} ➔ {{ .Target }}</small>
            </h2>
            <div class="graphdoc-rel-desc">{{ .Desc }}</div>
            <div class="graphdoc-rel-props">
                <h3>Properties</h3>
                {{ $relLabel := .Label }}
                {{ range .Props }}
                <div class="graphdoc-prop">
                    <h4 class="graphdoc-prop-name" id="graphdoc-rel-{{ $relLabel }}-prop-{{ .Name }}">
                        {{ .Name }}
                        <small class="graphdoc-prop-type">{{ .Type }}</small>
                    </h4>
                    <div class="graphdoc-prop-desc">{{ .Desc }}</div>
                </div>
                {{ else }}
                <p class="graphdoc-prop-none">There are no properties on this relationship.</p>
                {{ end }}
            </div>
        </div>
        {{ end }}
    </div>
</div>
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
func HTMLDocument(graph *pgraph.Graph, parser parsers.Parser) string {
    var out bytes.Buffer
    htmlTemplate.Execute(&out, graph)
    return out.String()
}
