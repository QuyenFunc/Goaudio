postgres:
	docker run --name postgresAu -p 2211:5432 -e POSTGES_USER=root -e POSTGES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root Audio
dropdb:
	docker exec -it postgres16 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" down
migrate_fix:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" force 17
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
.PHONY:postgres createdb dropdb migrateup migratedown sqlc test
