GO = go

.PHONY: run

run:
	$(GO) fmt ./...
	$(GO) build -buildvcs=false -ldflags='-s -w' -v
	./riven
