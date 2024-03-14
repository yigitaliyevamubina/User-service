FROM golang:1.20-alpine3.16 AS builder

RUN mkdir app
COPY . /app

WORKDIR /app

RUN go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate

RUN go build -o main cmd/main.go

CMD ["/app/main"]   