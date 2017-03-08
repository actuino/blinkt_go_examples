FROM resin/rpi-raspbian
#actuino/blinkt-unicorn-client:1

# Install Go 1.7.5 and enough of a build-chain to build wiringpi (in C)
# Added nano and nfs client for dev purposes (not needed to compile nor run)
RUN apt-get update && \
    apt-get install -qy nano nfs-common build-essential wiringpi git curl ca-certificates && \
    curl -sSLO https://storage.googleapis.com/golang/go1.7.5.linux-armv6l.tar.gz && \
	mkdir -p /usr/local/go && \
	tar -xvf go1.7.5.linux-armv6l.tar.gz -C /usr/local/go/ --strip-components=1

ENV PATH=$PATH:/usr/local/go/bin/
ENV GOPATH=/go/

RUN mkdir -p /go/src/github.com/actuino/blinkt_go_examples/unicorn-client
WORKDIR /go/src/github.com/actuino/blinkt_go_examples/unicorn-client

ENV UNICORN_SERVER_HOST "192.168.7.3"
ENV UNICORN_SERVER_PORT 80

COPY app.go	.
COPY colors.json .
RUN go get -d -v

RUN go build

CMD ["./unicorn-client"]
