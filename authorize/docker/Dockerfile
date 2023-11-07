FROM golang:1.20 as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /go-admin

COPY . .

RUN go mod tidy && go build -a -o authorize ./authorize

FROM alpine as final
WORKDIR /opt/app

COPY --from=builder /go-admin/authorize/authorize .
COPY --from=builder /go-admin/authorize/config ./config
EXPOSE 8999
CMD ["/opt/app/authorize", "server"]

