FROM golang:1.19.0-alpine3.16

WORKDIR /go/src

# gccの環境を入れる
RUN apk update && apk add alpine-sdk

# COPY ./go.mod ./go.sum  ./
# RUN go mod tidy
# COPY ./ ./

# CMD ["go","run","server.go"]
