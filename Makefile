run-migrations:
	migrate -path db/migrations -database postgresql://postgres:postgres@localhost:5432/mini_bank?sslmode=disable up

rollback-migrations:
	migrate -path db/migrations -database postgresql://postgres:postgres@localhost:5432/mini_bank?sslmode=disable down

sqlc:
	sqlc generate

.PHONY: run-migrations rollback-migrations sqlc