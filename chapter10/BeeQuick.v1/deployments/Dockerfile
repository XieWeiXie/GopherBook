FROM golang:1.13.4

LABEL maintainer="XieWei"

WORKDIR /go/GopherBook/chapter10/BeeQuick.v1

ENV GO111MODULE on

EXPOSE 8080

RUN apk --no-cache --update add vim git curl make gcc libc-dev bash

COPY chapter10/BeeQuick.v1/* .

RUN go mod download

RUN make production

CMD [ "bash", "-c", "/go/GopherBook/chapter10/BeeQuick.v1/BeeQuick sync2 db; /go/GopherBook/chapter10/BeeQuick.v1/BeeQuick"]
