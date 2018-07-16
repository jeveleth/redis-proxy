FROM golang:latest
WORKDIR /go/src/github.com/jeveleth/redis-proxy
COPY . .
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
    dep ensure && \
    CGO_ENABLED=0 GOARCH=amd64 GOOS=linux  go build -v -o proxy . && \
    go test -redis-addr host.docker.internal:6379 -v 
EXPOSE 8080
CMD ["./proxy"]
