ARG GO_VERSION=1.22
FROM golang:${GO_VERSION}-alpine

WORKDIR /src

RUN mkdir -p /dist && \
    mkdir -p /dist/server/go && \
    mkdir -p /dist/client/go && \
    mkdir -p /dist/server/ruby && \
    mkdir -p /dist/client/ruby

RUN apk update && \
    apk add --no-cache ruby protoc gcompat && \
    # ln -s /lib/libc.musl-x86_64.so.1 /lib/ld-linux-x86-64.so.2 && \
    gem install grpc-tools && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
    gem install grpc grpc-tools

COPY protoc.sh /usr/local/bin/protoc.sh

CMD sh /usr/local/bin/protoc.sh
