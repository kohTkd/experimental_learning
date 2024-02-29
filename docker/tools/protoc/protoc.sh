#!/bin/sh

function exec_protoc() {
  local service="${1}"
  cd /src/${service}

  echo "Starting ${service}"

  protoc \
    --go-grpc_out=/dist/server/${service}/go \
    --go_out=/dist/server/${service}/go \
    --go_opt=paths=source_relative \
    --go-grpc_opt paths=source_relative \
    ./*.proto

  grpc_tools_ruby_protoc \
    --ruby_out=/dist/client/ruby \
    --grpc_out=/dist/client/ruby \
    ./*.proto
  echo "Finished ${service}"
}

exec_protoc auth
