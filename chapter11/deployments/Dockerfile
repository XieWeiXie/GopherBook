FROM golang:1.13.4
LABEL maintainer="XieWei"

WORKDIR /go/GopherBook/chapter11
EXPOSE 8888
ENV GO111MODULE=on
VOLUME [ "/go/GopherBook/chapter11/logs" ]

RUN apt-get update && apt-get install -q -y vim git openssh-client cron bash && apt-get clean;

COPY chapter11/* .

RUN make install
RUN make prod
CMD [ "bash", "-c", "/go/GopherBook/chapter11/votes;"]
