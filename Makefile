.PHONY: help build run test clean docker-up docker-down docker-logs

help:
	@echo "Available targets:"
	@echo "  build        Build Go binary"
	@echo "  run          Run Go application locally"
	@echo "  test         Run tests"
	@echo "  clean        Remove built binary"
	@echo "  docker-up    Start Docker Compose services"
	@echo "  docker-down  Stop Docker Compose services"
	@echo "  docker-logs  View Docker Compose logs"

build:
	go build -o app

run:
	go run main.go

test:
	go test ./...

clean:
	rm -f app

docker-up:
	docker compose up -d

docker-down:
	docker compose down

docker-logs:
	docker compose logs -f