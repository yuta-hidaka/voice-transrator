FROM golang:1.21.1-alpine3.18

RUN apk update && apk upgrade && apk add --no-cache \   
    bash \
    make \ 
    protobuf-dev \
    autoconf \
    automake \
    alpine-sdk \
    curl \
    git \
    libtool \
    zip \
    npm

RUN go install github.com/cosmtrek/air@latest && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
    go install github.com/ktr0731/evans@v0.10.11 && \
    go install github.com/grpc-ecosystem/protoc-gen-grpc-gateway-ts@latest

RUN go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest


RUN npm install -g @protobuf-ts/plugin

RUN mkdir /proto

RUN export PATH="$PATH:$(go env GOPATH)/bin"

WORKDIR /app

COPY .air.toml /

EXPOSE 8080

CMD ["air", "-c", "/.air.toml"]