.PHONY: test lint check lines coverage
test:
		go test -v -race -timeout 45s ./...

.DEFAULT_GOAL := test

lint check:
	golangci-lint run

coverage:
	go test -coverprofile c.out ./...

lines:
	find . -name '*.go' | xargs wc -l
