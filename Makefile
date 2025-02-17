GO = go

.PHONY: run build fmt

run: build
	./riven

build: fmt
	$(GO) build -buildvcs=false -ldflags='-s -w' -v

fmt:
	$(GO) fmt ./...
