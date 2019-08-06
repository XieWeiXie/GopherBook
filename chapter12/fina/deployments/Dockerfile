FROM golang:1.12.7
LABEL maintainer="XieWei"
WORKDIR /go/src/GopherBook
EXPOSE 2345
ENV GO111MODULE=on
VOLUME [ "/go/src/GopherBook/chapter12/fina/logs" ]

RUN apt-get update && apt-get install -q -y vim git openssh-client cron bash && apt-get clean;

COPY chapter12/fina/go.mod chapter12/fina/go.mod
COPY chapter12/fina/go.sum chapter12/fina/go.sum
COPY chapter12/fina chapter12/fina
COPY chapter12/fina/Makefile chapter12/fina/Makefile
RUN cd chapter12/fina && make install
RUN cd chapter12/fina && make run
CMD [ "bash", "-c", "/go/src/GopherBook/chapter12/fina/fina sync; /go/src/GopherBook/chapter12/fina/fina;" ]
