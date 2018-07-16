test: local-redis	
	docker build .

local-redis:
	docker-compose up -d redis

docker-proxy-config:
	docker-compose up -d
	docker exec -it redis-proxy_proxy_1 ./proxy -help

# If you do *not* have Go set up on your machine
docker-proxy: local-redis
	docker-compose up -d
	docker exec -it redis-proxy_proxy_1 bash

# If you have Go set up on your machine
# local-proxy:
# 	go build -v -o proxy .
# 	./proxy -redis-addr localhost:6379

cleanup:
	docker-compose down
	docker system prune -f
	docker volume prune -f
	docker rmi -f `docker images -q`
	docker ps -a
	docker images -a

redis-cli: local-redis
	docker exec -it redis-proxy_redis_1 redis-cli

