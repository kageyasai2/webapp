FROM golang:1.8

RUN apt-get update -qq

WORKDIR /go/src/app
COPY ./app .

RUN go build