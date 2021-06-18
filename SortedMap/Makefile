help: ## Show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

run: ## Run application
	cat text.json | go run main.go

build: ## Build application
	go build .

lint: ## Run linters
	golangci-lint run

test: ### Run tests
	go test ./...

bench: ### Run benchmarks
	go test -bench=. -benchmem ./...