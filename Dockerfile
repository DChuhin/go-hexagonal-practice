ARG GO_VERSION=1.21
ARG ALPINE_VERSION=3.19

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as builder

COPY . /build

RUN cd /build && \
    go mod tidy && \
    go mod download && \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /service.bin cmd/main.go

FROM alpine:${ALPINE_VERSION}
RUN mkdir -p /app
COPY --from=builder /service.bin /app/service.bin
RUN chmod +x /app/service.bin

WORKDIR /app
ENTRYPOINT ["/app/service.bin"]
