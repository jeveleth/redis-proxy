test: local-redis
	docker build .

local-redis:
	docker-compose up -d redis

docker-proxy-config:
	docker-compose up -d
	docker exec -it redis-proxy_proxy_1 ./proxy -help

# To enter an interactive environment with Go
docker-proxy: local-redis
	docker-compose up -d
	docker exec -it redis-proxy_proxy_1 bash

cleanup:
	docker-compose down
	docker system prune -f
	docker volume prune -f
	docker rmi -f `docker images -q`
	docker ps -a
	docker images -a

redis-cli: local-redis
	docker exec -it redis-proxy_redis_1 redis-cli