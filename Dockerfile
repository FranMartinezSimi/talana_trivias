# Dockerfile
FROM golang:1.21-alpine

WORKDIR /app

RUN apk add --no-cache git bash

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY wait-for.sh /wait-for.sh

RUN go build -o main .

EXPOSE 8080

CMD ["/wait-for.sh", "postgres", "./main"]
