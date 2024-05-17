run:
	go run ./cmd/api/main.go

build:
	go build -o backend ./cmd/api/main.go

test:
	go test -cover ./...

format:
	gofmt -w -s .

swag-init:
	swag init --parseDependency --parseInternal -g **/**/*.go 

swag-fmt:
	swag fmt -g **/**/*.go
