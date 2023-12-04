FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add tzdata ca-certificates bash curl

COPY ./backend-go /app

ENTRYPOINT ["/app/backend-go", "server"]