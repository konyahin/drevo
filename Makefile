.POSIX:
.SUFFIXES:
.PHONY: all test install clean

all: check drevo test install

drevo: *.go
	gofmt -w . 
	go build

test: drevo
	go test

install: drevo
	go install .

check:
	go vet
	staticcheck

clean:
	rm drevo