test: 
	docker-compose up -d --build
	docker exec -it segment-redis-proxy_proxy_1 bash test.sh

local-redis:
	docker-compose up -d redis

local-proxy:
	go run *.go -redis-addr localhost:6379

nuclear:
	docker system prune -f -a
	docker volume prune -f

redis-cli:
	docker exec -it segment-redis-proxy_redis_1 redis-cli

# docker-proxy:
# 	docker exec -it segment-redis-proxy_proxy_1 ./proxy -redis-addr localhost:6379
tests:
	# create a test yaml to run docker like following
	go test -redis-addr localhost:6379 -cache-capacity 2
	