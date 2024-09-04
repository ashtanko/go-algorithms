.PHONY: test lint check lines l coverage
test:
	go test -v -race -timeout 45s ./...

.DEFAULT_GOAL := test

lint check l:
	golangci-lint run internal/... cmd/...

coverage:
	go test -coverprofile c.out ./...

lines:
	find . -name '*.go' | xargs wc -l
