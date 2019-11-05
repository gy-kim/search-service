.PHONY: test
test:
	go test ./... -v -count=1

.PHONY: build
build:
	go mod tidy
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o search-service