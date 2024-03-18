NAME=vec-calc-web
BINDIR=build
VERSION=1.0.0
BUILDTIME=$(shell date -u)
GOBUILD=go mod tidy && go build -ldflags '-s -w -X "main.version=$(VERSION)" -X "main.buildTime=$(BUILDTIME)"'

.PHONY: init

all: linux-amd64 windows-amd64

init:
	go mod tidy

run:
	cd $(BINDIR) && ./$(NAME)-linux-amd64-$(VERSION)

clean:
	rm -rf $(BINDIR)/*

bench:
	cd $(BINDIR) && ./$(NAME)-linux-amd64-$(VERSION) & sleep 1 && \
	wrk -t 20 -c 10000 -d 180s -s bench.lua --latency "http://localhost:8080/api/calc/mul"

linux-amd64: 
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./$(BINDIR)/$(NAME)-$@-$(VERSION)

windows-amd64: 
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o ./$(BINDIR)/$(NAME)-$@-$(VERSION).exe