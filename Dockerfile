FROM golang:1.13.8-alpine3.11 

ADD . /usr/src/myapp

WORKDIR /usr/src/myapp

RUN go build main.go

CMD ['./main']
