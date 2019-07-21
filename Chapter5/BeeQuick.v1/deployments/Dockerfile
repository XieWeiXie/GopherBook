FROM golang:1.11-alpine

LABEL maintainer="XieWei"

WORKDIR /go/src/GopherBook/Chapter5/BeeQuick.v1

ENV GO111MODULE on

EXPOSE 8080

RUN apk --no-cache --update add vim git curl make gcc libc-dev bash

COPY go.mod .
COPY go.sum .
COPY . .

RUN go mod download

RUN make production

CMD [ "bash", "-c", "/go/src/GopherBook/Chapter5/BeeQuick.v1 sync2 db; /go/src/GopherBook/Chapter5/BeeQuick.v1/BeeQuick"]
