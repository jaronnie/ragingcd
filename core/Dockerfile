FROM alpine:latest

WORKDIR /app

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN mkdir -p /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

RUN apk --no-cache add tzdata ca-certificates libc6-compat libstdc++ curl bash

EXPOSE 8081

COPY ./core /app/
COPY ./config/config.toml /app/config/

ENTRYPOINT ["/app/core", "server"]