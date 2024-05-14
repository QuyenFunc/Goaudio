postgres:
	docker run --name postgresAu -p 2211:5432 -e POSTGES_USER=root -e POSTGES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root Audio
dropdb:
	docker exec -it postgres16 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/GoAudio?sslmode=disable" up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/GoAudio?sslmode=disable" down
migrate_fix:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/GoAudio?sslmode=disable" force 17
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/QuyenFunc/Goaudio/db/sqlc Store
.PHONY:postgres createdb dropdb migrateup migratedown sqlc test mock