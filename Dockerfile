# Build Geth in a stock Go builder container
FROM golang:1.22.8-alpine3.19@sha256:fe5bea2e1ab3ffebe0267393fea88fcb197e2dbbb1e2dbabeec6dd9ccb0e1871 as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git libstdc++-dev

COPY . /opt
RUN cd /opt && make ronin

# Pull Geth into a second stage deploy alpine container
FROM alpine:3.19@sha256:ae65dbf8749a7d4527648ccee1fa3deb6bfcae34cbc30fc67aa45c44dcaa90ee

RUN apk add --no-cache ca-certificates
WORKDIR "/opt"

ENV PASSWORD ''
ENV PRIVATE_KEY ''
ENV BOOTNODES ''
ENV VERBOSITY 3
ENV SYNC_MODE 'snap'
ENV NETWORK_ID '2021'
ENV ETHSTATS_ENDPOINT ''
ENV NODEKEY ''
ENV FORCE_INIT 'true'
ENV RONIN_PARAMS ''
ENV INIT_FORCE_OVERRIDE_CHAIN_CONFIG 'false'
ENV ENABLE_FAST_FINALITY 'true'
ENV ENABLE_FAST_FINALITY_SIGN 'false'
ENV BLS_PRIVATE_KEY ''
ENV BLS_PASSWORD ''
ENV BLS_AUTO_GENERATE 'false'
ENV BLS_SHOW_PRIVATE_KEY 'false'
ENV GENERATE_BLS_PROOF 'false'

COPY --from=builder /opt/build/bin/ronin /usr/local/bin/ronin
COPY --from=builder /opt/genesis/ ./
COPY --from=builder /opt/docker/chainnode/entrypoint.sh ./

EXPOSE 7000 6060 8545 8546 30303 30303/udp

ENTRYPOINT ["./entrypoint.sh"]
