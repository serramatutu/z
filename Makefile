githooks:
	sh scripts/githooks.sh

build:
	go build -o bin/z cmd/z/main.go

test:
	go test ./...
