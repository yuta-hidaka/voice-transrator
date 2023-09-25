#!/bin/bash

echo "building..."

go mod tidy

echo "generating server gropc..."

PROTOC_PARAMS=(
    -I /app/proto/v1
    # go
    --go_out=./internal/proto
    --go_opt=paths=source_relative
    # grpc
    --go-grpc_out=./internal/proto
    --go-grpc_opt=paths=source_relative
    # gateway
    --grpc-gateway_out=./internal/proto
    --grpc-gateway_opt=logtostderr=true
    --grpc-gateway_opt=paths=source_relative
    --grpc-gateway_opt=generate_unbound_methods=true
    # gate way ts
    --grpc-gateway-ts_out=/proto
    # open API
    # --openapiv2_out ./internal/proto/openapi \
    # --openapiv2_opt logtostderr=true \
    # target
    /app/proto/v1/*.proto
)

protoc "${PROTOC_PARAMS[@]}"

# Client
# echo "generating client gropc..."
# protoc --ts_out=/ --proto_path /app /app/proto/v1/*.proto

echo "building go binary..."

go build -o ./tmp /app/cmd/api/main.go