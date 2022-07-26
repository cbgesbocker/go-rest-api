FROM golang:alpine3.15

WORKDIR /app

RUN apk add build-base

COPY go.mod ./

COPY go.sum ./

COPY *.go ./

COPY ./controllers ./controllers

COPY ./models ./models

VOLUME [ "./data", "./data" ]

RUN go mod download

RUN go build -o ./docker-gs-ping

EXPOSE 8080

ENTRYPOINT [ "./docker-gs-ping" ]