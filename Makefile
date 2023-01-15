AUTH_APP=authApp

build:
	@echo "Building auth binary..."
	env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_APP} ./cmd
	@echo "Done!"

up_build: build
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

remove_binary:
	@echo "Removing binary file"
	rm ${AUTH_APP} || true

remove_stop: remove_binary
	@echo "Stopping docker images"
	docker-compose down
