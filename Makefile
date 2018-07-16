test: local-redis	
	docker build .

local-redis:
	docker-compose up -d redis

local-proxy:
	go build -v -o proxy .
	./proxy -redis-addr localhost:6379

cleanup:
	docker system prune -f
	docker volume prune -f

redis-cli: local-redis
	docker exec -it segment-redis-proxy_redis_1 redis-cli
