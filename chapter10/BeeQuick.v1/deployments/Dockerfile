FROM golang:1.11-alpine

LABEL maintainer="XieWei"

WORKDIR /go/src/GopherBook

ENV GO111MODULE on

EXPOSE 8080

RUN apk --no-cache --update add vim git curl make gcc libc-dev bash

COPY Chapter5/BeeQuick.v1/go.mod .
COPY Chapter5/BeeQuick.v1/go.sum .
COPY Chapter5/BeeQuick.v1 .
COPY Chapter5/BeeQuick.v1/Makefile .

RUN go mod download

RUN make production

CMD [ "bash", "-c", "/go/src/GopherBook/Chapter5/BeeQuick.v1 sync2 db; /go/src/GopherBook/Chapter5/BeeQuick.v1/BeeQuick"]
