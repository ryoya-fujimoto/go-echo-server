FROM golang:1.2 AS build-env

ADD . /work
WORKDIR /work
RUN go build -o echo-server main.go

FROM ubuntu
COPY --from=build-env /work/echo-server /usr/local/bin/echo-server

RUN apt-get -y update; apt-get -y install nginx
