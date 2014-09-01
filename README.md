# GraphDoc

A system for building useful documentation of graph data models. I'm still
figuring out just how I want this thing to work, but in general the idea is that
the user can provide a simple file, like a Markdown file, with some special tags
added. GraphDoc will then generate a nice, usable web page or PDF. My plan is to
include a visualization along with the documentation. Ideally, the HTML version
of the documentation will feature an interactive visualization.

## Usage

Right now there isn't much here. There are four tags you can include, see the
`examples` directory for one or more examples. The four tags let you specify
graphs, nodes, relationships, and properties (relationships can't have
properties yet). Here are some examples of the tags:

```
::Graph GraphTitle::

::Node NodeName::

::Rel RELATED_TO -> NodeName::

::Prop propertyName:str::
```

Once you have created a document, you can use the command line tool to compile
it. Right now GraphDoc only produces HTML (and not very intelligently). For
example, to compile the "family" example, from the repo directory:

```
$ go install
$ graphdoc --outpath=family.html examples/family.md
```

Then take a look at `family.html`.

## Documentation

[![GoDoc](https://godoc.org/github.com/glesica/graphdoc?status.svg)](https://godoc.org/github.com/glesica/graphdoc)
