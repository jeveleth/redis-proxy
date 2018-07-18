FROM golang:latest
WORKDIR /go/src/github.com/jeveleth/redis-proxy
COPY . .
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
    dep ensure && \
    CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -v -o proxy .
EXPOSE 8080
CMD ["./proxy"]
