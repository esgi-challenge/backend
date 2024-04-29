run:
	go run ./cmd/api/main.go

build:
	go build ./cmd/api/main.go

test:
	go test -cover ./...

dev:
	echo "Starting docker environment"
	docker-compose -f docker-compose.yml up --build
