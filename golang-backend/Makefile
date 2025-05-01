.PHONY: build run clean migrateup migratedown sqlcgen
POSTGRES_USER := postgres
POSTGRES_PASSWORD := password
POSTGRES_DB := hospital
# Database URL
DB_URL := "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable"
build:
	go mod tidy
	go build -o bin/app

run: build
	./bin/app

clean:
	go clean
	rm -rf bin/

# Migrations
migrateup:
	migrate -path db/migrations -database "$(DB_URL)" up

migratedown:
	migrate -path db/migrations -database "$(DB_URL)" down

# SQLC code generation
setup:
	migrate -path db/migrations -database $(DB_URL) up
	cd db && sqlc generate

test:
	go test ./...
