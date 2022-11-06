DISTRIBUTED_CONFIG_GRPC_BINARY=distributedConfigApp

up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

up_build: distributed-config-grpc-service
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

distributed-config-grpc-service:
	@echo "Building user balance binary..."
	cd ./../. && env GOOS=linux CGO_ENABLED=0 go build -o ${DISTRIBUTED_CONFIG_GRPC_BINARY} ./cmd
	@echo "Done!"