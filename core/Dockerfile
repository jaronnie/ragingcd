FROM alpine:latest

WORKDIR /app

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk --no-cache add tzdata ca-certificates bash curl

COPY ./core /app/

ENTRYPOINT ["/app/core", "server"]