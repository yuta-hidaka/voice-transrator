# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang:1.12-alpine base image
FROM golang:1.21.1-alpine3.18

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
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
    # libssl-dev \
    zip \
    npm

RUN go install github.com/cosmtrek/air@latest && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
    go install github.com/ktr0731/evans@v0.10.11

RUN go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

RUN npm install -g @protobuf-ts/plugin

RUN mkdir /proto

RUN export PATH="$PATH:$(go env GOPATH)/bin"

# Set the Current Working Directory inside the container
WORKDIR /app

COPY .air.toml /

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["air", "-c", "/.air.toml"]