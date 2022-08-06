# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./
COPY Config/ ./
COPY Services/ ./
COPY config.yml ./

RUN go build -o /DressedApi

EXPOSE 8080

CMD [ "/DressedApi" ]