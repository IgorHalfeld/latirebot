FROM golang:1.13.8-alpine3.11 

ADD . /usr/src/myapp

WORKDIR /usr/src/myapp

RUN apk add build-base

RUN go get -u github.com/pressly/goose/cmd/goose
RUN goose --dir migrations sqlite3 latire-production.db up

ENV APP_ENV=production

RUN go build -v -o run-latire main.go
RUN mv run-latire $GOPATH/bin

