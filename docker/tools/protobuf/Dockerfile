ARG GO_VERSION=1.22
FROM golang:${GO_VERSION}-alpine

WORKDIR /src

RUN mkdir -p /dist && \
    mkdir -p /dist/server/go && \
    mkdir -p /dist/client/go && \
    mkdir -p /dist/server/ruby && \
    mkdir -p /dist/client/ruby

RUN apk update && \
    apk add ruby protoc && \
    gem install grpc-tools && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

CMD protoc --go-grpc_out=/dist/server/go --go_out=/dist/server/go --go_opt=paths=source_relative --go-grpc_opt require_unimplemented_servers=false --go-grpc_opt paths=source_relative --ruby_out=/dist/client/ruby ./**/*.proto
