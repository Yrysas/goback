# Makefile for finance_tracker project

# Database connection string
DB_URL = postgres://postgres:postgres@localhost:5432/financetracker?sslmode=disable

# Migration directory
MIGRATION_DIR = migrations

# Install golang-migrate if not present
install-migrate:
	@if ! command -v migrate >/dev/null 2>&1; then \
		echo "Installing golang-migrate..."; \
		go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest; \
	fi

# Create a new migration file
migrate-create:
	@echo "Creating new migration..."
	@migrate create -ext sql -dir $(MIGRATION_DIR) -seq $(name)

# Run migrations up
migrate-up: install-migrate
	@echo "Running migrations up..."
	@migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" up

# Run migrations down
migrate-down: install-migrate
	@echo "Running migrations down..."
	@migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" down 1

# Run migrations down all
migrate-down-all: install-migrate
	@echo "Running all migrations down..."
	@migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" down

# Show migration status
migrate-status: install-migrate
	@echo "Migration status:"
	@migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" version

# Build the project
build:
	@echo "Building the project..."
	@go build -o bin/finance_tracker

# Run the project
run: build
	@echo "Running the project..."
	@./bin/finance_tracker

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf bin/

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Tidy modules
tidy:
	@echo "Tidying modules..."
	@go mod tidy

.PHONY: install-migrate migrate-create migrate-up migrate-down migrate-down-all migrate-status build run clean test fmt tidy