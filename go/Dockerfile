FROM golang:1.19-alpine3.18

RUN apk add --update --no-cache alpine-sdk git

WORKDIR /app

COPY . .

RUN go build ./...
