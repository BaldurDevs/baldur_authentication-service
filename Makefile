build:
	@echo "Building auth binary..."
	env GOOS=linux CGO_ENABLED=0 go build -o authApp ./cmd
	@echo "Done!"

up_build: build
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"
