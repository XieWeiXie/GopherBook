FROM golang:1.12.7
LABEL maintainer="XieWei"
WORKDIR /go/src/GopherBook/chapter12
EXPOSE 2345
ENV GO111MODULE=on
VOLUME [ "/go/src/GopherBook/chapter12/fina/logs" ]

RUN apt-get update && apt-get install -q -y vim git openssh-client cron bash && apt-get clean;

COPY fina/go.mod .
COPY fina/go.sum .
COPY fina .
RUN make install
RUN make run
CMD [ "bash", "-c", "/go/src/GopherBook/chapter12/fina/fina sync; /go/src/GopherBook/chapter12/fina/fina;" ]
