# 使用官方的 golang 镜像作为基础镜像
FROM golang:1.22-alpine AS build

# 设置工作目录
WORKDIR /app

# 将本地代码复制到容器中
COPY . .

# 下载依赖包
RUN go mod download

# 编译 Go 应用程序
RUN go build -o main .

# 创建最终的镜像
FROM alpine:latest
# 安装 CA 证书以及设置工作目录
RUN apk --no-cache add ca-certificates
WORKDIR /root/
# 复制 .env 文件到容器中
COPY .env ./
# 从 builder 阶段复制可执行文件到最终镜像中
COPY --from=build /app/main .

EXPOSE 8080

# 运行应用程序
CMD ["./main"]
