DB_URL=postgresql://admin:password123@localhost:6510/auth_db?sslmode=disable

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up
migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down
sqlc:
	sqlc generate
mock:
	mockgen -package mockdb -destination db/mock/auth.go chat-app/auth-service/db/sqlc Auth 
test:
	go test -v -cover ./...

.phony:
	test migrateup migratedown sqlc mock