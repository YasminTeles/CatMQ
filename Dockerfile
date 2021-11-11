FROM golang:1.17

RUN mkdir /CatMQ

ADD . /CatMQ

WORKDIR /CatMQ

RUN make setup
RUN make build

EXPOSE 23023

CMD ["./main"]
