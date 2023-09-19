#!/bin/bash

go mod tidy

protoc \
-I /app/proto \
--go_opt=module=voice-translator \
--go_out=. \
--go-grpc_opt=module=voice-translator \
--go-grpc_out=. \
/app/proto/*.proto

grpc_tools_node_protoc \
-I /app/proto \
--js_out=import_style=commonjs:/grpc-web-client \
--grpc-web_out=import_style=typescript,mode=grpcwebtext:/grpc-web-client \
/app/proto/*.proto

go build -o ./tmp /app/cmd/api/main.go