FROM golang:latest
WORKDIR /go/src/github.com/jeveleth/redis-proxy
COPY . .
RUN go mod download && \
CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -v -o proxy .
EXPOSE 8080
CMD ["./proxy"]
