test: 
	docker-compose up -d --build 
	# go test
	
nuclear:
	docker system prune -f -a
	docker volume prune -f