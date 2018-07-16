FROM golang:latest

WORKDIR /go/src/github.com/jeveleth/segment-redis-proxy
COPY . .
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
    dep ensure && \
    CGO_ENABLED=0 GOARCH=amd64 GOOS=linux  go build -v -o proxy . && \
    go test -redis-addr host.docker.internal:6379 -v
FROM alpine:latest  
RUN apk --no-cache add ca-certificates && \
    apk update && apk add bash
WORKDIR /root/
COPY --from=0 /go/src/github.com/jeveleth/segment-redis-proxy/proxy .
EXPOSE 8080
CMD ["./proxy"]