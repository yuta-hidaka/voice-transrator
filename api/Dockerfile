# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang:1.12-alpine base image
FROM golang:1.21.0-alpine

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

RUN go install github.com/cosmtrek/air@v1.29.0 && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 && \
    go install github.com/ktr0731/evans@v0.10.11

RUN npm install -g @bazel/bazelisk
    
# RUN npm install -g @bufbuild/buf

# RUN npm install -g protoc-gen-js

# RUN wget https://github.com/protocolbuffers/protobuf-javascript/releases/download/v3.21.2/protobuf-javascript-3.21.2-linux-aarch_64.tar.gz
# RUN tar -xvf protobuf-javascript-3.21.2-linux-aarch_64.tar.gz
# RUN cp bin/protoc /usr/local/bin/protoc
# RUN wget -qO- https://get.pnpm.io/install.sh | ENV="$HOME/.bashrc" SHELL="$(which bash)" bash -

# ENV PNPM_HOME=/usr/local/bin

# RUN npm install -g pnpm 

# RUN pnpm setup --force

# RUN npm install google-protobuf

# RUN npm i -g protoc-gen-js grpc-tools google-protobuf

# RUN wget https://github.com/protocolbuffers/protobuf-javascript/releases/download/v3.21.2/protobuf-javascript-3.21.2-linux-aarch_64.tar.gz


RUN cd /tmp && wget https://github.com/grpc/grpc-web/archive/refs/tags/1.4.2.tar.gz && \
    tar -xvf 1.4.2.tar.gz && \
    cd grpc-web-1.4.2 && \
    make install-plugin

RUN export PATH="$PATH:$(go env GOPATH)/bin"

# Set the Current Working Directory inside the container
WORKDIR /app

COPY .air.toml /

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["air", "-c", "/.air.toml"]