.PHONY: test lint check lines
test:
		go test -v -race -timeout 45s ./...

.DEFAULT_GOAL := test

lint check:
	golangci-lint run

lines:
	find . -name '*.go' | xargs wc -l
