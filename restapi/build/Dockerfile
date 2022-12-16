FROM golang:1.19.1-alpine

RUN apk update && apk add git

RUN mkdir -p /go/src/github.com/restapi

ADD .. /go/src/github.com/restapi

RUN go install golang.org/x/tools/cmd/goimports@latest
RUN go install github.com/cosmtrek/air@latest

WORKDIR /go/src/github.com/restapi/cmd/server

CMD ["air", "-c", ".air.toml"]