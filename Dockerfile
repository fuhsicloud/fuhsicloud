FROM golang:1.13.0-alpine3.10 as build-env

ENV GO111MODULE=on
ENV BUILDPATH=github.com/fuhsicloud/fuhsicloud
ENV GOPROXY=https://goproxy.io
ENV GOPATH=/go
RUN mkdir -p /go/src/${BUILDPATH}
COPY ./ /go/src/${BUILDPATH}
RUN cd /go/src/${BUILDPATH} && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -v

FROM alpine:latest

RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        curl \
        && update-ca-certificates 2>/dev/null || true

COPY --from=build-env /go/bin/fuhsicloud /go/bin/fuhsicloud
COPY ./static /go/bin/static
COPY ./database /go/bin/database

WORKDIR /go/bin/
CMD ["/go/bin/fuhsicloud", "start", "-p", ":8080", "-c", "/etc/fuhsicloud/app.cfg", "-k", "/etc/fuhsicloud/config.yaml"]