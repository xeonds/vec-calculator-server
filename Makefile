NAME=vec-calc-web
BINDIR=build
VERSION=1.0.0
BUILDTIME=$(shell date -u)
GOBUILD=go mod tidy && go build -ldflags '-s -w -X "main.version=$(VERSION)" -X "main.buildTime=$(BUILDTIME)"'

.PHONY: init

all: linux-amd64 windows-amd64

init:
	(go mod tidy)

linux-amd64: 
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./$(BINDIR)/$(NAME)-$@-$(VERSION)

linux-amd64-dev: linux-amd64
	cd $(BINDIR) && ./$(NAME)-linux-amd64-$(VERSION)

windows-amd64: 
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o ./$(BINDIR)/$(NAME)-$@-$(VERSION).exe

clean:
	rm -rf $(BINDIR)/*
