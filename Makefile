run:
	go run ./cmd/api/main.go

build:
	go build -o backend ./cmd/api/main.go

test:
	go test -v -cover ./...

format:
	gofmt -w -s .

swag: swag-fmt swag-init

swag-init:
	@if command -v swag > /dev/null 2>&1; then \
		swag init --parseDependency --parseInternal -g **/**/*.go; \
	else \
		echo "You must install go 'swag' command."; \
	fi

swag-fmt:
	@if command -v swag > /dev/null 2>&1; then \
		swag fmt -g **/**/*.go; \
	else \
		echo "You must install go 'swag' command."; \
	fi

mock:
	@if [ -z "$(CRUD)" ]; then echo "The variable CRUD is not set."; exit 1; fi
	@if command -v mockgen > /dev/null 2>&1; then \
		mockgen -source=internal/$(CRUD)/usecase.go -destination=internal/$(CRUD)/mock/usecase_mock.go -package=mock; \
		mockgen -source=internal/$(CRUD)/repository.go -destination=internal/$(CRUD)/mock/repository_mock.go -package=mock; \
	else \
		echo "You must install go 'mockgen' command."; \
	fi

crud:
	@if python3 -c "import jinja2" > /dev/null 2>&1; then \
		python3 scripts/crud.py; \
	else \
		echo "You must install 'jinja2' pip package for python."; \
	fi
