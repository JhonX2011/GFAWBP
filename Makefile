# -- Comandos Go
GOCMD=go
GOBUILD=${GOCMD} build
GOCLEAN=${GOCMD} clean
GOTEST=${GOCMD} test ./...
GOGET=${GOCMD} get -u -v
GOMODTIDY=${GOCMD} mod tidy
GO_FMT=${GOCMD} fmt ./...
GORUN=${GOCMD} run
GOFMT=gofmt

lint:
	@golangci-lint -c ./public/linter/golangci-lint.yaml --out-format "colored-line-number" run ./... || exit 2

test-local:
	${GOCMD} test -race -v ./...

test:
	@go test -v -covermode="count" -coverprofile=./public/tmp/cover.out ./... && \
	echo "\n\ncoverage report\n" && \
	go tool cover -func=./public/tmp/cover.out

test-coverage: test
	@go tool cover -html ./public/tmp/cover.out -o ./public/tmp/cover.html && \
	open ./public/tmp/cover.html