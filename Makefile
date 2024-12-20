.PHONY: test t lint check lines l coverage
test t:
	go test -v -race -timeout 45s ./...

.DEFAULT_GOAL := test

lint check l:
	golangci-lint run internal/... cmd/...

coverage:
	go test -coverprofile c.out ./...

lines:
	find . -name '*.go' | xargs wc -l

# install lint: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
