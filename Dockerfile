

FROM ccr.ccs.tencentyun.com/images-base/golang:latest

MAINTAINER Razil "chenhuan@liaosearch.com"


RUN mkdir -p $GOPATH/src/webpimage

WORKDIR $GOPATH/src/webpimage

ADD ./ $GOPATH/src/webpimage

RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct

RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "main"



EXPOSE 8000

ENTRYPOINT ["./main"]


