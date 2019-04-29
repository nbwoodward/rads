GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFILE=parse_ndx.go
BINARY_NAME=parse_ndx
BINARY_UNIX=$(BINARY_NAME)_unix

all: build test

build:
	$(GOBUILD) $(GOFILE)

test:
	$(GOTEST) -v ./...
	#./parse_ndx

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run:
	./$(BINARY_NAME)
