MODULE := $(shell go list -m)
CMD_PACKAGES := $(shell go list ./cmd/...)
BINARIES := ${CMD_PACKAGES:${MODULE}/cmd/%=bin/%}

.PHONY: build install

build: ${BINARIES}
	@:

install:
	go install -v ./cmd/...

bin/%: cmd/%/main.go
	go build -v -o $@ $<
