DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.4-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank
	
dropdb:
	docker exec -it postgres14 dropdb simple_bank

migrateup:
	migrate --path db/migration --database "$(DB_URL)" --verbose up

migrateup1:
	migrate --path db/migration --database "$(DB_URL)" --verbose up 1

migratedown:
	migrate --path db/migration --database "$(DB_URL)" --verbose down

migratedown1:
	migrate --path db/migration --database "$(DB_URL)" --verbose down 1

db_docs: ## Generate database docs
	dbdocs build doc/db.dbml

db_schema: ## Generate database schema
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml 

sqlc:
	sqlc generate

test:
	go test -v -cover ./...
	
server:
	go run main.go

mock:
	mockgen --build_flags=--mod=mod --destination db/mock/store.go --package mockdb github.com/wizlif/simplebank/db/sqlc Store 

proto_fix:
	protolint lint -fix proto;

proto:
	rm -rf pb/*.go
	rm -rf doc/swagger/*.swagger.json
	rm -rf doc/statik/*
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb \
	--grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
	--openapiv2_out=doc/swagger \
	--openapiv2_opt=logtostderr=true,allow_merge=true,merge_file_name=simplebank \
    proto/*.proto
	statik -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 9090 -r repl


.PHONY: postgres createdb dropdb migratedown migratedown1 migrateup migrateup1 test server mock proto proto_fix evans
