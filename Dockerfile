# ビルド用環境
FROM golang:1.19.2-bullseye AS deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app

# ----------------------------------------------
# 本番環境
FROM debian:bullseye-slim AS deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app"]

# ----------------------------------------------
# 開発環境

FROM golang:1.19.2-alpine3.16 AS dev

WORKDIR /go/src

# gccの環境を入れる
RUN apk update && apk add alpine-sdk

COPY ./go.mod ./go.sum  ./
RUN go mod tidy
COPY ./ ./

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

CMD ["go","run","server.go"]
