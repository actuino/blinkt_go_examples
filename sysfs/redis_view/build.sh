#!/bin/sh

docker build --build-arg http_proxy=$http_proxy --build-arg https_proxy=$https_proxy -t alexellis2/redis-view:builder . -f Dockerfile.build && \
docker create --name sysfs-builder alexellis2/redis-view:builder && \
docker cp sysfs-builder:/go/src/github.com/alexellis/blinkt_go_examples/sysfs/progress/progress . && \
docker rm -f sysfs-builder && \
docker build -t alexellis2/redis-view:latest .


