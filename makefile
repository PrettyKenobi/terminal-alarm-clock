GOCMD=go
GOTEST=$(GOCMD) test
BINARY_NAME=terminal-alarm-clock
VERSION?=0.0.0
OUTPUT_DIR=out/bin

.PHONY: test build clean

build:
	mkdir -p $(OUTPUT_DIR)
	$(GOCMD) build -mod vendor -o $(OUTPUT_DIR)/$(BINARY_NAME) .

clean:
	rm -fr ./bin
	rm -fr ./out
	rm -f ./coverage.xml ./profile.cov
	rm -f ./... *.*~

test:
	$(GOTEST) -v -race ./...

coverage:
	$(GOTEST) -cover -covermode=count -coverprofile=profile.cov ./...
	$(GOCMD) tool cover -func profile.cov
