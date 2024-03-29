DB_URL=postgresql://admin:password123@localhost:6510/auth_db?sslmode=disable

docker:
	docker compose up -d
postgres:
	docker run --name auth --network chat_network -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=password123 -p 6510:5432 -d postgres:latest
createdb: 
	docker exec -it auth createdb --username admin --owner=admin auth_db
dropdb:
	docker exec -it auth dropdb auth_db --username admin
test:
	go test -v -cover ./...
server:
	go run main.go
chat:
	go run main.go -server chat
auth:
	go run main.go -server auth
frontend:
	cd frontend && yarn start
	
.PHONY: test postgres createdb dropdb server chat auth frontend docker
