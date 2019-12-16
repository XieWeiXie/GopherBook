FROM golang:1.12.7
LABEL maintainer="XieWei"

WORKDIR /go/GopherBook/chapter11
EXPOSE 8888
ENV GO111MODULE=on
VOLUME [ "/go/GopherBook/chapter11/logs" ]

RUN apt-get update && apt-get install -q -y vim git openssh-client cron bash && apt-get clean;

COPY . .
RUN make install
RUN make prod
CMD [ "bash", "-c", "/go/GopherBook/chapter11/votes;" ]
