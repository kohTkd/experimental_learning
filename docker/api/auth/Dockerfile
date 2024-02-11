ARG GO_VERSION=1.22
FROM golang:${GO_VERSION}-alpine

WORKDIR /app

COPY go.mod go.sum /app/

CMD [go run src/main.go]
