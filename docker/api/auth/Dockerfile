ARG GO_VERSION=1.22
FROM golang:${GO_VERSION}-alpine

WORKDIR /app

RUN apk update && \
    apk add git && \
    go install github.com/cosmtrek/air@latest

COPY go.mod go.sum /app/

RUN go mod download

CMD ["air", "-c", ".air.toml"]
