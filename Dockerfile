FROM golang:1.13.8-alpine3.11 

ADD . /usr/src/myapp

WORKDIR /usr/src/myapp

RUN go build -v main.go
RUN mv main $GOPATH/bin

ENTRYPOINT main
