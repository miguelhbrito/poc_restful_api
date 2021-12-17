run-server:
	echo "running api server"
	chmod +x scripts/run-server.sh
	sh scripts/run-server.sh
config-up:
	docker-compose up -d
config-down:
	docker-compose down
