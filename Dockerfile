FROM golang:1.17

RUN rm -rf /CatMQ
RUN mkdir /CatMQ

ADD . /CatMQ

WORKDIR /CatMQ

RUN make setup
RUN make build

EXPOSE 23023
