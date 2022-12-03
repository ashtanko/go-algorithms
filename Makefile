.PHONY: test
test:
		go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := test

lint:
	golangci-lint run

lines:
	find . -name '*.go' | xargs wc -l
