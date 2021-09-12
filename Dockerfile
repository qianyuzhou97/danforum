# 启动编译环境
FROM golang:1.17-alpine AS builder

# 配置编译环境
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

# 拷贝源代码到镜像中
COPY . /go/src/danforum

# 编译
WORKDIR /go/src/danforum
RUN go install ./cmd/...


FROM alpine:3.14
COPY --from=builder /go/bin/api /bin/api
# ENV ADDR=:8081

# 申明暴露的端口
EXPOSE 8000

# 设置服务入口
ENTRYPOINT [ "/bin/api" ]
