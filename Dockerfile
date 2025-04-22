FROM golang:1.22-alpine

LABEL org.opencontainers.image.authors="riffert.daniel@gmail.com"

WORKDIR /src

COPY *.go go.mod /src

RUN go build -o app && mv app /usr/

EXPOSE 8010

ENTRYPOINT /usr/app
