FROM golang:alpine
# 更换alpine下载镜像
RUN sed -i 's!http://dl-cdn.alpinelinux.org/!https://mirrors.ustc.edu.cn/!g' /etc/apk/repositories
RUN apk add --no-cache git gcc libc-dev curl
RUN mkdir -p /go/src
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOSUMDB=sum.golang.google.cn
RUN go env -w GO111MODULE=on
WORKDIR /go/src
COPY . .
RUN go mod tidy && go build -o sys_app main.go
EXPOSE 8081
CMD ["/go/src/sys_app"]