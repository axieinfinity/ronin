# Build Geth in a stock Go builder container
FROM golang:1.20.0-alpine3.17 as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

COPY . /opt
RUN cd /opt && make ronin

# Pull Geth into a second stage deploy alpine container
FROM alpine:3.13

RUN apk add --no-cache ca-certificates
WORKDIR "/opt"

ENV PASSWORD ''
ENV PRIVATE_KEY ''
ENV BOOTNODES ''
ENV VERBOSITY 3
ENV SYNC_MODE 'snap'
ENV NETWORK_ID '2021'
ENV ETHSTATS_ENDPOINT ''
ENV RPC_NODE ''
ENV NODEKEY ''
ENV FORCE_INIT 'true'
ENV RONIN_PARAMS ''
ENV INIT_FORCE_OVERRIDE_CHAIN_CONFIG 'false'
ENV ENABLE_FAST_FINALITY 'true'
ENV ENABLE_FAST_FINALITY_SIGN 'true'
ENV BLS_PRIVATE_KEY ''
ENV BLS_PASSWORD ''
ENV BLS_AUTO_GENERATE 'false'

COPY --from=builder /opt/build/bin/ronin /usr/local/bin/ronin
COPY --from=builder /opt/genesis/ ./
COPY --from=builder /opt/docker/chainnode/entrypoint.sh ./

EXPOSE 7000 6060 8545 8546 30303 30303/udp

ENTRYPOINT ["./entrypoint.sh"]
