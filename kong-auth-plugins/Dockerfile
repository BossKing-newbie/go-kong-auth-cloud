FROM golang:alpine as builder
# 更换alpine下载镜像
RUN sed -i 's!http://dl-cdn.alpinelinux.org/!https://mirrors.ustc.edu.cn/!g' /etc/apk/repositories
RUN apk add --no-cache git gcc libc-dev curl
RUN mkdir /go-plugins
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOSUMDB=sum.golang.google.cn
RUN go env -w GO111MODULE=on
COPY ./go-hello.go /go-plugins/go-hello.go
RUN cd /go-plugins && go mod init kong-go-plugin && go get github.com/Kong/go-pdk && go mod tidy && go build go-hello.go
#####################
## Release image
#####################
FROM kong:2.7.0-alpine
# Copy Go files
COPY --from=builder  /go-plugins/go-hello /usr/local/bin/
RUN mkdir -p /usr/local/kong
COPY gateway.yml /usr/local/kong/gateway.yml
ENV KONG_DECLARATIVE_CONFIG="/usr/local/kong/gateway.yml"
ENV KONG_DATABASE=off
ENV KONG_PLUGINS=bundled,go-hello
ENV KONG_PLUGINSERVER_NAMES=go-hello
ENV KONG_PLUGINSERVER_GO_HELLO_QUERY_CMD="go-hello -dump"

USER kong

