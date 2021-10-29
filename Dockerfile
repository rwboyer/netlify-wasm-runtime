# syntax=docker/dockerfile:1

FROM golang:1.17
WORKDIR /go/src/app
COPY . .

#RUN go get -d -v ./...
