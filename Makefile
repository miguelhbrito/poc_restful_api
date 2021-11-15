run-gateway:
	echo "running the api server"
	sudo ./scripts/run-server.sh
config-up:
	sudo docker-compose up -d
config-down:
	docker-compose down