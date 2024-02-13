FROM golang:1.22 AS builder

ARG VERSION=dev

COPY . /go/src/app
WORKDIR /go/src/app

RUN CGO_ENABLED=0 GOOS=linux go build -o service -ldflags="-X 'main.version=$VERSION'" main.go

FROM alpine:3.19

RUN mkdir -p /opt/auth-service
WORKDIR /opt/auth-service
RUN mkdir include
RUN mkdir data
COPY --from=builder /go/src/app/service service
COPY --from=builder /go/src/app/include include

HEALTHCHECK --interval=10s --timeout=5s --retries=3 CMD wget -nv -t1 --spider 'http://localhost/health-check' || exit 1

ENTRYPOINT ["./service"]
