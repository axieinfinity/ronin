FROM golang:1.17-buster AS builder
WORKDIR /opt
RUN mkdir /opt/ronin
COPY . /opt/ronin