FROM golang:alpine as builder
# 更换alpine下载镜像
RUN sed -i 's!http://dl-cdn.alpinelinux.org/!https://mirrors.ustc.edu.cn/!g' /etc/apk/repositories
RUN apk add --no-cache git gcc libc-dev curl
RUN mkdir /go-plugins
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOSUMDB=sum.golang.google.cn
RUN go env -w GO111MODULE=on
WORKDIR /go-plugins
COPY . .
RUN go mod tidy && go build -o auth-plugin main.go
#####################
## Release image
#####################
FROM kong:2.7.0-alpine
# Copy Go files
COPY --from=builder  /go-plugins/auth-plugin /usr/local/bin/
RUN mkdir -p /usr/local/kong
COPY gateway.yml /usr/local/kong/gateway.yml
ENV KONG_DECLARATIVE_CONFIG="/usr/local/kong/gateway.yml"
ENV KONG_DATABASE=off
ENV KONG_PLUGINS=bundled,auth-plugin
ENV KONG_PLUGINSERVER_NAMES=auth-plugin
ENV KONG_PLUGINSERVER_AUTH_PLUGIN_QUERY_CMD="auth-plugin -dump"

USER kong

