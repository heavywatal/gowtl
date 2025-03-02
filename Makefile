MODULE := $(shell go list -m)
CMD_PACKAGES := $(shell go list ./cmd/...)
BINARIES := ${CMD_PACKAGES:${MODULE}/cmd/%=bin/%}
PKG_SRCS := $(shell fd -t f '\.go$$' pkg)

.PHONY: build install

build: ${BINARIES}
	@:

install:
	go install -v ./cmd/...

bin/%: cmd/%/main.go ${PKG_SRCS}
	go build -v -o $@ $<
