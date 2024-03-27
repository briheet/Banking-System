.DEFAULT_GOAL := build

fmt:
	go fmt ./...

lint: fmt
	go lint ./...

vet: fmt
	go vet ./...

build: vet
	go build ./...

run: fmt
	go run ./...

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/go_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/go_bank?sslmode=disable" -verbose down

postgres:
	docker run --name postgreslatest -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest

createdb:
	docker exec -it postgreslatest createdb --username=root --owner=root go_bank

dropdb:
	docker exec -it postgreslatest dropdb go_bank

.PHONY: postgres createdb dropdb migrateup migratedown fmt lint vet build run
