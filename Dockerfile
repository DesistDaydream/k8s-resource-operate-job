FROM golang:1.16 as builder
WORKDIR /root/operate

ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOPROXY=https://goproxy.cn,https://goproxy.io,https://mirrors.aliyun.com/goproxy/,direct
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download && go get k8s.io/client-go@kubernetes-1.19.2

COPY ./ /root/operate
RUN go build -o operate ./*.go

FROM alpine
RUN echo "hosts: files dns" > /etc/nsswitch.conf
WORKDIR /root/operate
COPY --from=builder /root/operate/operate /usr/local/bin/operate
ENTRYPOINT ["/usr/local/bin/operate"]
