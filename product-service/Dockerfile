# Builder

FROM golang:1.19-alpine AS builder

ARG GITHUB_PATH=github.com/alexeykirinyuk/go_grpc_workshop/product-service
WORKDIR /home/${GITHUB_PATH}

RUN apk add --update make git protoc protobuf protobuf-dev curl

RUN BIN="/usr/local/bin" && VERSION="1.7.0" && \
  curl -sSL "https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" \
  -o "${BIN}/buf" && \
  chmod +x "${BIN}/buf"

COPY Makefile Makefile
RUN make deps-go

COPY go.mod go.mod
COPY go.sum go.sum
COPY pkg pkg
RUN make app-deps

COPY . .
RUN make build

# gRPC Server

FROM alpine:latest as server
LABEL org.opencontainers.image.source https://${GITHUB_PATH}
RUN apk --no-cache add ca-certificates
WORKDIR /root/

ARG GITHUB_PATH=github.com/alexeykirinyuk/go_grpc_workshop/product-service

COPY --from=builder /home/${GITHUB_PATH}/bin/grpc-server .
COPY --from=builder /home/${GITHUB_PATH}/config.yml .
COPY --from=builder /home/${GITHUB_PATH}/migrations/ ./migrations

RUN chown root:root ./grpc-server
RUN chmod +x ./grpc-server

EXPOSE 7053

ENTRYPOINT ["/root/grpc-server"]
