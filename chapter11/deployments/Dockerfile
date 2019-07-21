FROM golang:1.12.7
LABEL maintainer="XieWei"
WORKDIR /go/src/GopherBook/chapter11
EXPOSE 8888
ENV GO111MODULE=on
VOLUME [ "/go/src/GopherBook/chapter11/logs" ]
COPY . .
RUN make install
RUN make prod
CMD [ "bash", "-c", /go/src/GopherBook/chapter11/votes;" ]
