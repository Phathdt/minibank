.PHONY: migrate-up, migrate-down, run

migrate-up:
	migrate -path db/migrations -database $(DATABASE_URL) up

migrate-down:
	migrate -path db/migrations -database $(DATABASE_URL) down -all

run:
	go run cmd/app/main.go
