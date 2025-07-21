FROM golang:1.23.2-alpine

WORKDIR /usr/src/app

COPY . .

RUN go mod download

