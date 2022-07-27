FROM cosmtrek/air:v1.40.4

WORKDIR /app

# RUN apk add build-base

COPY go.mod ./

COPY go.sum ./

COPY *.go ./

COPY ./main.go ./

COPY ./.air.conf ./

COPY ./src ./src

RUN go mod download

EXPOSE 8080

ENTRYPOINT [ "air", "-c", ".air.conf" ]