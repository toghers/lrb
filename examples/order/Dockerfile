FROM golang:alpine AS builder

#设置环境变量
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY ./etc /app/etc
RUN go build -ldflags="-s -w" -o /app/order .


FROM scratch

WORKDIR /app
COPY --from=builder /app/order /app/order
COPY --from=builder /app/etc /app/etc


CMD ["./order", "-f", "etc/order.yaml"]
