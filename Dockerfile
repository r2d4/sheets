FROM golang:1.10.1 as builder
WORKDIR /go/src/github.com/r2d4/sh8s
COPY vendor ./vendor
COPY Makefile .

COPY cmd ./cmd
COPY pkg ./pkg
RUN mkdir -p /root/.cache/go-build
RUN  make

FROM alpine:3.7  
CMD ["./sh8s", "serve", "-a", "redis:6379"]
COPY --from=builder /go/src/github.com/r2d4/sh8s/out/sh8s .