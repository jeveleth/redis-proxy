FROM golang:latest
WORKDIR /go/src/github.com/jeveleth/segment-redis-proxy
COPY . .
# TODO: dep ensure
# RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux  go build -v -o proxy .
RUN go test -cache-capacity 2

# TODO: Get tests talking to redis
FROM alpine:latest  
RUN apk --no-cache add ca-certificates && \
    apk update && apk add bash
WORKDIR /root/
COPY --from=0 /go/src/github.com/jeveleth/segment-redis-proxy/proxy .
EXPOSE 8080
CMD ["./proxy"]