.PHONY: build test test-integration lint clean

build:
	go build -o timesum .

test:
	go test -v ./...

test-integration:
	go test -v -tags=integration ./...

lint:
	golangci-lint run

format:
	gofmt -s -w .

snapshot-release:
	goreleaser release --snapshot

clean:
	rm -f timesum
	rm -rf dist
