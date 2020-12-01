# Build Geth in a stock Go builder container
FROM golang:1.15-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

COPY . /opt
RUN cd /opt && make ronin

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
ENV ETHSTATS_ENDPOINT ''
ENV RPC_NODE ''


COPY --from=builder /opt/build/bin/ronin /usr/local/bin/ronin
COPY --from=builder /opt/genesis/ ./
COPY --from=builder /opt/docker/chainnode/entrypoint.sh ./

EXPOSE 8545 8546 30303 30303/udp

ENTRYPOINT ["./entrypoint.sh"]
