MARKDOWN=$(wildcard examples/*.md)
HTML=$(MARKDOWN:.md=.html)

examples : graphdoc $(HTML)

graphdoc :
	go build

%.html : %.md
	./graphdoc --outpath=$@ --informat=md $<

clean :
	rm examples/*.html
