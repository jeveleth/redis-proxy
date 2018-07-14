test: 
	docker-compose up -d --build 
	# go test

local: 
	go build -v -o redis-proxy .
	./redis-proxy

nuclear:
	docker system prune -f -a
	docker volume prune -f