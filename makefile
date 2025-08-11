.PHONY: build run docker-up clean

# Build the Go binary locally
build:
	go build -o taskmanager ./cmd/main.go

# Run the built binary locally (use after build)
run: build
	./taskmanager

# Build Docker images and start containers using docker-compose
docker-up:
	docker-compose up --build

# Stop containers and remove the binary
clean:
	docker-compose down
	rm -f taskmanager