# 第一阶段：构建 Go 应用
FROM golang:1.23-alpine AS builder

# 设置 Go 的工作目录
WORKDIR /app
# 复制源代码
COPY . .

# 编译 Go 应用

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

# 第二阶段：构建精简的运行时镜像
FROM alpine:latest

# 安装一些基础工具（可选）
# RUN apk --no-cache add ca-certificates
ENV TZ=Asia/Shanghai
RUN apk update && apk add --no-cache tzdata openntpd \
    && ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制二进制文件到容器中
COPY --from=builder /app/server .
COPY --from=builder /app/config.yaml .

# 暴露应用端口
EXPOSE 8080

# 启动应用
CMD ["./server"]
