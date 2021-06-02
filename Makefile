githooks:
	sh scripts/githooks.sh

build: test
	go build -o bin/z cmd/z/main.go

test:
	go test ./...

pr_ci_test:
	act pull_request -s GITHUB_TOKEN=${GITHUB_TOKEN}

release_test:
	goreleaser --snapshot --skip-publish --rm-dist
