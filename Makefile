run:
	go run ./cmd/api/main.go

build:
	go build -o backend ./cmd/api/main.go

test:
	go test -cover ./...

format:
	gofmt -w -s .

dev:
	echo "Starting docker environment"
	docker-compose -f docker-compose.yml up --build
