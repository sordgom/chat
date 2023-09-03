DB_URL=postgresql://admin:password123@localhost:6510/auth_db?sslmode=disable

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up
migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...

.phony:
	test migrateup migratedown sqlc