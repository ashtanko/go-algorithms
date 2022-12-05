.PHONY: test lint check lines
test:
		go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := test

lint check:
	golangci-lint run

lines:
	find . -name '*.go' | xargs wc -l
