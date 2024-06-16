ARG GO_VERSION=1.22.4
FROM golang:${GO_VERSION}-alpine

WORKDIR /app

RUN apk update && \
    apk add git mysql-client curl && \
    go install github.com/air-verse/air@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest && \
    curl -sSf https://atlasgo.sh | sh

COPY go.mod go.sum /app/

RUN go mod download

COPY . /app

CMD ["air", "-c", ".air.toml"]
