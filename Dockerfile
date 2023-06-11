ARG GO_VERSION=1.16
ARG BINARY_DIR=/go/bin/myutility
FROM golang:${GO_VERSION}-alpine AS builder
RUN apk --no-cache add git bash && \
    mkdir src/myutility.com && \
    mkdir src/myutility.com/cmd && \
    mkdir src/myutility.com/pkg

WORKDIR /src/myutility.com

COPY go.mod go.mod
COPY go.sum go.sum
COPY /cmd cmd
COPY /pkg pkg

RUN go mod download && \
    CGO_ENABLED=0 go build -o /src/myutility.com/tmp/bin/timezone ./cmd/timezone/timezone_converter.go

FROM alpine:3.13
RUN apk --no-cache add git bash vim curl && \
    mkdir /go && \
    mkdir /go/bin && \
    mkdir /go/bin/myutility

WORKDIR /go/bin/myutility
COPY --from=builder /src/myutility.com/tmp/bin/* /go/bin/myutility/

CMD exec /go/bin/myutility/timezone

