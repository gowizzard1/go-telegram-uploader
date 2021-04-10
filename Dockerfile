FROM golang:1.14 AS base
LABEL maintainer="Isaac Kiptanui <isaackmaiwa@gmail.com>"
WORKDIR /go/src/powergen/go-telegram-uploader

FROM base as builder

FROM alpine:latest as prod
RUN apk add --no-cache bash 
COPY --from=builder /go/src/powergen/go-telegram-uploader/app .
EXPOSE 10107
CMD ["./app"]

FROM golang:1.14 AS debug
EXPOSE 40000 10107
RUN go get github.com/go-delve/delve/cmd/dlv
WORKDIR /go/src/powergen/go-telegram-uploader
COPY --from=base ./go/src/powergen/go-telegram-uploader .
CMD "go run ."
