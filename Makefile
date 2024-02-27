postgres:
	docker run --name pgColibriTech -p 5432:5432 -e POSTGRES_USER=colibriuser -e POSTGRES_PASSWORD=colibripassword -d postgres:16-alpine
createdb:
	docker exec -it pgColibriTech createdb --username=colibriuser --owner=colibriuser colibriCRM
dropdb:
	docker exec -it pgColibriTech dropdb colibriCRM
migrateup:
	migrate -path ./db/migrations -database "postgresql://colibriuser:colibripassword@localhost:5432/colibriCRM?sslmode=disable" -verbose up
migratedown:
	migrate -path ./db/migrations -database "postgresql://colibriuser:colibripassword@localhost:5432/colibriCRM?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...

run:
	docker start pgColibriTech
	go run ./cmd/falcon/main.go
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test run