.PHONY: lint test cover
lint:
	golangci-lint run --timeout=1m ./...

test:
	go test -shuffle=on -race -covermode=atomic -coverprofile=coverage.out ./...

cover:
	go tool cover -func=coverage.out | tail -n1
