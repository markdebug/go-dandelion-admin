FROM golang:1.20 as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /go-admin

COPY . .

WORKDIR /go-admin/authorize

RUN go build -a -o authorize .

FROM alpine as final

RUN apk add --no-cache tzdata
ENV TZ="Asia/Shanghai"

WORKDIR /opt/app

COPY --from=builder /go-admin/authorize/authorize .
EXPOSE 8999
CMD ["/opt/app/authorize"]
