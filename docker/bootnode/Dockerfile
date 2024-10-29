# Build Geth in a stock Go builder container
FROM golang:1.15.5-alpine AS builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

COPY . /go-ethereum
RUN cd /go-ethereum && make bootnode

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
WORKDIR "/opt"

COPY --from=builder /go-ethereum/build/bin/bootnode /usr/local/bin/bootnode
COPY --from=builder /go-ethereum/docker/bootnode/entrypoint.sh ./

ENTRYPOINT ["./entrypoint.sh"]
CMD ["-verbosity", "6", "-nodekey", "bootnode.key", "--addr", ":30301"]

EXPOSE 30301/udp
EXPOSE 30303/udp
