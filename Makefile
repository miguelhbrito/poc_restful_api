run-server:
	echo "running the api server"
	./scripts/run-server.sh
config-up:
	sudo docker-compose up -d
config-down:
	docker-compose down