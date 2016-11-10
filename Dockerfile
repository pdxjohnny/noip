FROM alpine:latest
MAINTAINER John Andersen <johnandersenpdx@gmail.com>

RUN apk add --no-cache go git ca-certificates && \
  export GOPATH=/gopath && \
  mkdir -pv $GOPATH && \
  go install -v github.com/pdxjohnny/noip && \
  cp -v $GOPATH/bin/noip /bin/noip && \
  rm -rfv $GOPATH && \
  apk del go git
