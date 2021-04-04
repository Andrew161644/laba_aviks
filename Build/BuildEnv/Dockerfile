FROM golang:1.12-alpine AS build_base

RUN apk add --no-cache git

WORKDIR src
RUN go get -v -d google.golang.org/grpc
RUN go get -v -d github.com/gorilla/securecookie
RUN go get -v -d github.com/gorilla/sessions
RUN go get -v -d github.com/lib/pq