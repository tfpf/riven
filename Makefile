GO = go

.PHONY: run build fmt

run: fmt build
	./riven

build:
	$(GO) build -buildvcs=false -ldflags='-s -w' -v

fmt:
	$(GO) fmt ./...
