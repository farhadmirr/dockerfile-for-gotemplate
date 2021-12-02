# syntax=docker/dockerfile:1

FROM golang:1.15-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY ./templates ./

ADD templates ./templates

COPY ./* ./
RUN go build -o /docker-gs-ping

CMD [ "/docker-gs-ping" ]

EXPOSE 9090
