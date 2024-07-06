ARG GO_VERSION=1.22.4
FROM golang:${GO_VERSION}-alpine AS base

WORKDIR /app

RUN apk update && \
    apk add git mysql-client curl && \
    curl -sSf https://atlasgo.sh | sh

COPY go.mod go.sum /app/

RUN go mod tidy

## 開発環境用ステージ
FROM base AS development

RUN go install github.com/air-verse/air@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest

CMD ["air", "-c", ".air.toml"]

## ビルド用ステージ
FROM base AS builder
ARG TARGET_APP=auth

COPY . /app
WORKDIR /app/cmd/${TARGET_APP}

RUN mkdir /dist && \
    go build -tags netgo -ldflags '-s -w' -o /dist/auth

## デプロイ先環境用ステージ
FROM base AS deployment

COPY --from=builder /dist/auth /app/auth

CMD ["/app/auth"]
