FROM golang:1.12.7
LABEL maintainer="XieWei"
WORKDIR /go/src/GopherBook
EXPOSE 8888
ENV GO111MODULE=on
VOLUME [ "/go/src/GopherBook/chapter11/logs" ]

RUN apt-get update && apt-get install -q -y vim git openssh-client cron bash && apt-get clean;

COPY chapter11/go.mod .
COPY chapter11/go.sum .
COPY chapter11 .
RUN make install
RUN make prod
CMD [ "bash", "-c", "/go/src/GopherBook/chapter11/votes;" ]
