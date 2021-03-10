FROM golang:alpine as builder

RUN apk update && apk upgrade

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY . .

RUN go mod download

RUN go build ./cmd/server/main.go

EXPOSE 8080

CMD ["/app/main"]