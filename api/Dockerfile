# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang:1.12-alpine base image
FROM golang:1.21.0-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && apk add --no-cache bash

RUN apk add --no-cache make protobuf-dev

RUN go install github.com/cosmtrek/air@v1.29.0 && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 && \
    go install github.com/ktr0731/evans@v0.10.11

RUN export PATH="$PATH:$(go env GOPATH)/bin"

# Set the Current Working Directory inside the container
WORKDIR /app

COPY .air.toml /

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["air", "-c", "/.air.toml"]