FROM golang:1.12.7
LABEL maintainer="XieWei"
WORKDIR /go/src/GopherBook
EXPOSE 2345
ENV GO111MODULE=on
VOLUME [ "/go/src/GopherBook/chapter12/logs" ]

RUN apt-get update && apt-get install -q -y vim git openssh-client cron bash && apt-get clean;

COPY chapter12/fifa/go.mod .
COPY chapter12/fifa/go.sum .
COPY chapter12/fifa .
RUN make install
RUN make run
CMD [ "bash", "-c", "/go/src/GopherBook/chapter12/fifa/fifa sync; /go/src/GopherBook/chapter12/fifa/fifa;" ]
