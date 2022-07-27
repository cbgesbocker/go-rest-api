FROM golang:alpine3.15

WORKDIR /app

RUN apk add build-base

COPY go.mod ./

COPY go.sum ./

COPY *.go ./

COPY ./main.go ./

COPY ./src ./src

RUN go mod download

RUN go build -o ./docker-gs-ping

EXPOSE 8080

ENTRYPOINT [ "go", "run", "main.go" ]