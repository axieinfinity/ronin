# Build Geth in a stock Go builder container
FROM golang:1.13-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

ADD . /go-ethereum
RUN cd /go-ethereum && make geth

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
WORKDIR "/opt"
COPY --from=builder /go-ethereum/build/bin/geth .
COPY --from=builder /go-ethereum/genesis/consortium.json .

RUN ./geth init consortium.json

ARG password
ARG privatekey
RUN echo $password > ~/.accountpassword
RUN echo $privatekey > ~/.privatekey
RUN ./geth account import --password ~/.accountpassword  ~/.privatekey

CMD exec ./geth --bootnodes "enode://$bootnodeId@$bootnodeIp:30301" --networkid="42" --verbosity=4 --rpc --rpcaddr "0.0.0.0" --rpccorsdomain "*" --mine --etherbase $address --allow-insecure-unlock --unlock $address --password ~/.accountpassword

EXPOSE 8545 8546 8547 30303 30303/udp
