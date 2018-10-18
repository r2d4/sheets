FROM golang:1.10.1 as builder
WORKDIR /go/src/github.com/r2d4/sh8s
COPY vendor ./vendor
COPY Makefile .

COPY cmd ./cmd
COPY pkg ./pkg

RUN make install

FROM alpine:3.7  
CMD ["./sh8s", "serve"]
COPY --from=builder /go/bin/sh8s .