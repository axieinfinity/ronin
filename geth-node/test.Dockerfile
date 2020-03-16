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

CMD exec ./geth --networkid="42" --verbosity=4 --rpc --rpcaddr "0.0.0.0" --rpccorsdomain "*" 

EXPOSE 8545 8546 8547 30303 30303/udp
