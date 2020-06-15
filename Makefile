# Go parameters, thx to https://sohlich.github.io/post/go_makefile/
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOINSTALL=$(GOCMD) install
GOGET=$(GOCMD) get

build : clean
	$(GOBUILD) -v

test : clean build
	$(GOTEST) ./... -v

install : clean
	$(GOINSTALL)

clean :
	$(GOCLEAN)

.DEFAULT_GOAL = build
