FROM ubuntu:xenial as builder

RUN apt-get update \
     && apt-get install -y wget software-properties-common \
     && rm -rf /var/lib/apt/lists/*

WORKDIR "/root"

RUN add-apt-repository -y ppa:ethereum/ethereum

RUN apt-get update \
     && apt-get install -y ethereum

COPY docker/bootnode/entrypoint.sh ./

EXPOSE 30301

ENTRYPOINT ["./entrypoint.sh"]
CMD ["-verbosity", "6", "-nodekey", "bootnode.key", "--addr", ":30301"]
