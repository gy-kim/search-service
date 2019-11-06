.PHONY: test
test:
	go test ./... -v -count=1

.PHONY: build
build:
	go mod tidy
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o search-service

.PHONY: up
up:
	docker-compose up --build -d

.PHONY: down
down:
	docker-compose down

.PHONY: clean
clean:
	docker-compose down -v --rmi all --remove-orphans