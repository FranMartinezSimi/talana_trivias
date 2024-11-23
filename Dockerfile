FROM golang:1.22-alpine

WORKDIR /app

RUN apk add --no-cache git bash postgresql-client

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN chmod +x wait-for.sh

RUN go build -o main .

EXPOSE 8080

CMD ["./wait-for.sh", "postgres:5432", "--", "./main"]