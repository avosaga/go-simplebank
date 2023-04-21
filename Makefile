postgres:
	docker run --name postgresAlpine -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine

createdb:
	docker exec -it postgresAlpine createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgresAlpine dropdb simple_bank

build-image:
	docker build -t go-simplebank:latest .

run-container:
	docker run --name go-simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@postgresAlpine:5432/simple_bank?sslmode=disable" go-simplebank:latest

create-network:
	docker network create bank-network

network-connect:
	docker network connect bank-network postgresAlpine

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup-by-1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown-by-1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -build_flags=--mod=mod -package mockdb -destination db/mock/store.go github.com/avosaga/go-simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb sqlc migratedown migrateup migratedown-by-1 migrateup-by-1 test server mock network-connect run-container build-image create-network
