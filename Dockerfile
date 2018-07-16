FROM golang:latest
WORKDIR /go/src/github.com/jeveleth/segment-redis-proxy
COPY . .
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
    dep ensure && \
    CGO_ENABLED=0 GOARCH=amd64 GOOS=linux  go build -v -o proxy . && \
    go test -redis-addr localhost:6379 -cache-capacity 2

# TODO: Get tests talking to redis instance
FROM alpine:latest  
RUN apk --no-cache add ca-certificates && \
    apk update && apk add bash
WORKDIR /root/
COPY --from=0 /go/src/github.com/jeveleth/segment-redis-proxy/proxy .
EXPOSE 8080
CMD ["./proxy"]