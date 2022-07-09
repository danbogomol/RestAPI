run:
	docker-compose up -d

remove:
	docker rm -f server
	docker rm -f database
	docker rmi restapi_server

restart: remove run
