MARKDOWN=$(wildcard examples/*.md)
HTML=$(MARKDOWN:.md=.html)
GOFILES=$(wildcard */*.go)

examples : graphdoc $(HTML)

graphdoc : graphdoc.go $(GOFILES)
	go build

%.html : %.md
	./graphdoc --outpath=$@ --informat=md $<

clean :
	rm examples/*.html
