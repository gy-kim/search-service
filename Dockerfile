FROM golang:1.13-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o search-service .

EXPOSE 9000

CMD ["./search-service"]