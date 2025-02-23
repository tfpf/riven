GO = go

.PHONY: all run build fmt

all: fmt run

run: build
	./riven

build:
	$(GO) build -buildvcs=false -ldflags='-s -w' -v

fmt:
	$(GO) fmt ./...
