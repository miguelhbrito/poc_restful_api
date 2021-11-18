run-server:
	echo "running the api server"
	chmod +x scripts/run-server.sh
	sh scripts/run-server.sh
config-up:
	sudo docker-compose up -d
config-down:
	docker-compose down