# Build Geth in a stock Go builder container
FROM golang:1.13-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

COPY . /go-ethereum
RUN cd /go-ethereum && make geth

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
WORKDIR "/opt"

ENV PASSWORD ''
ENV PRIVATE_KEY ''
ENV BOOTNODES ''
ENV VERBOSITY 3
ENV SYNC_MODE 'full'
ENV NETWORK_ID '2021'


COPY --from=builder /go-ethereum/build/bin/geth /usr/local/bin/geth
COPY --from=builder /go-ethereum/genesis/ ./
COPY --from=builder /go-ethereum/docker/chainnode/entrypoint.sh ./

EXPOSE 8545 8546 30303 30303/udp

ENTRYPOINT ["./entrypoint.sh"]
