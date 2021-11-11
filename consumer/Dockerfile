FROM golang:1.17 AS builder

RUN rm -rf /Consumer
RUN mkdir /Consumer

ADD . /Consumer

WORKDIR /Consumer

RUN make setup
RUN make build-consumer

RUN chmod +x wait-for-it.sh
