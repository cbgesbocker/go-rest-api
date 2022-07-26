FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

COPY *.go ./

COPY ./controllers ./controllers

COPY ./models ./models

RUN go mod download

RUN go build -o ./docker-gs-ping

EXPOSE 8080

CMD [ "./docker-gs-ping" ]