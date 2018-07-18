build:
	docker-compose build

run: build
	docker-compose up -d

test: run
	docker exec proxy go test -v

local-redis: build
	docker-compose up -d redis

docker-proxy-config: run
	docker exec proxy ./proxy -help

# To enter an interactive environment with Go
docker-proxy: run
	docker exec -it proxy bash

# Clean up all of your docker images.
# Be careful if you have non redis-proxy images/containers that you want to keep
nuclear:
	docker-compose down
	docker system prune -f
	docker volume prune -f
	docker rmi -f `docker images -q`
	docker ps -a
	docker images -a

redis-cli: local-redis
	docker exec -it redis redis-cli