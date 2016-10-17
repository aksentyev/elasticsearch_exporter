FROM golang:1.7-alpine

MAINTAINER Grigory Aksentyev <grigory.aksentiev@gmail.com>

RUN mkdir -p /go/src/github.com/aksentyev/elastic_exporter
COPY . /go/src/github.com/aksentyev/elastic_exporter
RUN apk add --no-cache git build-base

ENV GOROOT /usr/local/go
RUN cd /go/src/github.com/aksentyev/elastic_exporter \
    && go get -d \
    && go test -v \
    && go build -o /bin/elastic_exporter \
    && rm -rf /go/src/github.com/aksentyev/elastic_exporter

RUN apk del --purge git build-base

RUN mkdir /config
WORKDIR /config

ENTRYPOINT ["/bin/elastic_exporter", "-log.level", "debug"]
CMD ["-update-interval", "3", "-scrape-interval", "3"]
