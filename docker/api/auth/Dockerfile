ARG GO_VERSION=1.22
FROM golang:${GO_VERSION}-alpine

WORKDIR /app

RUN apk update && \
    apk add git mysql-client && \
    go install github.com/cosmtrek/air@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest

COPY go.mod go.sum /app/

RUN go mod download

COPY . /app

CMD ["air", "-c", ".air.toml"]
