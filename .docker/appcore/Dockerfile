FROM golang:1.20.3-alpine

WORKDIR /app

RUN apk update && apk upgrade && \
    apk add build-base && \
    go install github.com/cosmtrek/air@latest

WORKDIR /app/cmd/app/core

CMD ["air"]
